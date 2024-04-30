package rest_client

import "testing"

// TestRestClient_ToHttpClient tests the RestClient.ToHttpClient function
func TestRestClient_ToHttpClient(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		c := NewRestClient()
		httpClient, err := c.ToHttpClient()
		if err != nil {
			t.Errorf("Error: %v", err)
			return
		}
		if httpClient == nil {
			t.Error("httpClient is nil")
			return
		}
	})
}
