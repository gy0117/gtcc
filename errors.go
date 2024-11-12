package gtcc

import "errors"

var ErrRepeatRegister = errors.New("duplicate register error")

var ErrServerNotFound = errors.New("server not found error")
