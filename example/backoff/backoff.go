package backoff

import (
	"errors"
	"time"
)

var OutOfRetries = errors.New("out of retries")

type Backoff interface {
	TryGet() (duration time.Duration, remainRetryCount uint, err error)
	Reset()
}
