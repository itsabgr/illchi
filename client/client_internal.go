package client

import (
	"context"
	"errors"
	"github.com/fasthttp/websocket"
	"github.com/itsabgr/broker-go"
	"github.com/valyala/fasthttp"
	"net/url"
	"os"
)

type client struct {
	ws     *websocket.Conn
	http   fasthttp.Client
	config Config
}

func (c *client) Receive(_ context.Context) (msg []byte, err error) {
begin:
	kind, msg, err := c.ws.ReadMessage()
	if err != nil {
		return nil, err
	}
	switch kind {
	case websocket.PingMessage, websocket.PongMessage:
		goto begin
	case websocket.CloseMessage:
		c.Close()
		return nil, os.ErrClosed
	}
	return msg, nil
}

func (c *client) Send(_ context.Context, host string, id broker.ID, msg []byte) error {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	url := url.URL{
		Scheme: "http",
		Host:   host,
		Path:   "/" + id.String(),
	}
	if c.config.TLS {
		url.Scheme = "https"
	}
	req.Header.SetMethod(fasthttp.MethodPost)
	req.SetRequestURI(url.String())
	req.SetBody(msg)
	res := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(res)
	err := c.http.Do(req, res)
	if err != nil {
		return err
	}
	switch res.StatusCode() {
	case 200, 204:
		return nil
	}
	return errors.New(string(res.Body()))
}

func (c *client) Close() error {
	return c.ws.Close()
}
