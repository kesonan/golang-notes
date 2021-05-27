//实现一种算法，找出单向链表中倒数第 k 个节点。返回该节点的值。
//
// 注意：本题相对原题稍作改动
//
// 示例：
//
// 输入： 1->2->3->4->5 和 k = 2
//输出： 4
//
// 说明：
//
// 给定的 k 保证是有效的。
// Related Topics 链表 双指针
// 👍 70 👎 0

package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKthToLast(t *testing.T) {
	assert.Equal(t, 4, kthToLast(NewList([]int{1, 2, 3, 4, 5}), 2))
	assert.Equal(t, 3, kthToLast(NewList([]int{1, 2, 3, 4, 5}), 3))
	assert.Equal(t, 1, kthToLast(NewList([]int{1, 2, 3, 4, 5}), 5))
	assert.Equal(t, 5, kthToLast(NewList([]int{1, 2, 3, 4, 5}), 1))
}

func kthToLast(head *ListNode, k int) int {
	if head == nil {
		return 0
	}

	var (
		cursor       = head
		firstCursor  = head
		secondCursor *ListNode
		count        = 0
	)

	for cursor != nil {
		count += 1
		if count == k {
			secondCursor = cursor
			break
		}
		cursor = cursor.Next

	}

	if secondCursor == nil {
		return 0
	}

	firstCursor = head
	cursor = secondCursor
	for cursor.Next != nil {
		cursor = cursor.Next
		firstCursor = firstCursor.Next
	}

	return firstCursor.Val
}
