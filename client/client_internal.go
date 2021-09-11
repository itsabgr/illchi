package client

import (
	"github.com/gorilla/websocket"
	"net/http"
)

type client struct {
	ws   websocket.Conn
	http http.Client
}
