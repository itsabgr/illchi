package broker

import (
	"crypto/tls"
	"github.com/itsabgr/go-handy"
	"github.com/kpango/fastime"
	"net"
)

func New(config Config) (broker Broker, err error) {
	handy.Catch(func(recovered interface{}) {
		err = recovered.(error)
	})
	b := new(brokerImpl)
	b.config = config
	b.init()
	return b, nil
}

func (b *brokerImpl) Close() error {
	return b.http.Shutdown()
}

func (b *brokerImpl) Listen() error {
	listener, err := net.Listen("tcp", b.config.Addr)
	if err != nil {
		return err
	}
	defer handy.Close(listener)
	if b.config.hasTlsOptions() {
		listener = tls.NewListener(listener, b.config.getTlsConfig())
	}
	return b.http.Serve(listener)
}

func (b *brokerImpl) Stat() *Statics {
	stat := poolStructStatics.Get().(*Statics)
	stat.UpTime = int32(b.startTime - fastime.UnixNow())
	stat.Connections = b.http.GetOpenConnectionsCount()
	return stat
}
