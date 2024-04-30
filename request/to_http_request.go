package request

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// ToHttpRequest creates an HTTP request from a request
func (r *request) ToHttpRequest() (*http.Request, error) {
	httpReq := http.Request{
		Header: make(http.Header),
	}

	httpReq.Method = r.method
	u, err := url.Parse(r.url)
	if err != nil {
		return nil, err
	}
	httpReq.URL = u

	// Set headers
	for k, v := range r.headers {
		httpReq.Header.Set(k, v)
	}

	// Set Auth
	if r.basicAuth.username != "" && r.basicAuth.password != "" {
		httpReq.SetBasicAuth(r.basicAuth.username, r.basicAuth.password)
	}
	if r.bearerAuth.token != "" {
		httpReq.Header.Set("Authorization", "Bearer "+r.bearerAuth.token)
	}

	// Set query params
	query := httpReq.URL.Query()
	for k, v := range r.queryParams {
		query.Add(k, v)
	}
	httpReq.URL.RawQuery = query.Encode()

	// Set path params
	for k, v := range r.pathParams {
		httpReq.URL.Path = strings.Replace(httpReq.URL.Path, ":"+k, v, -1)
	}

	// Set request body
	var body io.ReadCloser
	var length int64
	if len(r.formData) > 0 {
		httpReq.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		form := url.Values{}
		for k, v := range r.formData {
			form.Add(k, v)
		}
		body = io.NopCloser(strings.NewReader(form.Encode()))
		length = int64(len(form.Encode()))
	} else {
		// Set JSON data
		if r.jsonData != nil {
			httpReq.Header.Set("Content-Type", "application/json")
			jsonBody, err := json.Marshal(r.jsonData)
			if err != nil {
				return nil, err
			}
			body = io.NopCloser(strings.NewReader(string(jsonBody)))
			length = int64(len(jsonBody))
		}
	}
	if body != nil {
		httpReq.ContentLength = length
		httpReq.Body = body
	}

	return &httpReq, nil
}
