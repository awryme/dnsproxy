package rewrites

import (
	"net/netip"
	"slices"
)

type Entry struct {
	Domain   string
	AddrName string
	IP       netip.Addr
	Type     MatchType
}

type EntrySet struct {
	Strict []Entry
	Suffix []Entry
}

func (entries EntrySet) All() []Entry {
	return slices.Concat(entries.Strict, entries.Suffix)
}

func (entries EntrySet) SortedByDomain() []Entry {
	allRewrites := entries.All()
	slices.SortFunc(allRewrites, func(a, b Entry) int {
		if a.Domain < b.Domain {
			return -1
		}
		return 1
	})
	return allRewrites
}
