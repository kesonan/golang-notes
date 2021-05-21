package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewList(list []int) *ListNode {
	dummy := new(ListNode)
	head := new(ListNode)
	dummy.Next = head
	cursor := head
	for i, e := range list {
		cursor.Val = e
		if i < len(list)-1 {
			next := new(ListNode)
			cursor.Next = next
			cursor = next
		}
	}

	return dummy.Next
}

func (n *ListNode) Array() []int {
	var (
		list   []int
		cursor = n
	)
	for cursor != nil {
		list = append(list, cursor.Val)
		cursor = cursor.Next
	}

	return list
}
