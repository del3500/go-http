package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func getIPAddress(domain string) (string, error) {

	url := fmt.Sprintf("https://cloudflare-dns.com/dns-query?name=%s&type=A", domain)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("accept", "application/dns-json")

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error sending request: %w", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %w", err)
	}
	var dnsResp DNSResponse
	json.Unmarshal(body, &dnsResp)
	return dnsResp.Answer[0].Data, nil
}
