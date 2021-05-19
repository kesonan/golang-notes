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
