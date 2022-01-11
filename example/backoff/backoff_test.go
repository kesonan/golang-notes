package backoff

import (
	"fmt"
	"testing"
	"time"
)

func TestNewBackoff(t *testing.T) {
	initConfig := Config{
		InitBackoff: time.Second,
		Multiplier:  1.2,
		Jitter:      0.2,
		MaxBackoff:  10 * time.Second,
		InitRetries: 5,
	}

	bc := NewBackoff(initConfig)
	for idx := 0; idx < 5; idx++ {
		fmt.Println("-------------")
		for {
			dur, retries, err := bc.TryGet()
			if err != nil {
				break
			}
			fmt.Println(retries, ":", dur)
		}
		bc.Reset()
	}
}
