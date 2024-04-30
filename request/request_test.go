package request

import (
	"github.com/sitnikovik/go-rest-client/proxy"
	"net/http"
	"testing"
	"time"
)

// TestNewRequest tests the NewRequest function
func TestNewRequest(t *testing.T) {
	r := NewRequest()

	if r == nil {
		t.Errorf("NewRequest() = nil")
	}

	if r.GetMethod() != http.MethodGet {
		t.Errorf("GetMethod() = %s but want %s", r.GetMethod(), http.MethodGet)
	}
	if r.GetUrl() != "" {
		t.Errorf("GetUrl() = %s but want empty string", r.GetUrl())
	}
	if r.GetProxy() != nil {
		t.Errorf("GetProxy() = %v but want nil", r.GetProxy())
	}
	if r.GetStatusExpected() != http.StatusOK {
		t.Errorf("GetStatusExpected() = %d but want 0", http.StatusOK)
	}
	if r.GetRequestTimeout() != 0 {
		t.Errorf("GetRequestTimeout() = %d but want 0", r.GetRequestTimeout())
	}
}

// TestRequest_GetMethod tests the GetMethod method
func TestRequest_GetMethod(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		r := NewRequest()

		if r.GetMethod() != http.MethodGet {
			t.Errorf("GetMethod() = %s but want %s", r.GetMethod(), http.MethodGet)
		}
	})
}

// TestRequest_GetUrl tests the GetUrl method
func TestRequest_GetUrl(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		r := NewRequest()

		if r.GetUrl() != "" {
			t.Errorf("GetUrl() = %s but want empty", r.GetUrl())
		}
	})
}

// TestRequest_SetUrl tests the SetUrl method
func TestRequest_SetUrl(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		r := NewRequest().
			SetUrl("http://localhost")

		if r.GetUrl() != "http://localhost" {
			t.Errorf("GetUrl() = %s but want http://localhost", r.GetUrl())
		}
	})
}

// TestRequest_GetProxy tests the GetProxy method
func TestRequest_GetProxy(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		r := NewRequest().
			SetProxy(proxy.NewProxy())

		if r.GetProxy() == nil {
			t.Errorf("GetProxy() = nil")
		}
	})
}

// TestRequest_SetProxy tests the SetProxy method
func TestRequest_SetProxy(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		r := NewRequest().
			SetProxy(proxy.NewProxy())

		if r.GetProxy() == nil {
			t.Errorf("GetProxy() = nil")
		}
	})
}

// TestRequest_GetStatusExpected tests the GetStatusExpected method
func TestRequest_GetStatusExpected(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		r := NewRequest()

		if r.GetStatusExpected() != http.StatusOK {
			t.Errorf("GetStatusExpected() = %d but want %d", r.GetStatusExpected(), http.StatusOK)
		}
	})
}

// TestRequest_SetStatusExpected tests the SetStatusExpected method
func TestRequest_SetStatusExpected(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		r := NewRequest().
			SetStatusExpected(http.StatusExpectationFailed)

		if r.GetStatusExpected() != http.StatusExpectationFailed {
			t.Errorf("GetStatusExpected() = %d but want %d", r.GetStatusExpected(), http.StatusExpectationFailed)
		}
	})
}

// TestRequest_GetRequestTimeout tests the GetRequestTimeout method
func TestRequest_GetRequestTimeout(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		r := NewRequest()

		if r.GetRequestTimeout() != 0 {
			t.Errorf("GetRequestTimeout() = %d but want 0", r.GetRequestTimeout())
		}
	})
}

// TestRequest_SetRequestTimeout tests the SetRequestTimeout method
func TestRequest_SetRequestTimeout(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		tmt := 10 * time.Second
		r := NewRequest().
			SetRequestTimeout(tmt)

		if r.GetRequestTimeout() != tmt {
			t.Errorf("GetRequestTimeout() = %d but want %d", r.GetRequestTimeout(), tmt)
		}
	})
}

// TestRequest_SetHeaders tests the SetHeaders method
func TestRequest_SetHeaders(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		r := NewRequest().
			SetHeaders(map[string]string{"Content-Type": "application/json"})

		if r.GetHeader("Content-Type") != "application/json" {
			t.Errorf("GetHeader() = %s but want application/json", r.GetHeader("Content-Type"))
		}
	})
}

// TestRequest_AddHeader tests the AddHeader method
func TestRequest_AddHeader(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		r := NewRequest().
			AddHeader("Content-Type", "application/json")

		if r.GetHeader("Content-Type") != "application/json" {
			t.Errorf("GetHeader() = %s but want application/json", r.GetHeader("Content-Type"))
		}

		r.AddHeader("Content-Type", "application/xml")
		if r.GetHeader("Content-Type") != "application/xml" {
			t.Errorf("GetHeader() = %s but want application/xml", r.GetHeader("Content-Type"))
		}

		r.AddHeader("Authorization", "Bearer token")
		if r.GetHeader("Authorization") != "Bearer token" {
			t.Errorf("GetHeader() = %s but want Bearer token", r.GetHeader("Authorization"))
		}
		if r.GetHeader("Content-Type") != "application/xml" {
			t.Errorf("GetHeader() = %s but want application/xml", r.GetHeader("Content-Type"))
		}
	})
}

