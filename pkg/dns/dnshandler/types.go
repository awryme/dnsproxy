package dnshandler

import (
	"context"

	"github.com/miekg/dns"
)

type Client interface {
	Send(ctx context.Context, msg *dns.Msg) (*dns.Msg, error)
}

type Replacer interface {
	ReplaceAddr(req *dns.Msg) (*dns.Msg, bool)
}
