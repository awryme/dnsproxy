package dnsreplacer

import "github.com/awryme/dnsproxy/pkg/rewrites"

type Storage interface {
	GetRewrites() rewrites.Rewrites
}
