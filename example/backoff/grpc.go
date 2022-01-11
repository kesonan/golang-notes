// Package backoff from grpc
package backoff

import (
	"math/rand"
	"time"
)

var _ Backoff = (*GRPC)(nil)

type GRPCConfig struct {
	InitBackoff time.Duration
	Multiplier  float64
	Jitter      float64
	MaxBackoff  time.Duration
	InitRetries uint
	retries     uint
}

type GRPC struct {
	config GRPCConfig
	r      *rand.Rand
}

func NewGRPCBackoff(config GRPCConfig) Backoff {
	config.retries = config.InitRetries
	return &GRPC{
		config: config,
		r:      rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (bc *GRPC) TryGet() (time.Duration, uint, error) {
	if bc.config.retries <= 0 {
		return 0, bc.config.retries, OutOfRetries
	}
	backoff, max := float64(bc.config.InitBackoff), float64(bc.config.MaxBackoff)
	if backoff < max && bc.config.retries > 0 {
		backoff *= bc.config.Multiplier
		bc.config.retries--
	}
	if backoff > max {
		backoff = max
	}
	backoff *= 1 + bc.config.Jitter*(bc.r.Float64()*2-1)
	if backoff < 0 {
		backoff = 0
	}
	return time.Duration(backoff), bc.config.retries, nil
}

func (bc *GRPC) Reset() {
	bc.r.Seed(time.Now().UnixNano())
	bc.config.retries = bc.config.InitRetries
}
