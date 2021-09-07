package broker

type Broker interface {
	Close() error
	Listen() error
	Stat() *Statics
}
