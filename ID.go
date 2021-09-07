package broker

import (
	"github.com/valyala/fasthttp"
	"strconv"
)

type ID uint32

func (id ID) String() string {
	return strconv.FormatUint(uint64(id), 10)
}
func parseIDFromHttpRequest(r *fasthttp.RequestCtx) (ID, error) {
	path := r.Path()[1:]
	id, err := strconv.ParseUint(string(path), 10, 32)
	return ID(id), err
}
