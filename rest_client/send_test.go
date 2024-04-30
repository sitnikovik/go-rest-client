package rest_client

import (
	"encoding/json"
	"github.com/sitnikovik/go-rest-client/internal/config"
	"github.com/sitnikovik/go-rest-client/proxy"
	"github.com/sitnikovik/go-rest-client/request"
	"io"
	"testing"
)

// TestRestClient_Send tests the Send method
func TestRestClient_Send(t *testing.T) {
	t.Run("ok with google.com 200", func(t *testing.T) {
		req := request.NewRequest().
			SetUrl("https://www.google.com").
			AsGet()

		c := NewRestClient()
		resp, err := c.Send(req)
		if err != nil {
			t.Errorf("Error: %v", err)
			return
		}
		if resp.StatusCode != 200 {
			t.Errorf("Response code is not ok: %v", resp.StatusCode)
			return
		}
	})

	t.Run("ok with proxy", func(t *testing.T) {
		conf := config.FromFile().Proxy

		prx := proxy.NewProxy().
			SetHost(conf.Host).
			SetPort(conf.Port)
		req := request.NewRequest().
			SetUrl("https://api.ipify.org/?format=json").
			SetProxy(prx).
			AsGet()

		// Send
		c := NewRestClient()
		resp, err := c.Send(req)
		if err != nil {
			t.Errorf("Error: %v", err)
			return
		}

		// Response
		if resp.StatusCode != 200 {
			t.Errorf("Response code is not ok: %v", resp.StatusCode)
			return
		}
		bb, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Errorf("Error on convert body: %v", err)
			return
		}
		v := struct {
			Ip string `json:"ip"`
		}{}
		if err = json.Unmarshal(bb, &v); err != nil {
			t.Errorf("Error on unmarshal body: %v", err)
			return
		}

		// Assert
		if v.Ip != conf.Host {
			t.Errorf("Proxy is not used: %v", v.Ip)
			return
		}
	})

	t.Run("ok with proxy exclude hosts", func(t *testing.T) {
		conf := config.FromFile().Proxy

		prx := proxy.NewProxy().
			SetHost(conf.Host).
			SetPort(conf.Port).
			SetExcludedHosts([]string{"api.ipify.org"})
		req := request.NewRequest().
			SetUrl("https://api.ipify.org/?format=json").
			SetProxy(prx).
			AsGet()

		// Send
		c := NewRestClient()
		resp, err := c.Send(req)
		if err != nil {
			t.Errorf("Error: %v", err)
			return
		}

		// Response
		if resp.StatusCode != 200 {
			t.Errorf("Response code is not ok: %v", resp.StatusCode)
			return
		}
		bb, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Errorf("Error on convert body: %v", err)
			return
		}
		v := struct {
			Ip string `json:"ip"`
		}{}
		if err = json.Unmarshal(bb, &v); err != nil {
			t.Errorf("Error on unmarshal body: %v", err)
			return
		}

		// Assert
		if v.Ip == conf.Host {
			t.Errorf("Proxy not excluded the host. Detected ip: %s", v.Ip)
			return
		}
	})

	t.Run("ok on err status code and no err returned", func(t *testing.T) {
		req := request.NewRequest().
			SetUrl("https://software.hixie.ch/utilities/cgi/test-tools/http-error?status=400+Bad+Request").
			AsGet()

		c := NewRestClient()
		resp, err := c.Send(req)
		if err != nil {
			t.Errorf("Error: %v", err)
			return
		}

		if resp.StatusCode != 400 {
			t.Errorf("Response code is not 400: %v", resp.StatusCode)
			return
		}
	})
}
