//请设计一个栈，除了常规栈支持的pop与push函数以外，还支持min函数，该函数返回栈元素中的最小值。执行push、pop和min操作的时间复杂度必须为O(
//1)。 示例：
// MinStack minStack = new MinStack();
// minStack.push(-2);
// minStack.push(0);
// minStack.push(-3);
// minStack.getMin();   --> 返回 -3.
// minStack.pop();
// minStack.top ();      --> 返回 0.
// minStack.getMin();   --> 返回 -2.
// Related Topics 栈
// 👍 48 👎 0

package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinStack(t *testing.T) {
	ms := Constructor()
	ms.Push(-2)
	ms.Push(0)
	ms.Push(-3)
	assert.Equal(t, -3, ms.GetMin())
	ms.Pop()
	assert.Equal(t, 0, ms.Top())
	assert.Equal(t, -2, ms.GetMin())
}

type MinStack struct {
	list []int
	min  []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	this.list = append(this.list, x)
	if len(this.min) > 0 {
		top := this.min[len(this.min)-1]
		if x < top {
			this.min = append(this.min, x)
		} else {
			this.min = append(this.min, top)
		}
	} else {
		this.min = append(this.min, x)
	}
}

func (this *MinStack) Pop() {
	if len(this.list) == 0 {
		return
	}
	this.list = this.list[:len(this.list)-1]
	this.min = this.min[:len(this.min)-1]
}

func (this *MinStack) Top() int {
	if len(this.list) == 0 {
		return 0
	}

	return this.list[len(this.list)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.min) == 0 {
		return 0
	}
	return this.min[len(this.min)-1]
}
