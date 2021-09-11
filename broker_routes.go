package broker

import (
	"github.com/fasthttp/websocket"
	"github.com/valyala/fasthttp"
)

func (b *brokerImpl) routeCORS(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.Set("Access-Control-Allow-Origin", b.config.Origin)
	ctx.Response.Header.Set("Access-Control-Allow-Methods", accessControlAllowMethodsString)
	ctx.Response.Header.Set("Access-Control-Max-Age", accessControlMaxAgeString)
	ctx.SetStatusCode(fasthttp.StatusNoContent)
	ctx.SetConnectionClose()
}
func (b *brokerImpl) routeSend(ctx *fasthttp.RequestCtx) {
	messageLength := ctx.Request.Header.ContentLength()
	if messageLength <= 0 {
		ctx.SetStatusCode(fasthttp.StatusLengthRequired)
		ctx.SetBodyString(ErrBadContentLength.Error())
		ctx.SetConnectionClose()
		return
	}
	//if uint(messageLength) > b.config.MaxMessageSize {
	//	ctx.SetStatusCode(fasthttp.StatusRequestEntityTooLarge)
	//	ctx.SetBodyString(fasthttp.ErrBodyTooLarge.Error())
	//	ctx.SetConnectionClose()
	//	return
	//}
	targetID, err := parseIDFromHttpRequest(ctx)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.SetBodyString(err.Error())
		ctx.SetConnectionClose()
		return
	}
	targetWsConn := b.wsConnMap.Get(uintptr(targetID))
	if targetWsConn == nil {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		ctx.SetBodyString("Not Found")
		ctx.SetConnectionClose()
		return
	}
	err = b.authenticate(ctx)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusForbidden)
		ctx.SetBodyString(err.Error())
		ctx.SetConnectionClose()
		return
	}
	body := ctx.Request.Body()
	err = targetWsConn.WriteMessage(websocket.BinaryMessage, body)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadGateway)
		ctx.SetBodyString(err.Error())
		ctx.SetConnectionClose()
		return
	}
	ctx.ResetBody()
	ctx.SetStatusCode(fasthttp.StatusNoContent)
}
func (b *brokerImpl) routeUpgrade(ctx *fasthttp.RequestCtx) {
	desiredID, err := parseIDFromHttpRequest(ctx)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.SetBodyString(err.Error())
		ctx.SetConnectionClose()
		return
	}
	if conn := b.wsConnMap.Get(uintptr(desiredID)); conn != nil {
		ctx.SetStatusCode(fasthttp.StatusConflict)
		ctx.SetBodyString(ErrConflictID.Error())
		ctx.SetConnectionClose()
		return
	}
	if err := b.authenticate(ctx); err != nil {
		ctx.SetStatusCode(fasthttp.StatusForbidden)
		ctx.SetBodyString(err.Error())
		ctx.SetConnectionClose()
	}
	err = b.upgrader.Upgrade(ctx, func(conn *websocket.Conn) {
		defer conn.Close()
		if !b.wsConnMap.Add(uintptr(desiredID), conn) {
			return
		}
		defer b.wsConnMap.Delete(uintptr(desiredID))
		conn.SetReadLimit(1)
		conn.NextReader()
	})
	if err != nil {
		ctx.SetBodyString(err.Error())
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.SetConnectionClose()
	}
}
func (b *brokerImpl) routeStatics(ctx *fasthttp.RequestCtx) {
	ctx.SetBody(b.Stat().JSON())
}