// TestRequest_GetHeader tests the GetHeader method
func TestRequest_GetHeader(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		r := NewRequest().
			SetHeaders(map[string]string{"Content-Type": "application/json"})

		if r.GetHeader("Content-Type") != "application/json" {
			t.Errorf("GetHeader() = %s but want application/json", r.GetHeader("Content-Type"))
		}
	})
}

// TestRequest_HasHeader tests the HasHeader method
func TestRequest_HasHeader(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		r := NewRequest().
			SetHeaders(map[string]string{"Content-Type": "application/json"})

		if !r.HasHeader("Content-Type") {
			t.Errorf("HasHeader() = false but want true")
		}
	})
}

// TestRequest_GetUserAgent tests the GetUserAgent method
func TestRequest_GetUserAgent(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		r := NewRequest().
			SetUserAgent("Mozilla/5.0")

		if r.GetUserAgent() != "Mozilla/5.0" {
			t.Errorf("GetUserAgent() = %s but want Mozilla/5.0", r.GetUserAgent())
		}
	})
}

// TestRequest_SetUserAgent tests the SetUserAgent method
func TestRequest_SetUserAgent(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		ua := "Mozilla/5.0"
		r := NewRequest().
			SetUserAgent(ua)

		if r.GetUserAgent() != ua {
			t.Errorf("GetUserAgent() = %s but want %s", r.GetUserAgent(), ua)
		}

		ua = "Chrome/5.0"
		r.SetUserAgent(ua)
		if r.GetUserAgent() != ua {
			t.Errorf("GetUserAgent() = %s but want %s", r.GetUserAgent(), ua)
		}
	})
}

// TestRequest_AsPost tests the AsPost method
func TestRequest_AsGet(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		r := NewRequest().
			AsGet()

		if r.GetMethod() != http.MethodGet {
			t.Errorf("GetMethod() = %s but want %s", r.GetMethod(), http.MethodGet)
		}
	})
}

// TestRequest_AsPost tests the AsPost method
func TestRequest_AsPost(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		r := NewRequest().
			AsPost()

		if r.GetMethod() != http.MethodPost {
			t.Errorf("GetMethod() = %s but want %s", r.GetMethod(), http.MethodPost)
		}
	})
}

// TestRequest_AsPut tests the AsPut method
func TestRequest_AsPut(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		r := NewRequest().
			AsPut()

		if r.GetMethod() != http.MethodPut {
			t.Errorf("GetMethod() = %s but want %s", r.GetMethod(), http.MethodPut)
		}
	})
}

// TestRequest_AsPatch tests the AsPatch method
func TestRequest_AsPatch(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		r := NewRequest().
			AsPatch()

		if r.GetMethod() != http.MethodPatch {
			t.Errorf("GetMethod() = %s but want %s", r.GetMethod(), http.MethodPatch)
		}
	})
}

// TestRequest_AsDelete tests the AsDelete method
func TestRequest_AsDelete(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		r := NewRequest().
			AsDelete()

		if r.GetMethod() != http.MethodDelete {
			t.Errorf("GetMethod() = %s but want %s", r.GetMethod(), http.MethodDelete)
		}
	})
}

// TestRequest_SetFormData tests the SetFormData method
func TestRequest_SetFormData(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		r := NewRequest().
			SetFormData(map[string]string{"key": "value"})

		// check no panic
		r.SetJsonData(map[string]string{"key": "value"})
	})
}

// TestRequest_SetJsonData tests the SetJsonData method
func TestRequest_SetJsonData(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		r := NewRequest().
			SetJsonData(map[string]string{"key": "value"})

		// check no panic
		r.SetFormData(map[string]string{"key": "value"})
	})
}

// TestRequest_SetPathParams tests the SetPathParams method
func TestRequest_SetPathParams(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		// check no panic
		NewRequest().SetPathParams(map[string]string{"key": "value"})
	})
}

// TestRequest_AddPathParam tests the AddPathParam method
func TestRequest_AddPathParam(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		// check no panic
		NewRequest().AddPathParam("key", "value")
	})
}

// TestRequest_SetQueryParams tests the SetQueryParams method
func TestRequest_SetQueryParams(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		// check no panic
		NewRequest().SetQueryParams(map[string]string{"key": "value"})
	})
}

// TestRequest_AddQueryParam tests the AddQueryParam method
func TestRequest_AddQueryParam(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		// check no panic
		NewRequest().AddQueryParam("key", "value")
	})
}
