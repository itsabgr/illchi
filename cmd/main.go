package main

import (
	"github.com/itsabgr/broker-go"
)

func main() {
	broker.New(broker.Config{
		Addr: ":80",
	})
}
