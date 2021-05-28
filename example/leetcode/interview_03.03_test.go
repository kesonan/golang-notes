//堆盘子。设想有一堆盘子，堆太高可能会倒下来。因此，在现实生活中，盘子堆到一定高度时，我们就会另外堆一堆盘子。请实现数据结构SetOfStacks，模拟这种行
//为。SetOfStacks应该由多个栈组成，并且在前一个栈填满时新建一个栈。此外，SetOfStacks.push()和SetOfStacks.pop()应该与
//普通栈的操作方法相同（也就是说，pop()返回的值，应该跟只有一个栈时的情况一样）。 进阶：实现一个popAt(int index)方法，根据指定的子栈，执行p
//op操作。
//
// 当某个栈为空时，应当删除该栈。当栈中没有元素或不存在该栈时，pop，popAt 应返回 -1.
//
// 示例1:
//
//  输入：
//["StackOfPlates", "push", "push", "popAt", "pop", "pop"]
//[[1], [1], [2], [1], [], []]
// 输出：
//[null, null, null, 2, 1, -1]
//
//
// 示例2:
//
//  输入：
//["StackOfPlates", "push", "push", "push", "popAt", "popAt", "popAt"]
//[[2], [1], [2], [3], [0], [0], [0]]
// 输出：
//[null, null, null, null, 2, 1, 3]
//
// Related Topics 设计
// 👍 20 👎 0

package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStackOfPlates(t *testing.T) {
	t.Run("size 1", func(t *testing.T) {
		s := ConstructorStackOfPlates(1)
		s.Push(1)
		s.Push(2)
		assert.Equal(t, 2, s.PopAt(1))
		assert.Equal(t, 1, s.Pop())
		assert.Equal(t, -1, s.Pop())
	})
	t.Run("size 2", func(t *testing.T) {
		s := ConstructorStackOfPlates(2)
		s.Push(1)
		s.Push(2)
		s.Push(3)
		assert.Equal(t, 2, s.PopAt(0))
		assert.Equal(t, 1, s.PopAt(0))
		assert.Equal(t, 3, s.PopAt(0))
	})

}

type StackOfPlates struct {
	list [][]int
	cap  int
}

func ConstructorStackOfPlates(cap int) StackOfPlates {
	return StackOfPlates{
		cap: cap,
	}
}

func (this *StackOfPlates) Push(val int) {
	if this.cap == 0 {
		return
	}

	var tail []int
	if len(this.list) > 0 {
		tail = this.list[len(this.list)-1]
	} else {
		this.list = append(this.list, tail)
	}

	if len(tail) == this.cap {
		tail = []int{val}
		this.list = append(this.list, tail)
	} else if len(tail) < this.cap {
		tail = append(tail, val)
		this.list[len(this.list)-1] = tail
	} else {
		return
	}
}

func (this *StackOfPlates) Pop() int {
	if this.cap == 0 {
		return -1
	}

	if len(this.list) == 0 {
		return -1
	}
	tail := this.list[len(this.list)-1]
	if len(tail) == 0 {
		return -1
	}

	value := tail[len(tail)-1]
	if len(tail) == 1 {
		this.list = this.list[:len(this.list)-1]
	}
	if len(tail) > 1 {
		this.list[len(this.list)-1] = tail[:len(tail)-1]
	}
	return value
}

func (this *StackOfPlates) PopAt(index int) int {
	if this.cap == 0 {
		return -1
	}

	if index >= len(this.list) {
		return -1
	}

	tail := this.list[index]
	if len(tail) == 0 {
		return -1
	}

	value := tail[len(tail)-1]
	if len(tail) == 1 {
		var remain [][]int
		if index+1 < len(this.list) {
			remain = this.list[index+1:]
		}
		this.list = append([][]int(nil), this.list[:index]...)
		this.list = append(this.list, remain...)
	}

	if len(tail) > 1 {
		this.list[index] = tail[:len(tail)-1]
	}
	return value
}
