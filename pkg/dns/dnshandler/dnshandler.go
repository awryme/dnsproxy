package dnshandler

import (
	"context"
	"log/slog"

	"github.com/awryme/slogf"
	"github.com/miekg/dns"
	"github.com/oklog/ulid/v2"
)

type Handler struct {
	ctx      context.Context
	logf     slogf.Logf
	client   Client
	replacer Replacer
}

func New(ctx context.Context, logf slogf.Logf, client Client, replacer Replacer) *Handler {
	return &Handler{ctx, logf, client, replacer}
}

func (h *Handler) ServeDNS(w dns.ResponseWriter, req *dns.Msg) {
	reqId := ulid.Make().String()
	logf := h.logf.With(
		slog.String("req_id", reqId),
		slog.String("remote_addr", w.RemoteAddr().String()),
	)
	logDnsRequest(logf, "received dns request", req)

	resp, ok := h.replacer.ReplaceAddr(req)
	if ok {
		h.writeResponse(logf, w, req, resp, "rewrite")
		return
	}

	resp, err := h.client.Send(h.ctx, req)
	if err != nil {
		logf("failed to send dns request", slogf.Error(err))
		h.responseError(logf, w, req)
		return
	}
	h.writeResponse(logf, w, req, resp, "upstream")
}

func (h *Handler) writeResponse(logf slogf.Logf, w dns.ResponseWriter, req *dns.Msg, resp *dns.Msg, responseType string) {
	resp.SetReply(req)
	err := w.WriteMsg(resp)
	if err != nil {
		logf("failed to send dns response", slog.String("response_type", responseType), slogf.Error(err))
		return
	}
	logDnsResponse(logf, "sent dns response", resp, responseType)
}

func (h *Handler) responseError(logf slogf.Logf, w dns.ResponseWriter, req *dns.Msg) {
	err := w.WriteMsg(new(dns.Msg).SetRcode(req, dns.RcodeServerFailure))
	if err != nil {
		logf("failed to write dns response with server failure", slogf.Error(err))
	}
}

func logDnsRequest(logf slogf.Logf, msg string, req *dns.Msg) {
	for idx, q := range req.Question {
		typeValue := dns.TypeToString[q.Qtype]
		classValue := dns.ClassToString[q.Qclass]
		logf(msg,
			slog.Int("q_idx", idx),
			slog.String("name", q.Name),
			slog.String("type", typeValue),
			slog.String("class", classValue),
		)
	}
}

func logDnsResponse(logf slogf.Logf, msg string, resp *dns.Msg, responseType string) {
	mapAddressToAttr := func(rr dns.RR) slog.Attr {
		if rr == nil {
			return slog.Attr{}
		}
		switch v := rr.(type) {
		case *dns.A:
			return slog.String("ip", v.A.String())
		case *dns.AAAA:
			return slog.String("ip", v.AAAA.String())
		default:
			return slog.Attr{}
		}
	}
	for idx, ans := range resp.Answer {
		header := ans.Header()
		typeValue := dns.TypeToString[header.Rrtype]
		classValue := dns.ClassToString[header.Class]
		logf(msg,
			slog.String("response_type", responseType),
			slog.Int("ans_idx", idx),
			slog.String("name", header.Name),
			slog.String("type", typeValue),
			slog.String("class", classValue),
			slog.Uint64("ttl", uint64(header.Ttl)),
			mapAddressToAttr(ans),
		)
	}
}
