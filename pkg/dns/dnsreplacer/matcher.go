package dnsreplacer

import "strings"

func StrictMatcher(dnsQ, domain string) bool {
	trimFQDN(&dnsQ, &domain)
	return dnsQ == domain
}

func SuffixMatcher(dnsQ, domain string) bool {
	trimFQDN(&dnsQ, &domain)
	if dnsQ == domain {
		return true
	}

	suffixDomain := "." + domain

	return strings.HasSuffix(dnsQ, suffixDomain)
}

func trimFQDN(addrRefs ...*string) {
	for _, ref := range addrRefs {
		*ref = strings.TrimSuffix(*ref, ".")
	}
}
