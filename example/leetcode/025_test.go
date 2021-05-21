//给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。
//
// k 是一个正整数，它的值小于或等于链表的长度。
//
// 如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
//
// 进阶：
//
//
// 你可以设计一个只使用常数额外空间的算法来解决此问题吗？
// 你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。
//
//
//
//
// 示例 1：
//
//
//输入：head = [1,2,3,4,5], k = 2
//输出：[2,1,4,3,5]
//
//
// 示例 2：
//
//
//输入：head = [1,2,3,4,5], k = 3
//输出：[3,2,1,4,5]
//
//
// 示例 3：
//
//
//输入：head = [1,2,3,4,5], k = 1
//输出：[1,2,3,4,5]
//
//
// 示例 4：
//
//
//输入：head = [1], k = 1
//输出：[1]
//
//
//
//
//
// 提示：
//
//
// 列表中节点的数量在范围 sz 内
// 1 <= sz <= 5000
// 0 <= Node.val <= 1000
// 1 <= k <= sz
//
// Related Topics 链表
// 👍 1106 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseKGroup(t *testing.T) {
	assert.Equal(t, []int{2, 1, 4, 3, 5}, reverseKGroup(NewList([]int{1, 2, 3, 4, 5}), 2).Array())
	assert.Equal(t, []int{3, 2, 1, 4, 5}, reverseKGroup(NewList([]int{1, 2, 3, 4, 5}), 3).Array())
	assert.Equal(t, []int{4, 3, 2, 1, 5}, reverseKGroup(NewList([]int{1, 2, 3, 4, 5}), 4).Array())
	assert.Equal(t, []int{5, 4, 3, 2, 1}, reverseKGroup(NewList([]int{1, 2, 3, 4, 5}), 5).Array())
	assert.Equal(t, []int{1, 2, 3, 4, 5}, reverseKGroup(NewList([]int{1, 2, 3, 4, 5}), 1).Array())
	assert.Equal(t, []int{2, 1, 4, 3, 6, 5}, reverseKGroup(NewList([]int{1, 2, 3, 4, 5, 6}), 2).Array())
}

// stack
func reverseKGroup(head *ListNode, k int) *ListNode {
	var (
		list       [][]*ListNode
		cursor     = head
		count      = -1
		segment    []*ListNode
		headCursor = head
	)

	for cursor != nil {
		count += 1
		segment = append(segment, cursor)
		if count == k-1 {
			list = append(list, segment)
			count = -1
			headCursor = cursor.Next
			segment = []*ListNode{}
		}

		cursor = cursor.Next
	}

	if len(list) == 0 {
		return nil
	}

	var (
		nh *ListNode
		nc *ListNode
	)
	for index, e := range list {
		for i := k - 1; i >= 0; i-- {
			if index == 0 && i == k-1 {
				nh = e[i]
			}

			if nc == nil {
				nc = e[i]
				continue
			}

			nc.Next = e[i]
			nc = e[i]
		}
	}

	if nc != nil {
		nc.Next = headCursor
	}

	return nh
}
