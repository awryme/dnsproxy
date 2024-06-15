package rewrites

import (
	"net/netip"
)

const SelfAddrName = "self"

type Addr struct {
	Name string
	IP   netip.Addr
}
