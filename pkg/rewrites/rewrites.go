package rewrites

import (
	"net/netip"
	"slices"
)

type MatchType string

func (m MatchType) String() string {
	return string(m)
}

const (
	MatchTypeStrict MatchType = "strict"
	MatchTypeSuffix MatchType = "suffix"
)

type Entry struct {
	Domain   string
	AddrName string
	IP       netip.Addr
	Type     MatchType
}

type Rewrites struct {
	Strict []Entry
	Suffix []Entry
}

func (r Rewrites) All() []Entry {
	return slices.Concat(r.Strict, r.Suffix)
}

func (r Rewrites) SortedByDomain() []Entry {
	allRewrites := r.All()
	slices.SortFunc(allRewrites, func(a, b Entry) int {
		if a.Domain < b.Domain {
			return -1
		}
		return 1
	})
	return allRewrites
}
