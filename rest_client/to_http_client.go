package rest_client

import (
	"net/http"
	"net/url"
)

// ToHttpClient converts the client to an HTTP client
func (c *restClient) ToHttpClient() (*http.Client, error) {
	httpClient := http.Client{}

	// Proxy
	if c.proxy != nil {
		proxyURL, err := url.Parse(c.proxy.Dsn())
		if err != nil {
			return nil, err
		}
		proxy := http.ProxyURL(proxyURL)
		httpClient.Transport = &http.Transport{Proxy: proxy}
	}

	// Request timeout
	if c.requestTimeout > 0 {
		httpClient.Timeout = c.requestTimeout
	}

	return &httpClient, nil
}
