package scanner

import (
	"crypto/tls"
	"fmt"
	"net"
	"time"
)

type SSLScanner struct{}

func (s *SSLScanner) Scan(target string) (interface{}, error) {

	conn, err := tls.DialWithDialer(
		&net.Dialer{
			Timeout: 5 * time.Second,
		},
		"tcp",
		fmt.Sprintf("%s:443", target),
		&tls.Config{
			InsecureSkipVerify: true,
		},
	)

	if err != nil {
		return nil, err
	}

	defer conn.Close()

	cert := conn.ConnectionState().PeerCertificates[0]

	return map[string]interface{}{
		"certificate": map[string]interface{}{
			"subject":      cert.Subject.String(),
			"issuer":       cert.Issuer.String(),
			"valid_from":   cert.NotBefore,
			"valid_until":  cert.NotAfter,
			"is_expired":   time.Now().After(cert.NotAfter),
			"is_self_sign": cert.Subject.String() == cert.Issuer.String(),
			"san":          cert.DNSNames,
		},
		"created_at": time.Now(),
	}, nil
}
