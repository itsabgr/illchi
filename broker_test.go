package broker

import (
	"crypto/tls"
	"fmt"
	"github.com/fasthttp/websocket"
	"github.com/itsabgr/go-handy"
	"net"
	"sync"
	"testing"
	"time"
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
	aBroker, err := New(Config{
		Addr: brokerAddr,
	})
	if err != nil {
		t.Fatal(err)
	}
	//defer aBroker.Close()
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func(t *testing.T) {
		fmt.Printf("listen %s\n", brokerAddr)
		wg.Done()
		t.Fatal(aBroker.Listen())
	}(t)
	wg.Wait()
	websocket.DefaultDialer.HandshakeTimeout = 1 * time.Second
	websocket.DefaultDialer.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	conn, resp, err := websocket.DefaultDialer.Dial(fmt.Sprintf("ws://%s/100?foo=bad", brokerAddr), nil)
	if err != nil {
		t.Log(resp)
		t.Fatal(err)
	}
	defer conn.Close()
	t.Log("connected")
}
