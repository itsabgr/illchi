package broker

import "errors"

var ErrBadContentLength = errors.New("bad content length")
var ErrConflictID = errors.New("conflict id")
