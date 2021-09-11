package client

import (
	"fmt"
	"github.com/itsabgr/broker-go"
	"github.com/itsabgr/go-handy"
	"github.com/rocketlaunchr/https-go"
	"net"
	"sync"
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
	cert, key, err := https.GenerateKeys(https.GenerateOptions{
		Host: "127.0.0.1",
		IsCA: true,
	})
	handy.Throw(err)
	aBroker, err := broker.New(broker.Config{
		Addr: brokerAddr,
		Cert: cert,
		Key:  key,
	})
	if err != nil {
		t.Fatal(err)
	}

	defer aBroker.Close()
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func() {
		wg.Done()
		fmt.Printf("listen %s\n", brokerAddr)
		handy.Throw(aBroker.Listen())
	}()
	wg.Wait()
	client1, err := Dial(nil, Config{
		Broker:        brokerAddr,
		ID:            1,
		SkipVerifySSL: true,
	})
	if err != nil {
		t.Fatal(err)
	}
	defer client1.Close()
	client2, err := Dial(nil, Config{
		Broker:        brokerAddr,
		ID:            2,
		SkipVerifySSL: true,
	})
	if err != nil {
		t.Fatal(err)
	}
	defer client2.Close()
	err = client1.Send(nil, brokerAddr, 2, []byte("hello"))
	if err != nil {
		t.Fatal(err)
	}
	msg, err := client2.Receive(nil)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(msg)

}
