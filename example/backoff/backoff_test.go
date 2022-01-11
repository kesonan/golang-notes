package backoff

import (
	"fmt"
	"testing"
	"time"
)

func TestGRPCBackoff(t *testing.T) {
	initConfig := GRPCConfig{
		InitBackoff: time.Second,
		Multiplier:  1.2,
		Jitter:      0.2,
		MaxBackoff:  10 * time.Second,
		InitRetries: 5,
	}

	bc := NewGRPCBackoff(initConfig)
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

func TestWikiBackoff(t *testing.T) {
	initConfig := WikiConfig{
		InitBackoff: time.Second,
		Jitter:      0.2,
		MaxBackoff:  10 * time.Second,
		InitRetries: 10,
	}

	bc := NewWikiBackoff(initConfig)
	for idx := 0; idx < 2; idx++ {
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
