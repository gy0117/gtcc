package gtcc

import "time"

type Config struct {
	// transaction execution timeout
	TxTimeout time.Duration
	// polling transaction execution status
	MonitorInterval time.Duration
}
