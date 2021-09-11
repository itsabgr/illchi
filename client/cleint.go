package client

import (
	"context"
	"github.com/fasthttp/websocket"
	"github.com/valyala/fasthttp"
	"net/url"
	"time"
)

func Dial(ctx context.Context, config Config) (Client, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	url := url.URL{
		Scheme: "ws",
		Host:   config.Broker,
		Path:   "/" + config.ID.String(),
	}
	if config.TLS {
		url.Scheme = "wss"
	}
	if config.Params != nil {
		url.RawQuery = config.Params.Encode()
	}
	dialer := websocket.Dialer{
		HandshakeTimeout:  2 * time.Second,
		EnableCompression: false,
	}
	conn, _, err := dialer.DialContext(ctx, url.String(), nil)
	if err != nil {
		return nil, err
	}

	cli := &client{
		ws: conn,
		http: fasthttp.Client{
			Name:                "go-client/0.1.0",
			MaxConnsPerHost:     1,
			MaxConnWaitTimeout:  2 * time.Second,
			MaxConnDuration:     10 * time.Second,
			MaxIdleConnDuration: 2 * time.Second,
			WriteTimeout:        10 * time.Second,
			ReadTimeout:         10 * time.Second,
		},
	}
	//if config.hasTlsOptions() {
	//	cli.http.
	//}
	return cli, nil
}
