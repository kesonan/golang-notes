package backoff

import (
	"math"
	"math/rand"
	"time"
)

type Wiki struct {
	config WikiConfig
	r      *rand.Rand
}

type WikiConfig struct {
	InitBackoff time.Duration
	Jitter      float64
	MaxBackoff  time.Duration
	InitRetries uint
	retries     uint
}

func NewWikiBackoff(config WikiConfig) Backoff {
	config.retries = config.InitRetries
	return &Wiki{
		config: config,
		r:      rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (bc *Wiki) TryGet() (time.Duration, uint, error) {
	if bc.config.retries <= 0 {
		return 0, bc.config.retries, OutOfRetries
	}
	backoff, max := float64(bc.config.InitBackoff), float64(bc.config.MaxBackoff)
	if backoff < max && bc.config.retries > 0 {
		c := float64(bc.config.InitRetries - bc.config.retries + 1)
		// (2^c -1)/2
		backoff *= (math.Exp2(c) - 1) / float64(2)
		bc.config.retries--
	}
	if backoff > max {
		backoff = max
	}
	// (-1,1)
	backoff *= 1 + bc.config.Jitter*(bc.r.Float64()*2-1)
	if backoff < 0 {
		backoff = 0
	}
	return time.Duration(backoff), bc.config.retries, nil
}

func (bc *Wiki) Reset() {
	bc.r.Seed(time.Now().UnixNano())
	bc.config.retries = bc.config.InitRetries
}
