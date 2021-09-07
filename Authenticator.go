package broker

import (
	"github.com/valyala/fasthttp"
)

type Authenticator func(Broker, *fasthttp.RequestCtx) error
