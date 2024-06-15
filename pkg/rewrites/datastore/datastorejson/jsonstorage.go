package datastorejson

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/netip"
	"os"
	"slices"

	"github.com/awryme/dnsproxy/pkg/rewrites/datastore"
	"github.com/awryme/slogf"
)

type Storage struct {
	filePath string
	// config   *jsonConfig
}

func New(logf slogf.Logf, path string) (datastore.Storage, error) {
	s := &Storage{
		filePath: path,
	}
	logf("created rewrites cache storage", slog.String("json_path", path))

	return s, nil
}

func (s *Storage) readConfig() (*jsonConfig, error) {
	file, err := os.Open(s.filePath)
	if err != nil {
		return nil, fmt.Errorf("open rewrites json storage path %s: %w", s.filePath, err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("read rewrites json storage (path = %s): %w", s.filePath, err)
	}
	if len(data) == 0 {
		return &jsonConfig{
			Addrs:    map[string]string{},
			Rewrites: []rewriteEntry{},
		}, nil
	}

	var config jsonConfig
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("decode json rewrites config (path = %s): %w", s.filePath, err)
	}
	return &config, nil
}

func (s *Storage) writeConfig(cfg *jsonConfig) error {
	file, err := os.Create(s.filePath)
	if err != nil {
		return fmt.Errorf("open rewrites json storage path %s: %w", s.filePath, err)
	}
	defer file.Close()

	err = json.NewEncoder(file).Encode(cfg)
	if err != nil {
		return fmt.Errorf("encode rewrites config to file %s: %w", s.filePath, err)
	}
	return nil
}

func (s *Storage) ListRewrites() ([]datastore.Entry, error) {
	cfg, err := s.readConfig()
	if err != nil {
		return nil, fmt.Errorf("read rewrites storage: %w", err)
	}
	entries := make([]datastore.Entry, 0, len(cfg.Rewrites))
	for _, rewrite := range cfg.Rewrites {
		entries = append(entries, datastore.Entry{
			Domain:   rewrite.Domain,
			AddrName: rewrite.AddrName,
			Type:     rewrite.Type,
		})
	}
	return entries, nil
}

func (s *Storage) ListAddrs() (datastore.Addrs, error) {
	cfg, err := s.readConfig()
	if err != nil {
		return nil, fmt.Errorf("read rewrites storage: %w", err)
	}
	addrs := make(datastore.Addrs, len(cfg.Addrs))
	for name, ipstr := range cfg.Addrs {
		ip, err := netip.ParseAddr(ipstr)
		if err != nil {
			return nil, fmt.Errorf("parse addr %s from store: %w", ipstr, err)
		}
		addrs[name] = ip
	}
	return addrs, nil
}

func (s *Storage) AddAddr(name string, ip string) error {
	cfg, err := s.readConfig()
	if err != nil {
		return fmt.Errorf("read addrs storage: %w", err)
	}
	if cfg.Addrs == nil {
		cfg.Addrs = map[string]string{}
	}
	cfg.Addrs[name] = ip
	if err := s.writeConfig(cfg); err != nil {
		return fmt.Errorf("write addrs storage: %w", err)
	}
	return nil
}

func (s *Storage) DeleteAddr(name string) error {
	cfg, err := s.readConfig()
	if err != nil {
		return fmt.Errorf("read addrs storage: %w", err)
	}
	if len(cfg.Addrs) == 0 {
		return nil
	}
	delete(cfg.Addrs, name)
	if err := s.writeConfig(cfg); err != nil {
		return fmt.Errorf("write addrs storage: %w", err)
	}
	return nil
}

func (s *Storage) AddRewrite(domain string, matchType string, addr string) error {
	cfg, err := s.readConfig()
	if err != nil {
		return fmt.Errorf("read addrs storage: %w", err)
	}
	cfg.Rewrites = append(cfg.Rewrites, rewriteEntry{
		Domain:   domain,
		Type:     matchType,
		AddrName: addr,
	})
	if err := s.writeConfig(cfg); err != nil {
		return fmt.Errorf("write addrs storage: %w", err)
	}
	return nil
}

func (s *Storage) DeleteRewrite(domain string, matchType string) error {
	cfg, err := s.readConfig()
	if err != nil {
		return fmt.Errorf("read addrs storage: %w", err)
	}
	cfg.Rewrites = slices.DeleteFunc(cfg.Rewrites, func(entry rewriteEntry) bool {
		return entry.Domain == domain && entry.Type == matchType
	})
	if err := s.writeConfig(cfg); err != nil {
		return fmt.Errorf("write addrs storage: %w", err)
	}
	return nil
}
