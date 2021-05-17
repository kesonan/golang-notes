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
