package dnslink

import (
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/ipfs-gateway-monitor/pkg/config"
	"github.com/cloudflare/ipfs-gateway-monitor/pkg/context"
)

func FetchIpfsDomainAsUrlHostname(ctx *context.Context, cfg *config.DnslinkConfig) (time.Duration, error) {
	target := cfg.IpfsDomain
	// TODO: Assert that "target" CNAME's to the gateway and has DNSLink set to the IPFS content.
	start := time.Now()
	resp, err := http.Get(fmt.Sprintf("https://%s", target))
	elapsed := time.Since(start)
	if err != nil {
		return elapsed, err
	}
	if resp.StatusCode != http.StatusOK {
		return elapsed, fmt.Errorf("The status code is not correct: expected: %v got: %v", http.StatusOK, resp.StatusCode)
	}
	return elapsed, err
}
