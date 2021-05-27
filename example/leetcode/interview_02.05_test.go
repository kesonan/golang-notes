//给定两个用链表表示的整数，每个节点包含一个数位。
//
// 这些数位是反向存放的，也就是个位排在链表首部。
//
// 编写函数对这两个整数求和，并用链表形式返回结果。
//
//
//
// 示例：
//
// 输入：(7 -> 1 -> 6) + (5 -> 9 -> 2)，即617 + 295
//输出：2 -> 1 -> 9，即912
//
//
// 进阶：思考一下，假设这些数位是正向存放的，又该如何解决呢?
//
// 示例：
//
// 输入：(6 -> 1 -> 7) + (2 -> 9 -> 5)，即617 + 295
//输出：9 -> 1 -> 2，即912
//
// Related Topics 链表 数学
// 👍 73 👎 0

package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddTwoNumbers(t *testing.T) {
	assert.Equal(t, NewList([]int{2, 1, 9}).Array(), addTwoNumbers(NewList([]int{7, 1, 6}), NewList([]int{5, 9, 2})).Array())
	assert.Equal(t, NewList([]int{7, 4, 3}).Array(), addTwoNumbers(NewList([]int{2, 4, 3}), NewList([]int{5})).Array())
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	var (
		dummy  = new(ListNode)
		first  = l1
		second = l2
		step   = 0
	)
	dummy.Next = new(ListNode)
	next := dummy.Next
	for first != nil || second != nil {
		v1, v2 := 0, 0
		if first != nil {
			v1 = first.Val
			first = first.Next
		}
		if second != nil {
			v2 = second.Val
			second = second.Next
		}

		sum := v1 + v2 + step
		v := sum
		if sum > 9 {
			step = sum / 10
			v = sum % 10
		} else {
			step = 0
		}

		next.Val = v
		if first != nil || second != nil {
			next.Next = new(ListNode)
			next = next.Next
		} else {
			if step > 0 {
				n := new(ListNode)
				n.Val = step
				next.Next = n
			}
		}
	}

	return dummy.Next
}
