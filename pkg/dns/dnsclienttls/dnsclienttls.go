package dnsclienttls

import (
	"context"
	"fmt"
	"net"

	"github.com/awryme/dnsproxy/pkg/dns/udpdnsresolver"
	"github.com/miekg/dns"
)

type Client struct {
	// dialer     *tls.Dialer
	client     *dns.Client
	remoteAddr string
}

func New(remoteTlsAddr string, baseDnsAddr string) *Client {
	baseResolver := udpdnsresolver.New(baseDnsAddr)
	baseDialer := &net.Dialer{
		Resolver: baseResolver,
	}
	client := &dns.Client{
		Net:    "tcp-tls",
		Dialer: baseDialer,
	}
	return &Client{
		// dialer: &tls.Dialer{
		// 	NetDialer: baseDialer,
		// },
		remoteAddr: remoteTlsAddr,
		client:     client,
	}
}

const UDPPacketSize = 512

func (c *Client) Send(ctx context.Context, msg *dns.Msg) (*dns.Msg, error) {
	// fixme: port?
	addr := fmt.Sprintf("%s:%d", c.remoteAddr, 853)
	respMsg, _, err := c.client.ExchangeContext(ctx, msg, addr)
	if err != nil {
		return nil, fmt.Errorf("exchange dns with addr %s: %w", addr, err)
	}
	return respMsg, nil
}
