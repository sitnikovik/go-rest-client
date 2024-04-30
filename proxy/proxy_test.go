package proxy

import (
	"fmt"
	"github.com/sitnikovik/go-rest-client/internal/config"
	"testing"
)

// TestProxy_Dsn tests the Dsn method
func TestProxy_Dsn(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		conf := config.FromFile()

		p := NewProxy().
			SetHost(conf.Proxy.Host).
			SetPort(conf.Proxy.Port).
			SetLogin(conf.Proxy.Login).
			SetPassword(conf.Proxy.Password)

		actual := p.Dsn()
		expected := fmt.Sprintf(
			"http://%s:%s@%s:%d",
			conf.Proxy.Login,
			conf.Proxy.Password,
			conf.Proxy.Host,
			conf.Proxy.Port,
		)

		if actual != expected {
			t.Errorf("DSN (%s) is not equals expected: %s)", actual, expected)
		}
	})
}

// TestProxy_GetHost tests the GetHost method
func TestProxy_GetHost(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		conf := config.FromFile()

		p := NewProxy().
			SetHost(conf.Proxy.Host)

		actual := p.GetHost()
		expected := conf.Proxy.Host

		if actual != expected {
			t.Errorf("Host (%s) is not equals expected: %s)", actual, expected)
		}
	})
}

// TestProxy_SetHost tests the SetHost method
func TestProxy_SetHost(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		conf := config.FromFile()

		p := NewProxy().
			SetHost(conf.Proxy.Host)

		actual := p.GetHost()
		expected := conf.Proxy.Host

		if actual != expected {
			t.Errorf("Host (%s) is not equals expected: %s)", actual, expected)
		}
	})
}

// TestProxy_GetPort tests the GetPort method
func TestProxy_GetPort(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		conf := config.FromFile()

		p := NewProxy().
			SetPort(conf.Proxy.Port)

		actual := p.GetPort()
		expected := conf.Proxy.Port

		if actual != expected {
			t.Errorf("Port (%d) is not equals expected: %d)", actual, expected)
		}
	})
}

// TestProxy_SetPort tests the SetPort method
func TestProxy_SetPort(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		conf := config.FromFile()

		p := NewProxy().
			SetPort(conf.Proxy.Port)

		actual := p.GetPort()
		expected := conf.Proxy.Port

		if actual != expected {
			t.Errorf("Port (%d) is not equals expected: %d)", actual, expected)
		}
	})
}

// TestProxy_GetLogin tests the GetLogin method
func TestProxy_GetLogin(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		conf := config.FromFile()

		p := NewProxy().
			SetLogin(conf.Proxy.Login)

		actual := p.GetLogin()
		expected := conf.Proxy.Login

		if actual != expected {
			t.Errorf("Login (%s) is not equals expected: %s)", actual, expected)
		}
	})
}

// TestProxy_SetLogin tests the SetLogin method
func TestProxy_SetLogin(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		conf := config.FromFile()

		p := NewProxy().
			SetLogin(conf.Proxy.Login)

		actual := p.GetLogin()
		expected := conf.Proxy.Login

		if actual != expected {
			t.Errorf("Login (%s) is not equals expected: %s)", actual, expected)
		}
	})
}

// TestProxy_GetPassword tests the GetPassword method
func TestProxy_GetPassword(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		conf := config.FromFile()

		p := NewProxy().
			SetPassword(conf.Proxy.Password)

		actual := p.GetPassword()
		expected := conf.Proxy.Password

		if actual != expected {
			t.Errorf("Password (%s) is not equals expected: %s)", actual, expected)
		}
	})
}

// TestProxy_SetPassword tests the SetPassword method
func TestProxy_SetPassword(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		conf := config.FromFile()

		p := NewProxy().
			SetPassword(conf.Proxy.Password)

		actual := p.GetPassword()
		expected := conf.Proxy.Password

		if actual != expected {
			t.Errorf("Password (%s) is not equals expected: %s)", actual, expected)
		}
	})
}

// TestProxy_NewProxy tests the NewProxy function
func TestProxy_NewProxy(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		p := NewProxy()

		if p == nil {
			t.Errorf("Proxy is nil")
		}
		if p.GetHost() != "" {
			t.Errorf("Host is not empty")
		}
		if p.GetPort() != 0 {
			t.Errorf("Port is not 0")
		}
		if p.GetLogin() != "" {
			t.Errorf("Login is not empty")
		}
		if p.GetPassword() != "" {
			t.Errorf("Password is not empty")
		}
		if p.GetExcludedHosts() == nil {
			t.Errorf("Excluded hosts is nil")
		}
		if len(p.GetExcludedHosts()) != 0 {
			t.Errorf("Excluded hosts is not empty")
		}
	})
}

// TestProxy_GetExcludedHosts tests the GetExcludedHosts method
func TestProxy_GetExcludedHosts(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		p := NewProxy().
			SetExcludedHosts([]string{"localhost"})

		actual := p.GetExcludedHosts()
		expected := []string{"localhost"}

		if len(actual) != len(expected) {
			t.Errorf("Excluded hosts length (%d) is not equals expected: %d)", len(actual), len(expected))
		}
		for i := range actual {
			if actual[i] != expected[i] {
				t.Errorf("Excluded host (%s) is not equals expected: %s)", actual[i], expected[i])
			}
		}
	})
}

// TestProxy_SetExcludedHosts tests the SetExcludedHosts method
func TestProxy_SetExcludedHosts(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		p := NewProxy().
			SetExcludedHosts([]string{"localhost"})

		actual := p.GetExcludedHosts()
		expected := []string{"localhost"}

		if len(actual) != len(expected) {
			t.Errorf("Excluded hosts length (%d) is not equals expected: %d)", len(actual), len(expected))
		}
		for i := range actual {
			if actual[i] != expected[i] {
				t.Errorf("Excluded host (%s) is not equals expected: %s)", actual[i], expected[i])
			}
		}
	})
}

// TestProxy_AddExcludedHost tests the AddExcludedHost method
func TestProxy_AddExcludedHost(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		p := NewProxy().
			SetExcludedHosts([]string{"localhost"}).
			AddExcludedHost("localhost1")

		actual := p.GetExcludedHosts()
		expected := []string{"localhost", "localhost1"}

		if len(actual) != len(expected) {
			t.Errorf("Excluded hosts length (%d) is not equals expected: %d)", len(actual), len(expected))
		}
		for i := range actual {
			if actual[i] != expected[i] {
				t.Errorf("Excluded host (%s) is not equals expected: %s)", actual[i], expected[i])
			}
		}
	})
}

// TestProxy_HasExcludedHost tests the HasExcludedHost method
func TestProxy_HasExcludedHost(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		p := NewProxy().
			SetExcludedHosts([]string{"localhost"})

		actual := p.HasExcludedHost("localhost")

		if actual != true {
			t.Errorf("Excluded host (%t) is not equals expected: %t)", actual, true)
		}
	})
}
