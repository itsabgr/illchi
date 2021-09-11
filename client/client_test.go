package client

import (
	"github.com/itsabgr/broker-go"
	"github.com/itsabgr/go-handy"
	"net"
	"testing"
)

func getFreeAddr() string {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	handy.Throw(err)

	l, err := net.ListenTCP("tcp", addr)
	handy.Throw(err)
	defer l.Close()
	return l.Addr().String()
}

func TestOverall(t *testing.T) {
	brokerAddr := getFreeAddr()
	aBroker, err := broker.New(broker.Config{
		Addr: brokerAddr,
	})
	handy.Throw(err)
	defer aBroker.Close()
	_, err = Dial(nil, Config{
		Add,
	})
	handy.Throw(err)
}
