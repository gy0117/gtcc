package gtcc

import "sync"

// TxRegister 注册中心
type TxRegister struct {
	sync.RWMutex
	serverMap map[string]TxServer
}

func newRegister() *TxRegister {
	return &TxRegister{
		serverMap: make(map[string]TxServer),
	}
}

// 1. 注册
func (txr *TxRegister) register(server TxServer) error {
	txr.RLock()
	_, ok := txr.serverMap[server.GetId()]
	if ok {
		txr.RUnlock()
		return ErrRepeatRegister
	}

	txr.Lock()
	defer txr.Unlock()
	_, ok = txr.serverMap[server.GetId()]
	if !ok {
		txr.serverMap[server.GetId()] = server
	}
	return nil
}

// 2. 获取服务
func (txr *TxRegister) getServers(sids ...string) ([]TxServer, error) {
	res := make([]TxServer, len(sids))

	txr.RLock()
	defer txr.RUnlock()

	for _, id := range sids {
		server, ok := txr.serverMap[id]
		if !ok {
			return nil, ErrServerNotFound
		}
		res = append(res, server)
	}
	return res, nil
}
