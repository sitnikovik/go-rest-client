package rest_client

import (
	"encoding/json"
	"fmt"
	"github.com/sitnikovik/go-rest-client/request"
	"io"
)

// SendAndDecodeToBytes sends the request and decodes response to byte slice
func (c *restClient) SendAndDecodeToBytes(r request.Request) ([]byte, error) {
	// Send request
	resp, err := c.Send(r)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	// Decode response
	bb, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != r.GetStatusExpected() {
		return bb, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return bb, nil
}

// SendAndDecodeFromJson sends the request and decodes response to the given interface
func (c *restClient) SendAndDecodeFromJson(r request.Request, v interface{}) error {
	// Send request
	resp, err := c.Send(r)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	// Decode response
	bb, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(bb, v); err != nil {
		return err
	}

	if resp.StatusCode != r.GetStatusExpected() {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}
