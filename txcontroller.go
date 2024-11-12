package gtcc

import "context"

type TxController struct {
	ctx        context.Context
	conf       *Config
	txStore    TxStore
	txRegister *TxRegister
}

func NewTxController(txStore TxStore) *TxController {
	ctx, _ := context.WithCancel(context.Background())
	txc := &TxController{
		ctx:        ctx,
		conf:       newConfig(),
		txStore:    txStore,
		txRegister: newRegister(),
	}
	go txc.run()
	return txc
}

func (c *TxController) run() {

}
