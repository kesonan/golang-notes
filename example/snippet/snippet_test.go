package snippet

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Foo struct {
	member int
}

func (f *Foo) Get() int {
	return rand.Int()
}

func TestAnyCase(t *testing.T) {
	t.Run("CallFunctionByNil", func(t *testing.T) {
		var value atomic.Value
		value.Store((*Foo)(nil))
		ret := value.Load().(*Foo)
		assert.Nil(t, ret)
		assert.True(t, ret.Get() >= 0)
	})

	t.Run("atomic.CompareAndSwapInt32", func(t *testing.T) {
		var cur int32
		assert.False(t, atomic.CompareAndSwapInt32(&cur, 1, 1))
		atomic.StoreInt32(&cur, 1)
		assert.True(t, atomic.CompareAndSwapInt32(&cur, 1, 1))
	})

	t.Run("channel", func(t *testing.T) {
		ch := make(chan chan int, 0)
		go func() {
			ch <- make(chan int)
		}()

		select {
		case v := <-ch:
			fmt.Println(v)
		}
	})
}
