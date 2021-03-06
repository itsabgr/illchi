package client

import (
	"github.com/itsabgr/illchi"
	"net/url"
)

type Config struct {
	Broker string
	ID     broker.ID
	Params url.Values
	TLS    bool
	//Cert, Key []byte
	//MaxMessageSize uint
}

//func (c *Config) hasTLS() bool {
//	return c.Cert != nil && c.Key != nil && len(c.Cert) > 0 && len(c.Key) > 0
//}
//
//func (c *Config) getTlsConfig() *tls.Config {
//	cert, err := tls.X509KeyPair(c.Cert, c.Key)
//	handy.Throw(err)
//	return &tls.Config{
//		Certificates: []tls.Certificate{cert},
//	}
//}
