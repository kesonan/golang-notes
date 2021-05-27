//实现一个MyQueue类，该类用两个栈来实现一个队列。 示例： MyQueue queue = new MyQueue(); queue.push(1);
//queue.push(2); queue.peek();  // 返回 1 queue.pop();   // 返回 1 queue.empty(); // 返
//回 false 说明： 你只能使用标准的栈操作 -- 也就是只有 push to top, peek/pop from top, size 和 is empty
// 操作是合法的。 你所使用的语言也许不支持栈。你可以使用 list 或者 deque（双端队列）来模拟一个栈，只要是标准的栈操作即可。 假设所有操作都是有效的
//（例如，一个空的队列不会调用 pop 或者 peek 操作）。 Related Topics 栈
// 👍 39 👎 0

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
