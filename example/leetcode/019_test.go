//给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
//
// 进阶：你能尝试使用一趟扫描实现吗？
//
//
//
// 示例 1：
//
//
//输入：head = [1,2,3,4,5], n = 2
//输出：[1,2,3,5]
//
//
// 示例 2：
//
//
//输入：head = [1], n = 1
//输出：[]
//
//
// 示例 3：
//
//
//输入：head = [1,2], n = 1
//输出：[1]
//
//
//
//
// 提示：
//
//
// 链表中结点的数目为 sz
// 1 <= sz <= 30
// 0 <= Node.val <= 100
// 1 <= n <= sz
//
// Related Topics 链表 双指针
// 👍 1359 👎 0

package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

// 双指针
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	var left, right = dummy, head
	for i := 0; i < n; i++ {
		right = right.Next
	}

	for right != nil {
		right = right.Next
		left = left.Next
	}

	left.Next = left.Next.Next
	return dummy.Next
}

// 数组
//func removeNthFromEnd(head *ListNode, n int) *ListNode {
//	if head == nil {
//		return nil
//	}
//
//	m := make([]*ListNode, 0)
//	m = append(m, head)
//	next := head.Next
//	for next != nil {
//		m = append(m, next)
//		next = next.Next
//	}
//
//	index := len(m) - n
//	if index == 0 {
//		return m[index].Next
//	} else if index == len(m)-1 {
//		if len(m) > 1 {
//			m[index-1].Next = nil
//			return head
//		}
//
//		return nil
//	} else {
//		m[index-1].Next = m[index].Next
//		return head
//	}
//}
