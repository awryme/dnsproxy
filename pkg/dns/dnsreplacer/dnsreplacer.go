package dnsreplacer

import (
	"net"
	"net/netip"

	"github.com/awryme/dnsproxy/pkg/rewrites"
	"github.com/miekg/dns"
)

const defaultTTL = 60

type Replacer struct {
	storage Storage
}

func New(storage Storage) *Replacer {
	return &Replacer{
		storage: storage,
	}
}

func (r *Replacer) ReplaceAddr(req *dns.Msg) (*dns.Msg, bool) {
	rewrites := r.storage.GetRewrites()

	resp := new(dns.Msg)
	var replacedOk bool

	for _, q := range req.Question {
		rr, ok := r.matchAddr(rewrites, q)
		if ok {
			resp.Answer = append(resp.Answer, rr)
			replacedOk = true
		}
	}
	return resp, replacedOk
}

func (r *Replacer) matchAddr(rewrites rewrites.Rewrites, q dns.Question) (dns.RR, bool) {
	for _, entry := range rewrites.Strict {
		if StrictMatcher(q.Name, entry.Domain) {
			return makeDnsAnswer_A(q, entry.IP), true
		}
	}
	for _, entry := range rewrites.Suffix {
		if SuffixMatcher(q.Name, entry.Domain) {
			return makeDnsAnswer_A(q, entry.IP), true
		}
	}
	return nil, false
}

func makeDnsAnswer_A(q dns.Question, ip netip.Addr) dns.RR {
	return &dns.A{
		Hdr: dns.RR_Header{
			Name:   q.Name,
			Rrtype: q.Qtype,
			Class:  q.Qclass,
			Ttl:    defaultTTL,
		},
		A: net.IP(ip.AsSlice()),
	}
}
