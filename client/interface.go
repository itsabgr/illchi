package client

import (
	"context"
	"github.com/itsabgr/illchi"
)

type Client interface {
	Receive(ctx context.Context) (msg []byte, err error)
	Send(ctx context.Context, Host string, ID broker.ID, msg []byte) (err error)
	Close() error
}
