// Package backoff from grpc
package backoff

import (
	"errors"
	"math/rand"
	"time"
)

var OutOfRetries = errors.New("out of retries")

type Config struct {
	InitBackoff time.Duration
	Multiplier  float32
	Jitter      float32
	MaxBackoff  time.Duration
	InitRetries uint
	retries     uint
}

type Backoff struct {
	config Config
	r      *rand.Rand
}

func NewBackoff(config Config) *Backoff {
	config.retries = config.InitRetries
	return &Backoff{
		config: config,
		r:      rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (bc *Backoff) TryGet() (time.Duration, uint, error) {
	if bc.config.retries <= 0 {
		return 0, bc.config.retries, OutOfRetries
	}
	backoff, max := float32(bc.config.InitBackoff), float32(bc.config.MaxBackoff)
	if backoff < max && bc.config.retries > 0 {
		backoff *= bc.config.Multiplier
		bc.config.retries--
	}
	if backoff > max {
		backoff = max
	}
	backoff *= 1 + bc.config.Jitter*(rand.Float32()*2-1)
	if backoff < 0 {
		backoff = 0
	}
	return time.Duration(backoff), bc.config.retries, nil
}

func (bc *Backoff) Reset() {
	bc.r.Seed(time.Now().UnixNano())
	bc.config.retries = bc.config.InitRetries
}
