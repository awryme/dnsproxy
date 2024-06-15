package datastorejson

type rewriteEntry struct {
	Domain   string `json:"domain"`
	Type     string `json:"type"`
	AddrName string `json:"addr_name"`
}

type jsonConfig struct {
	Addrs    map[string]string `json:"addrs"`
	Rewrites []rewriteEntry    `json:"rewrites"`
}
