package proxy

import "fmt"

// Proxy is the interface that wraps the basic methods of a proxy
type Proxy interface {
	// Dsn returns the DSN of the proxy
	Dsn() string

	// GetHost returns the host of the proxy
	GetHost() string
	// SetHost sets the host of the proxy
	SetHost(host string) Proxy

	// GetPort returns the port of the proxy
	GetPort() int
	// SetPort sets the port of the proxy
	SetPort(port int) Proxy

	// GetLogin returns the login of the proxy
	GetLogin() string
	// SetLogin sets the login of the proxy
	SetLogin(login string) Proxy

	// GetPassword returns the password of the proxy
	GetPassword() string
	// SetPassword sets the password of the proxy
	SetPassword(password string) Proxy

	// GetExcludedHosts returns the excluded hosts of the proxy
	GetExcludedHosts() []string
	// SetExcludedHosts sets the excluded hosts of the proxy
	SetExcludedHosts(hosts []string) Proxy
	// AddExcludedHost adds an excluded host to the proxy
	AddExcludedHost(host string) Proxy
	// HasExcludedHost returns true if the host is excluded
	HasExcludedHost(host string) bool
}

// proxy is a basic implementation of the Proxy interface
type proxy struct {
	host          string
	port          int
	login         string
	password      string
	excludedHosts []string
}

// NewProxy returns a new instance of the Proxy interface
func NewProxy() Proxy {
	return &proxy{
		excludedHosts: make([]string, 0),
	}
}

// Dsn returns the DSN of the proxy
func (p *proxy) Dsn() string {
	return fmt.Sprintf("http://%s:%s@%s:%d", p.login, p.password, p.host, p.port)
}

// GetHost returns the host of the proxy
func (p *proxy) GetHost() string {
	return p.host
}

// SetHost sets the host of the proxy
func (p *proxy) SetHost(host string) Proxy {
	p.host = host

	return p
}

// GetPort returns the port of the proxy
func (p *proxy) GetPort() int {
	return p.port
}

// SetPort sets the port of the proxy
func (p *proxy) SetPort(port int) Proxy {
	p.port = port

	return p
}

// GetLogin returns the login of the proxy
func (p *proxy) GetLogin() string {
	return p.login
}

// SetLogin sets the login of the proxy
func (p *proxy) SetLogin(login string) Proxy {
	p.login = login

	return p
}

// GetPassword returns the password of the proxy
func (p *proxy) GetPassword() string {
	return p.password
}

// SetPassword sets the password of the proxy
func (p *proxy) SetPassword(password string) Proxy {
	p.password = password

	return p
}

// GetExcludedHosts returns the excluded hosts of the proxy
func (p *proxy) GetExcludedHosts() []string {
	return p.excludedHosts
}

// SetExcludedHosts sets the excluded hosts of the proxy
func (p *proxy) SetExcludedHosts(hosts []string) Proxy {
	p.excludedHosts = hosts

	return p
}

// AddExcludedHost adds an excluded host to the proxy
func (p *proxy) AddExcludedHost(host string) Proxy {
	p.excludedHosts = append(p.excludedHosts, host)

	return p
}

// HasExcludedHost returns true if the host is excluded
func (p *proxy) HasExcludedHost(host string) bool {
	for _, h := range p.excludedHosts {
		if h == host {
			return true
		}
	}

	return false
}
