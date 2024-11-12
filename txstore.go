package gtcc

import "context"

// TxStore 事务表
type TxStore interface {
	// CreateTx 创建事务，返回全局事务id
	CreateTx(ctx context.Context, servers []TxServer) (gid string, err error)
	// UpdateTx 更新事务，更新各个服务的try操作结束
	UpdateTx(ctx context.Context, gid string, serverId string, ack bool) error
	// SubmitTx 提交事务
	SubmitTx(ctx context.Context, gid string, success bool) error
}
