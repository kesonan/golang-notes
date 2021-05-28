//栈排序。 编写程序，对栈进行排序使最小元素位于栈顶。最多只能使用一个其他的临时栈存放数据，但不得将元素复制到别的数据结构（如数组）中。该栈支持如下操作：pu
//sh、pop、peek 和 isEmpty。当栈为空时，peek 返回 -1。
//
// 示例1:
//
//  输入：
//["SortedStack", "push", "push", "peek", "pop", "peek"]
//[[], [1], [2], [], [], []]
// 输出：
//[null,null,null,1,null,2]
//
//
// 示例2:
//
//  输入：
//["SortedStack", "pop", "pop", "push", "pop", "isEmpty"]
//[[], [], [], [1], [], []]
// 输出：
//[null,null,null,null,null,true]
//
//
// 说明:
//
//
// 栈中的元素数目在[0, 5000]范围内。
//
// Related Topics 设计
// 👍 36 👎 0

package leetcode

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortedStack(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		s := ConstructorSortedStack()
		s.Push(1)
		s.Push(2)
		assert.Equal(t, 1, s.Peek())
		s.Pop()
		assert.Equal(t, 2, s.Peek())
	})

	t.Run("case 2", func(t *testing.T) {
		s := ConstructorSortedStack()
		s.Pop()
		s.Pop()
		s.Push(1)
		s.Pop()
		assert.True(t, s.IsEmpty())
	})
}

type SortedStack struct {
	list []int
}

func ConstructorSortedStack() SortedStack {
	return SortedStack{}
}

func (this *SortedStack) Push(val int) {
	this.list = append(this.list, val)
	sort.Slice(this.list, func(i, j int) bool {
		return this.list[i] > this.list[j]
	})
}

func (this *SortedStack) Pop() {
	if len(this.list) == 0 {
		return
	}
	this.list = this.list[:len(this.list)-1]
}

func (this *SortedStack) Peek() int {
	if len(this.list) == 0 {
		return -1
	}
	return this.list[len(this.list)-1]
}

func (this *SortedStack) IsEmpty() bool {
	return len(this.list) == 0
}
