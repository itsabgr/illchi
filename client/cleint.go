package client

import (
	"context"
	"crypto/tls"
	"github.com/fasthttp/websocket"
	"net/url"
)

func Dial(ctx context.Context, config Config) (Client, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	url := url.URL{
		Scheme: "wss",
		Host:   config.Broker,
		Path:   "/" + config.ID,
	}
	dialer := websocket.Dialer{
		EnableCompression: false,
		TLSClientConfig:   config.getTlsConfig(),
	}
	dialer.DialContext(ctx)
}
