package gtcc

import "context"

type TxController struct {
	ctx        context.Context
	conf       *Config
	txStore    TxStore
	txRegister TxRegister
}
