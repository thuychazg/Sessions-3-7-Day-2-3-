package scanner

import (
	"fmt"
	"net"
	"time"
)

type PortScanner struct{}

func isPrivateIP(ip string) bool {

	addr := net.ParseIP(ip)

	if addr == nil {
		return false
	}

	if addr.IsLoopback() {
		return true
	}

	if addr.IsPrivate() {
		return true
	}

	return false
}

func (s *PortScanner) Scan(target string) (interface{}, error) {

	if !isPrivateIP(target) {
		return nil, fmt.Errorf("refuse to scan public ip")
	}

	ports := []int{22, 80, 443, 8080}

	var openPorts []map[string]interface{}

	for _, port := range ports {

		conn, err := net.DialTimeout(
			"tcp",
			fmt.Sprintf("%s:%d", target, port),
			3*time.Second,
		)

		if err == nil {

			openPorts = append(openPorts,
				map[string]interface{}{
					"port":     port,
					"protocol": "tcp",
					"state":    "open",
				})

			conn.Close()
		}
	}

	return map[string]interface{}{
		"ip_address": target,
		"open_ports": openPorts,
		"created_at": time.Now(),
	}, nil
}
