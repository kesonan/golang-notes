package leetcode

func detectCycle(head *ListNode) *ListNode {
	var (
		m      = make(map[*ListNode]struct{})
		cursor = head
	)

	for cursor != nil {
		if _, ok := m[cursor]; ok {
			return cursor
		}

		m[cursor] = struct{}{}
		cursor = cursor.Next
	}

	return nil
}
