package broker

import (
	"crypto/tls"
	"github.com/itsabgr/go-handy"
)

type Config struct {
	Addr           string
	Origin         string
	Authenticator  Authenticator
	Cert, Key      []byte
	StartTime      int64
	MaxMessageSize uint
}

func (c *Config) hasTlsOptions() bool {
	return c.Cert != nil && c.Key != nil && len(c.Cert) > 0 && len(c.Key) > 0
}

func (c *Config) getTlsConfig() *tls.Config {
	cert, err := tls.X509KeyPair(c.Cert, c.Key)
	handy.Throw(err)
	return &tls.Config{
		Certificates: []tls.Certificate{cert},
	}
}
