# dnsproxy
DNS proxy, forwards dns to a different dns server, adds custom dns rewrites, provides UI for rewrites

Use web UI (default port = 8080) or file rewrites.json to provide rewrites

Special rewrite address `self` is automatically set to resolve to the address dnsproxy is listening on (set by flags)

Matching is configured by match types:
- strict - matches specific dns query to IP
- suffix - matches set dns query and all prefixed addresses to IP
  - suffix matcher for `twitter.com` will match `twitter.com` and `*.twitter.com`

If dns query is not matched by rewrites, it is forwarded to configured dns upstream

## Support DNS listening protocols
- Regular UDP DNS

## Supported DNS upstream protocols
- DNS over TLS

## flags/arguments
Output from `dnsproxy -h`

```
Usage: dnsproxy --rewrites-cfg=STRING --dns.upstream=DNS.UPSTREAM [flags]

Flags:
  -h, --help                         Show context-sensitive help.
      --rewrites-cfg=STRING          path to config with rewrites ($DNSPROXY_REWRITES_CFG)
      --addr="127.0.0.1"             address to listen on (udp, http), without port ($DNSPROXY_ADDR)
      --dns.port=53                  dns server: port to listen on (udp) ($DNSPROXY_DNS_PORT)
      --dns.upstream=DNS.UPSTREAM    dns server: upstream dns url (schemas: tls) ($DNSPROXY_DNS_UPSTREAM)
      --dns.base-dns="1.1.1.1"       dns server: base dns address for tls & https upstream resolve ($DNSPROXY_DNS_BASE_DNS)
      --ui.port=8080                 ui server: port to listen on (http) ($DNSPROXY_UI_PORT)
      --ui.log-access                log ui api requests ($DNSPROXY_UI_LOG_ACCESS)
```

## rewrites.json structure
Comments are added here, not supported by format!
```
{
    addrs: {
        "mycustomserver": "1.2.3.4" // sets 1.2.3.4 as mycustomserver
    },
    rewrites: [
        {
            "domain": "twitter.com",
            "type": "suffix",
            "addr_name": "self" // matches twitter.com and *.twitter.com to the same address as dnsproxy itself
        },
        {
            "domain": "service.dev",
            "type": "suffix",
            "addr_name": "mycustomserver" // matches only service.dev to mycustomserver (1.2.3.4)
        }
    ]
}
```
