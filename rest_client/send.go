package rest_client

import (
	"github.com/sitnikovik/go-rest-client/proxy"
	"github.com/sitnikovik/go-rest-client/request"
	"net/http"
	"net/url"
)

// Send sends the request and returns the employee.
// No returns error if the employee status code is not 2xx!
func (c *restClient) Send(r request.Request) (*http.Response, error) {
	// Merge client params with request
	if err := c.mergeParamsWithRequest(r); err != nil {
		return nil, err
	}

	// Convert request to HTTP request
	httpReq, err := r.ToHttpRequest()
	if err != nil {
		return nil, err
	}

	// Prepare client
	cl, err := c.ToHttpClient()
	if err != nil {
		return nil, err
	}
	// User agent
	ua := c.userAgent
	if r.GetUserAgent() != "" {
		ua = r.GetUserAgent()
	}
	if ua != "" {
		httpReq.Header.Set("User-Agent", ua)
	}

	// Send request
	resp, err := cl.Do(httpReq)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// mergeParamsWithRequest prepares the client by the request
func (c *restClient) mergeParamsWithRequest(r request.Request) error {
	// Set proxy
	var prx proxy.Proxy
	if r.GetProxy() != nil {
		prx = r.GetProxy()
	} else {
		prx = c.proxy
	}
	if prx != nil && len(prx.GetExcludedHosts()) > 0 {
		u, err := url.Parse(r.GetUrl())
		if err != nil {
			return err
		}
		if prx.HasExcludedHost(u.Host) {
			prx = nil
		}
	}
	c.proxy = prx

	// Request timeout
	if r.GetRequestTimeout() != 0 {
		c.requestTimeout = r.GetRequestTimeout()
	}

	return nil
}
