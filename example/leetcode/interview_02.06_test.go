//编写一个函数，检查输入的链表是否是回文的。
//
//
//
// 示例 1：
//
// 输入： 1->2
//输出： false
//
//
// 示例 2：
//
// 输入： 1->2->2->1
//输出： true
//
//
//
//
// 进阶：
//你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？
// Related Topics 链表
// 👍 63 👎 0

package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	assert.True(t, isPalindrome(NewList([]int{1, 1})))
	assert.True(t, isPalindrome(NewList([]int{1, 2, 1})))
	assert.False(t, isPalindrome(NewList([]int{1, 2, 2})))
	assert.True(t, isPalindrome(NewList([]int{1, 2, 2, 1})))
	assert.True(t, isPalindrome(NewList([]int{0})))
	assert.True(t, isPalindrome(NewList([]int{0, 0, 0, 0})))
	assert.True(t, isPalindrome(NewList([]int{0, 0, 1, 1, 0, 0})))
}

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}

	var (
		list        []int
		cursor      = head
		left, right = 0, 0
	)

	for cursor != nil {
		list = append(list, cursor.Val)
		cursor = cursor.Next
	}
	right = len(list) - 1

	for left < right {
		if list[left] != list[right] {
			return false
		}

		left += 1
		right -= 1
	}

	return true
}
