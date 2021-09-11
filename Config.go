package broker

import (
	"crypto/tls"
	"github.com/itsabgr/go-handy"
	"github.com/rocketlaunchr/https-go"
)

type Config struct {
	Addr          string
	Origin        string
	Authenticator Authenticator
	Cert, Key     []byte
	//StartTime      int64
	//MaxMessageSize uint
}

func (c *Config) getTlsConfig() *tls.Config {
	if c.Key == nil && c.Cert == nil {
		var err error
		c.Cert, c.Key, err = https.GenerateKeys(https.GenerateOptions{
			Host: c.Origin,
		})
		handy.Throw(err)
	}
	cert, err := tls.X509KeyPair(c.Cert, c.Key)
	handy.Throw(err)
	return &tls.Config{
		Certificates: []tls.Certificate{cert},
	}
}
