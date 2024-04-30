package rest_client

import (
	"github.com/sitnikovik/go-rest-client/request"
	"reflect"
	"testing"
)

// TestRestClient_SendAndDecodeToBytes tests the SendAndDecodeToBytes method
func TestRestClient_SendAndDecodeToBytes(t *testing.T) {
	t.Run("ok with google.com 200", func(t *testing.T) {
		req := request.NewRequest().
			SetUrl("https://www.google.com").
			AsGet()

		bb, err := NewRestClient().SendAndDecodeToBytes(req)

		if err != nil {
			t.Errorf("Error: %v but want nil", err)
			return
		}
		if bb == nil {
			t.Errorf("Response body is nil")
			return
		}
		if len(bb) == 0 {
			t.Errorf("Response body is empty")
			return
		}
	})

	t.Run("err with employee equals special string and err code expected", func(t *testing.T) {
		req := request.NewRequest().
			SetUrl("https://software.hixie.ch/utilities/cgi/test-tools/http-error?status=500+Internal+Server+Error").
			SetStatusExpected(500).
			AsGet()

		bb, err := NewRestClient().SendAndDecodeToBytes(req)

		if err != nil {
			t.Errorf("Error (%v) but want nil", err)
			return
		}
		if bb == nil {
			t.Errorf("Response body is nil")
			return
		}
		if len(bb) == 0 {
			t.Errorf("Response body is empty")
			return
		}

		actual := string(bb)
		expected := "Responded with: 500 Internal Server Error\n"
		if actual != expected {
			t.Errorf("Response body (%s) is not equals expected: %s)", actual, expected)
			return
		}
	})

	t.Run("err on 500 returned", func(t *testing.T) {
		req := request.NewRequest().
			SetUrl("https://software.hixie.ch/utilities/cgi/test-tools/http-error?status=500+Internal+Server+Error").
			AsGet()

		_, err := NewRestClient().SendAndDecodeToBytes(req)

		if err == nil {
			t.Errorf("Error is nil but error is expected")
			return
		}
	})
}

// TestRestClient_SendAndDecodeFromJson tests the TestRestClient_SendAndDecodeFromJson method
func TestRestClient_SendAndDecodeFromJson(t *testing.T) {
	t.Run("ok with google.com 200", func(t *testing.T) {
		req := request.NewRequest().
			SetUrl("https://mocki.io/v1/e597d386-ef58-45fa-81ad-dd9a62a066f7").
			AsGet()

		type employee struct {
			Name           string `json:"name"`
			Age            int    `json:"age"`
			Salary         int    `json:"salary"`
			SalaryCurrency string `json:"salaryCurrency"`
		}
		expected := employee{
			Name:           "John Doe",
			Age:            20,
			Salary:         10000,
			SalaryCurrency: "USD",
		}

		v := struct {
			Employee employee `json:"employee"`
		}{}
		err := NewRestClient().SendAndDecodeFromJson(req, &v)

		if err != nil {
			t.Errorf("Error: %v but want nil", err)
			return
		}
		if reflect.DeepEqual(v.Employee, expected) == false {
			t.Errorf("Response body (%v) is not equals expected (%v)", v.Employee, expected)
			return
		}
	})
}
