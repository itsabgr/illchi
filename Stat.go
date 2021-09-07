package broker

import (
	"encoding/json"
	"sync"
)
import "github.com/itsabgr/go-handy"

//Statics is Server status and Statics
type Statics struct {
	UpTime      int32  `json:"uptime"`
	Connections int32  `json:"conns"`
	Version     string `json:"v"`
}

func (s *Statics) JSON() []byte {
	b, err := json.Marshal(s)
	handy.Throw(err)
	return b
}

var poolStructStatics = sync.Pool{
	New: func() handy.Any {
		return &Statics{
			Version: Version,
		}
	},
}
