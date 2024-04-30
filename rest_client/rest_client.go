package rest_client

import (
	"github.com/sitnikovik/go-rest-client/proxy"
	"github.com/sitnikovik/go-rest-client/request"
	"net/http"
	"time"
)

// RestClient is the interface that wraps the basic methods of a REST client.
type RestClient interface {
	// GetRequestTimeout returns the request timeout of the client
	GetRequestTimeout() time.Duration
	// SetRequestTimeout sets the request timeout of the client
	SetRequestTimeout(timeout time.Duration)

	// GetUserAgent returns the user agent of the client
	GetUserAgent() string
	// SetUserAgent sets the user agent of the client
	SetUserAgent(userAgent string)

	// GetProxy returns the proxy of the client
	GetProxy() proxy.Proxy
	// SetProxy sets the proxy of the client
	SetProxy(proxy proxy.Proxy)

	// ToHttpClient converts the client to an HTTP client
	ToHttpClient() (*http.Client, error)

	// Send sends the request and returns the employee
	Send(r request.Request) (*http.Response, error)

	// SendAndDecodeToBytes sends the request and decodes the employee
	SendAndDecodeToBytes(r request.Request) ([]byte, error)

	// SendAndDecodeFromJson sends the request and decodes response to the given interface
	SendAndDecodeFromJson(r request.Request, v interface{}) error
}

// restClient is a basic implementation of the Client interface
type restClient struct {
	requestTimeout time.Duration
	userAgent      string
	proxy          proxy.Proxy
}

// NewRestClient returns a new instance of the Client interface
func NewRestClient() RestClient {
	return &restClient{}
}

// GetRequestTimeout returns the request timeout of the client
func (c *restClient) GetRequestTimeout() time.Duration {
	return c.requestTimeout
}

// SetRequestTimeout sets the request timeout of the client
func (c *restClient) SetRequestTimeout(timeout time.Duration) {
	c.requestTimeout = timeout
}

// GetUserAgent returns the user agent of the client
func (c *restClient) GetUserAgent() string {
	return c.userAgent
}

// SetUserAgent sets the user agent of the client
func (c *restClient) SetUserAgent(userAgent string) {
	c.userAgent = userAgent
}

// GetProxy returns the proxy of the client
func (c *restClient) GetProxy() proxy.Proxy {
	return c.proxy
}

// SetProxy sets the proxy of the client
func (c *restClient) SetProxy(proxy proxy.Proxy) {
	c.proxy = proxy
}
