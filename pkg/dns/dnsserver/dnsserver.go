package dnsserver

import (
	"context"
	"fmt"
	"log/slog"
	"net/url"

	"github.com/awryme/dnsproxy/pkg/dns/dnsclienttls"
	"github.com/awryme/dnsproxy/pkg/dns/dnshandler"
	"github.com/awryme/dnsproxy/pkg/dns/dnsreplacer"
	"github.com/awryme/slogf"
	"github.com/miekg/dns"
)

type Params struct {
	Addr            string
	Port            int
	Upstream        *url.URL
	BaseDns         string
	RewritesStorage dnsreplacer.Storage
}

func Start(ctx context.Context, logf slogf.Logf, params Params) error {
	logf = logf.With(slog.String("component", "dnsserver"))
	upstream := *params.Upstream
	switch upstream.Scheme {
	case "":
		return fmt.Errorf("missing upstream scheme in %s", upstream.String())
	case "tls":
		upstream.Scheme = ""
	default:
		return fmt.Errorf("upstream scheme is not tls")
	}
	upstreamType := "tls"
	upstreamAddr := upstream.Host

	replacer := dnsreplacer.New(params.RewritesStorage)

	upstreamClient := dnsclienttls.New(upstreamAddr, params.BaseDns)
	logf("created upstream client",
		slog.String("type", upstreamType),
		slog.String("addr", upstreamAddr),
		slog.String("base_dns", params.BaseDns),
	)

	handler := dnshandler.New(ctx, logf, upstreamClient, replacer)

	listenAddr := fmt.Sprintf("%s:%d", params.Addr, params.Port)
	server := dns.Server{
		Addr:    listenAddr,
		Net:     "udp",
		Handler: handler,
	}
	logf("starting dns server", slog.String("addr", params.Addr), slog.Int("port", params.Port))
	err := server.ListenAndServe()
	if err != nil {
		return fmt.Errorf("run dns server failed on addr %s: %w", listenAddr, err)
	}
	return nil
}
