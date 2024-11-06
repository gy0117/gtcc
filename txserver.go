package gtcc

import "context"

type TxIn struct {
	ServerId string                 `json:"server_id"`
	TxId     string                 `json:"tx_id"`
	Data     map[string]interface{} `json:"data"`
}

type TxOut struct {
	ServerId string `json:"server_id"`
	TxId     string `json:"tx_id"`
	Ack      int8   `json:"ack"`
}

// TxServer Services participating in distributed transactions must implement this interface
type TxServer interface {
	// GetId get server unique id
	GetId() string
	// Try first phase
	Try(ctx context.Context, in *TxIn) (*TxOut, error)
	// Confirm second phase operation
	Confirm(ctx context.Context, txId string) (*TxOut, error)
	// Cancel second phase operation
	Cancel(ctx context.Context, txId string) (*TxOut, error)
}
