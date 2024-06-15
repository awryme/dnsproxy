package datastore

import (
	"net/netip"
)

type Addrs map[string]netip.Addr

type Entry struct {
	Domain   string
	AddrName string
	Type     string
}

// Storage stores and gives rewrites
// It should cache them between updates
type Storage interface {
	ListRewrites() ([]Entry, error)
	AddRewrite(domain string, matchType string, addr string) error
	DeleteRewrite(domain string, matchType string) error

	ListAddrs() (Addrs, error)
	AddAddr(name string, ip string) error
	DeleteAddr(name string) error
}
