package leetcode

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	dummyA := new(ListNode)
	dummyB := new(ListNode)
	dummyA.Next = headA
	dummyB.Next = headB
	m := make(map[*ListNode]struct{})
	if headA == nil || headB == nil {
		return nil
	}
	var (
		cursor = dummyA.Next
	)
	for cursor != nil {
		m[cursor] = struct{}{}
		cursor = cursor.Next
	}

	cursor = dummyB.Next
	for cursor != nil {
		if _, ok := m[cursor]; ok {
			return cursor
		}
		cursor = cursor.Next
	}

	return nil
}
