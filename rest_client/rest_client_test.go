package rest_client

import (
	"github.com/sitnikovik/go-rest-client/proxy"
	"testing"
	"time"
)

// TestNewRestClient tests the NewRestClient function
func TestNewRestClient(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		c := NewRestClient()

		if c == nil {
			t.Errorf("NewRestClient() = nil")
			return
		}

		if c.GetRequestTimeout() != 0 {
			t.Errorf("Request timeout is not 0")
			return
		}
		if c.GetUserAgent() != "" {
			t.Errorf("User agent is not empty")
			return
		}
		if c.GetProxy() != nil {
			t.Errorf("Proxy is not nil")
			return
		}
	})
}

// TestRestClient_GetRequestTimeout tests the GetRequestTimeout method
func TestRestClient_GetRequestTimeout(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		c := NewRestClient()

		if c.GetRequestTimeout() != 0 {
			t.Errorf("Request timeout is not 0")
			return
		}
	})
}

// TestRestClient_SetRequestTimeout tests the SetRequestTimeout method
func TestRestClient_SetRequestTimeout(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		c := NewRestClient()
		c.SetRequestTimeout(10 * time.Second)

		if c.GetRequestTimeout() != 10*time.Second {
			t.Errorf("Request timeout is not %d", 10*time.Second)
		}
	})
}

// TestRestClient_GetUserAgent tests the GetUserAgent method
func TestRestClient_GetUserAgent(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		c := NewRestClient()

		if c.GetUserAgent() != "" {
			t.Errorf("User agent is not empty")
		}
	})
}

// TestRestClient_SetUserAgent tests the SetUserAgent method
func TestRestClient_SetUserAgent(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		c := NewRestClient()
		c.SetUserAgent("test")

		if c.GetUserAgent() != "test" {
			t.Errorf("User agent is not test")
		}
	})
}

// TestRestClient_GetProxy tests the GetProxy method
func TestRestClient_GetProxy(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		c := NewRestClient()

		if c.GetProxy() != nil {
			t.Errorf("Proxy is not nil")
		}
	})
}

// TestRestClient_SetProxy tests the SetProxy method
func TestRestClient_SetProxy(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		c := NewRestClient()
		c.SetProxy(proxy.NewProxy())

		if c.GetProxy() == nil {
			t.Errorf("Proxy is nil")
		}
	})
}
