package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyQueue(t *testing.T) {
	queue := Constructor0304()
	queue.Push(1)
	queue.Push(2)
	assert.Equal(t, 1, queue.Peek())
	assert.Equal(t, 1, queue.Pop())
	assert.False(t, queue.Empty())
}

type MyQueue struct {
	list []int
}

/** Initialize your data structure here. */
func Constructor0304() MyQueue {
	return MyQueue{}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.list = append(this.list, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	peek := this.Peek()
	n := len(this.list)
	if n > 0 {
		this.list = this.list[1:]
	}
	return peek
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if len(this.list) == 0 {
		return 0
	}

	n := len(this.list)
	var popElement int
	if n > 0 {
		popElement = this.list[0]
	}
	return popElement
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.list) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
