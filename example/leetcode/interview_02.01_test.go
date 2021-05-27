//编写代码，移除未排序链表中的重复节点。保留最开始出现的节点。
//
// 示例1:
//
//
// 输入：[1, 2, 3, 3, 2, 1]
// 输出：[1, 2, 3]
//
//
// 示例2:
//
//
// 输入：[1, 1, 1, 1, 2]
// 输出：[1, 2]
//
//
// 提示：
//
//
// 链表长度在[0, 20000]范围内。
// 链表元素在[0, 20000]范围内。
//
//
// 进阶：
//
// 如果不得使用临时缓冲区，该怎么解决？
// Related Topics 链表
// 👍 105 👎 0

package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicateNodes(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3}, removeDuplicateNodes(NewList([]int{1, 2, 3, 3, 2, 1})).Array())
}

func removeDuplicateNodes(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	occurred := make(map[int]struct{})
	pos := head
	occurred[head.Val] = struct{}{}
	for pos.Next != nil {
		cur := pos.Next
		if _, ok := occurred[cur.Val]; !ok {
			occurred[cur.Val] = struct{}{}
			pos = pos.Next
		} else {
			pos.Next = pos.Next.Next
		}
	}

	pos.Next = nil
	return head
}
