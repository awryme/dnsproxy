package rewrites

import (
	"fmt"
	"net/netip"
	"slices"

	"github.com/awryme/dnsproxy/pkg/rewrites/datastore"
)

// CacheStorage stores and gives rewrites and addr mappings
// It should cache them between updates
type CacheStorage struct {
	datastore datastore.Storage

	selfIP netip.Addr

	// todo: close with locks
	cachedRewrites Rewrites
	cachedAddrs    []Addr
}

func NewStorage(datastore datastore.Storage, selfAddr string) (*CacheStorage, error) {
	selfIP, err := netip.ParseAddr(selfAddr)
	if err != nil {
		return nil, fmt.Errorf("parse self addr: %w", err)
	}
	s := &CacheStorage{
		datastore: datastore,
		selfIP:    selfIP,
	}

	if err := s.updateCaches(); err != nil {
		return nil, fmt.Errorf("update store caches: %w", err)
	}
	return s, nil
}

func (s *CacheStorage) GetRewrites() Rewrites {
	return s.cachedRewrites
}

func (s *CacheStorage) GetAddrs() []Addr {
	return s.cachedAddrs
}

func (s *CacheStorage) AddRewrite(domain string, matchType string, addr string) error {
	// todo: check if exists
	err := s.datastore.AddRewrite(domain, matchType, addr)
	if err != nil {
		return fmt.Errorf("add rewrite %s (type = %s, addr = %s) to store: %w", domain, matchType, addr, err)
	}
	return s.updateCaches()
}

func (s *CacheStorage) DeleteRewrite(domain string, matchType string) error {
	err := s.datastore.DeleteRewrite(domain, matchType)
	if err != nil {
		return fmt.Errorf("delete rewrite %s (type = %s) from store: %w", domain, matchType, err)
	}
	return s.updateCaches()
}

func (s *CacheStorage) AddAddr(name string, ip string) error {
	if name == SelfAddrName {
		return fmt.Errorf("adding self addr is forbidden")
	}
	// todo: check existing?
	err := s.datastore.AddAddr(name, ip)
	if err != nil {
		return fmt.Errorf("add addr %s (ip = %s) tp store: %w", name, ip, err)
	}
	return s.updateCaches()
}

func (s *CacheStorage) DeleteAddr(name string) error {
	if name == SelfAddrName {
		return fmt.Errorf("deleting self addr is forbidden")
	}

	for _, entry := range s.GetRewrites().All() {
		if entry.AddrName == name {
			return fmt.Errorf("deleting addr %s is forbidden, it is used in rewrites", name)
		}
	}
	err := s.datastore.DeleteAddr(name)
	if err != nil {
		return fmt.Errorf("delete addr %s from store: %w", name, err)
	}
	return s.updateCaches()
}

func (s *CacheStorage) updateCaches() error {
	storeAddrs, err := s.datastore.ListAddrs()
	if err != nil {
		return fmt.Errorf("list addrs from store: %w", err)
	}
	if _, ok := storeAddrs[SelfAddrName]; ok {
		return fmt.Errorf("overwriting self addr is forbidden")

	}

	addrs := make([]Addr, 0, len(storeAddrs))
	for name, ip := range storeAddrs {
		addrs = append(addrs, Addr{
			Name: name,
			IP:   ip,
		})
	}

	slices.SortFunc(addrs, func(a, b Addr) int {
		if a.Name < b.Name {
			return -1
		}
		return 1
	})

	storeRewrites, err := s.datastore.ListRewrites()
	if err != nil {
		return fmt.Errorf("list rewrites from store: %w", err)
	}
	rewrites, err := mapStoreRewrites(storeRewrites, storeAddrs, s.selfIP)
	if err != nil {
		return fmt.Errorf("map rewrite from store: %w", err)
	}

	selfAddr := Addr{
		Name: SelfAddrName,
		IP:   s.selfIP,
	}
	s.cachedAddrs = append([]Addr{selfAddr}, addrs...)
	s.cachedRewrites = rewrites
	return nil
}

func mapStoreRewrites(storeRewrites []datastore.Entry, storeAddrs datastore.Addrs, selfIP netip.Addr) (res Rewrites, _ error) {
	for _, storeRewrite := range storeRewrites {
		ip, ok := mapDatastoreIP(storeRewrite.AddrName, storeAddrs, selfIP)
		if !ok {
			return res, fmt.Errorf("rewrite address name %s not found", storeRewrite.AddrName)
		}
		matchType := MatchType(storeRewrite.Type)
		switch matchType {
		case MatchTypeStrict:
			res.Strict = append(res.Strict, Entry{
				Domain:   storeRewrite.Domain,
				IP:       ip,
				AddrName: storeRewrite.AddrName,
				Type:     matchType,
			})
		case MatchTypeSuffix:
			res.Suffix = append(res.Suffix, Entry{
				Domain:   storeRewrite.Domain,
				IP:       ip,
				AddrName: storeRewrite.AddrName,
				Type:     matchType,
			})
		default:
			return res, fmt.Errorf("received bad match type for rewrite: %s", storeRewrite.Type)
		}
	}
	return res, nil
}

func mapDatastoreIP(name string, storeAddrs datastore.Addrs, selfIP netip.Addr) (netip.Addr, bool) {
	if name == SelfAddrName {
		return selfIP, true
	}
	ip, ok := storeAddrs[name]
	return ip, ok
}
