package broker

import (
	"github.com/fasthttp/websocket"
	"github.com/valyala/fasthttp"
	"time"
)

type brokerImpl struct {
	http      fasthttp.Server
	config    Config
	wsConnMap webSocketMap
	upgrader  websocket.FastHTTPUpgrader
	startTime int64
}

func (b *brokerImpl) init() {
	b.http.Handler = b.httpHandler
	if len(b.config.Origin) == 0 {
		b.config.Origin = "*"
	}
	b.startTime = time.Now().UTC().Unix()
}
func (b *brokerImpl) authenticate(ctx *fasthttp.RequestCtx) error {
	if b.config.Authenticator == nil {
		return nil
	}
	err := b.config.Authenticator(b, ctx)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusMethodNotAllowed)
		ctx.SetConnectionClose()
	}
	return err
}

func (b *brokerImpl) httpHandler(ctx *fasthttp.RequestCtx) {

	switch string(ctx.Method()) {
	case fasthttp.MethodGet:
		switch string(ctx.Path()) {
		case "/":
			b.routeStatics(ctx)
		default:
			b.routeUpgrade(ctx)
		}
	case fasthttp.MethodPost:
		b.routeSend(ctx)
	case fasthttp.MethodOptions:
		b.routeCORS(ctx)
	default:
		ctx.SetStatusCode(fasthttp.StatusMethodNotAllowed)
		ctx.SetConnectionClose()
	}
}
