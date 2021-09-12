package main

import (
	"github.com/itsabgr/illchi"
)

func main() {
	broker.New(broker.Config{
		Addr: ":80",
	})
}
