package scanner

import (
	"net"
	"time"
)

type IPScanner struct{}

func (s *IPScanner) Scan(target string) (interface{}, error) {

	names, _ := net.LookupAddr(target)

	return map[string]interface{}{
		"ip_address": target,
		"geolocation": map[string]interface{}{
			"country": "Unknown",
			"city":    "Unknown",
		},
		"asn": map[string]interface{}{
			"number": 0,
			"name":   "Unknown",
		},
		"reverse_dns": names,
		"created_at":  time.Now(),
	}, nil
}
