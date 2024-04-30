package request

import (
	"github.com/sitnikovik/go-rest-client/proxy"
	"net/http"
	"time"
)

// Request is the interface that wraps the basic methods of a request.
type Request interface {
	// GetRequestTimeout returns the request timeout
	GetRequestTimeout() time.Duration
	// SetRequestTimeout sets the request timeout
	SetRequestTimeout(timeout time.Duration) Request

	// GetUrl returns the URL
	GetUrl() string
	// SetUrl sets the URL
	SetUrl(url string) Request

	// GetStatusExpected returns the status code expected from the response.
	// If the response status code is different, an error is returned.
	GetStatusExpected() int
	// SetStatusExpected sets the status code expected from the response.
	// If the response status code is different, an error is returned.
	SetStatusExpected(status int) Request

	// GetUserAgent returns the user agent
	GetUserAgent() string
	// SetUserAgent sets the user agent
	SetUserAgent(userAgent string) Request

	// GetProxy returns the proxy
	GetProxy() proxy.Proxy
	// SetProxy sets the proxy
	SetProxy(proxy proxy.Proxy) Request

	// GetHeader returns header by key
	GetHeader(key string) string
	// HasHeader returns true if the header exists
	HasHeader(key string) bool
	// SetHeaders sets headers
	SetHeaders(headers map[string]string) Request
	// AddHeader sets a header
	AddHeader(key, value string) Request

	// SetBasicAuth sets basic auth
	SetBasicAuth(username, password string) Request
	// SetBearerAuth sets bearer auth
	SetBearerAuth(token string) Request

	// SetQueryParams sets query parameters
	SetQueryParams(params map[string]string) Request
	// AddQueryParam sets a query parameter
	AddQueryParam(key, value string) Request

	// SetPathParams sets path parameters
	SetPathParams(params map[string]string) Request
	// AddPathParam sets a path parameter
	AddPathParam(key, value string) Request

	// SetFormData sets form data  body
	SetFormData(data map[string]string) Request
	// SetJsonData sets JSON data  body
	SetJsonData(data interface{}) Request

	// GetMethod returns the HTTP method
	GetMethod() string
	// AsGet sets the HTTP method to GET
	AsGet() Request
	// AsPost sets the HTTP method to POST
	AsPost() Request
	// AsPut sets the HTTP method to PUT
	AsPut() Request
	// AsPatch sets the HTTP method to PATCH
	AsPatch() Request
	// AsDelete sets the HTTP method to DELETE
	AsDelete() Request

	// ToHttpRequest creates an HTTP request from a request
	ToHttpRequest() (*http.Request, error)
}

// request is a basic implementation of the Request interface
type request struct {
	requestTimeout time.Duration
	url            string
	method         string
	statusExpected int
	userAgent      string
	proxy          proxy.Proxy
	basicAuth      struct {
		username string
		password string
	}
	bearerAuth struct {
		token string
	}
	headers     map[string]string
	queryParams map[string]string
	pathParams  map[string]string
	formData    map[string]string
	jsonData    interface{}
}

// NewRequest returns a new instance of the Request interface
func NewRequest() Request {
	return &request{
		method:         http.MethodGet,
		statusExpected: http.StatusOK,

		headers:     make(map[string]string),
		queryParams: make(map[string]string),
		pathParams:  make(map[string]string),
		formData:    make(map[string]string),
	}
}

// GetRequestTimeout returns the request timeout
func (r *request) GetRequestTimeout() time.Duration {
	return r.requestTimeout
}

// SetRequestTimeout sets the request timeout
func (r *request) SetRequestTimeout(timeout time.Duration) Request {
	r.requestTimeout = timeout

	return r
}

// GetUrl returns the URL
func (r *request) GetUrl() string {
	return r.url
}

// SetUrl sets the URL
func (r *request) SetUrl(url string) Request {
	r.url = url

	return r
}

// GetStatusExpected returns the status code expected from the response.
func (r *request) GetStatusExpected() int {
	return r.statusExpected
}

// SetStatusExpected sets the status code expected from the response.
func (r *request) SetStatusExpected(status int) Request {
	r.statusExpected = status

	return r
}

// GetUserAgent returns the user agent
func (r *request) GetUserAgent() string {
	return r.headers["User-Agent"]
}

// SetUserAgent sets the user agent
func (r *request) SetUserAgent(userAgent string) Request {
	return r.AddHeader("User-Agent", userAgent)
}

// GetProxy returns the proxy
func (r *request) GetProxy() proxy.Proxy {
	return r.proxy
}

// SetProxy sets the proxy
func (r *request) SetProxy(proxy proxy.Proxy) Request {
	r.proxy = proxy

	return r
}

// GetHeader returns header by key
func (r *request) GetHeader(key string) string {
	return r.headers[key]
}

// HasHeader returns true if the header exists
func (r *request) HasHeader(key string) bool {
	_, ok := r.headers[key]

	return ok
}

// SetHeaders sets headers
func (r *request) SetHeaders(headers map[string]string) Request {
	r.headers = headers

	return r
}

// AddHeader sets a header
func (r *request) AddHeader(key, value string) Request {
	r.headers[key] = value

	return r
}

// SetBasicAuth sets basic auth
func (r *request) SetBasicAuth(username, password string) Request {
	r.basicAuth.username = username
	r.basicAuth.password = password

	return r
}

// SetBearerAuth sets bearer auth
func (r *request) SetBearerAuth(token string) Request {
	r.bearerAuth.token = token

	return r
}

// SetQueryParams sets query parameters
func (r *request) SetQueryParams(params map[string]string) Request {
	r.queryParams = params

	return r
}

// AddQueryParam sets a query parameter
func (r *request) AddQueryParam(key, value string) Request {
	r.queryParams[key] = value

	return r
}

// SetPathParams sets path parameters
func (r *request) SetPathParams(params map[string]string) Request {
	r.pathParams = params

	return r
}

// AddPathParam sets a path parameter
func (r *request) AddPathParam(key, value string) Request {
	r.pathParams[key] = value

	return r
}

// SetFormData sets form data body and clears JSON data
func (r *request) SetFormData(data map[string]string) Request {
	r.jsonData = nil
	r.formData = data

	return r
}

// SetJsonData sets JSON data body and clears form data
func (r *request) SetJsonData(data interface{}) Request {
	r.formData = nil
	r.jsonData = data

	return r
}

// GetMethod returns the HTTP method
func (r *request) GetMethod() string {
	return r.method
}

// AsGet sets the HTTP method to GET
func (r *request) AsGet() Request {
	r.method = http.MethodGet

	return r
}

// AsPost sets the HTTP method to POST
func (r *request) AsPost() Request {
	r.method = http.MethodPost

	return r
}

// AsPut sets the HTTP method to PUT
func (r *request) AsPut() Request {
	r.method = http.MethodPut

	return r
}

// AsPatch sets the HTTP method to PATCH
func (r *request) AsPatch() Request {
	r.method = http.MethodPatch

	return r
}

// AsDelete sets the HTTP method to DELETE
func (r *request) AsDelete() Request {
	r.method = http.MethodDelete

	return r
}
