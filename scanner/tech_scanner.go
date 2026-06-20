package scanner

import (
	"net/http"
	"time"
)

type TechScanner struct{}

func (s *TechScanner) Scan(target string) (interface{}, error) {

	client := http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get("http://" + target)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	headers := map[string]string{
		"server":       resp.Header.Get("Server"),
		"x-powered-by": resp.Header.Get("X-Powered-By"),
		"content-type": resp.Header.Get("Content-Type"),
	}

	return map[string]interface{}{
		"domain":  target,
		"headers": headers,
		"technologies": []map[string]interface{}{
			{
				"name":       resp.Header.Get("Server"),
				"confidence": 100,
			},
		},
		"created_at": time.Now(),
	}, nil
}
