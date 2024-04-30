package request

import (
	"testing"
)

// TestRequest_ToHttpRequest tests the ToHttpRequest method
func TestRequest_ToHttpRequest(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		r := NewRequest().
			SetUrl("https://google.com/:pathKey/").
			AddQueryParam("key", "value").
			AddPathParam("pathKey", "pathValue").
			SetFormData(map[string]string{"formKey": "formValue"}).
			AsPost()

		req, err := r.ToHttpRequest()
		if err != nil {
			t.Errorf("ToHttpRequest() error = %v, want nil", err)
			return
		}
		if req == nil {
			t.Errorf("ToHttpRequest() request is nil")
			return
		}
		if req.Method != r.GetMethod() {
			t.Errorf(
				"ToHttpRequest() Method = %s, want %s",
				req.Method, r.GetMethod(),
			)
			return
		}
		if req.URL.Path == r.GetUrl() {
			t.Errorf(
				"ToHttpRequest() URL.Path = %s, want %s",
				req.URL.Path, r.GetUrl(),
			)
			return
		}
		if req.URL.Query().Get("key") != "value" {
			t.Errorf(
				"ToHttpRequest() URL.Query().Get(key) = %s, want value",
				req.URL.Query().Get("key"),
			)
			return
		}
		if req.URL.Path != "/pathValue/" {
			t.Errorf(
				"ToHttpRequest() URL.Path = %s, want https://google.com/pathValue/",
				req.URL.Path,
			)
			return
		}
		if req.Header.Get("Content-Type") != "application/x-www-form-urlencoded" {
			t.Errorf(
				"ToHttpRequest() Content-Type = %s, want application/x-www-form-urlencoded",
				req.Header.Get("Content-Type"),
			)
			return
		}
		if req.ContentLength != 17 {
			t.Errorf(
				"ToHttpRequest() ContentLength = %d, want 19",
				req.ContentLength,
			)
			return
		}
		if req.Body == nil {
			t.Errorf("ToHttpRequest() Body is nil")
			return
		}
		defer func() {
			_ = req.Body.Close()
		}()
	})
}
