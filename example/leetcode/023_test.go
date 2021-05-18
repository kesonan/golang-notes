//给你一个链表数组，每个链表都已经按升序排列。
//
// 请你将所有链表合并到一个升序链表中，返回合并后的链表。
//
//
//
// 示例 1：
//
// 输入：lists = [[1,4,5],[1,3,4],[2,6]]
//输出：[1,1,2,3,4,4,5,6]
//解释：链表数组如下：
//[
//  1->4->5,
//  1->3->4,
//  2->6
//]
//将它们合并到一个有序链表中得到。
//1->1->2->3->4->4->5->6
//
//
// 示例 2：
//
// 输入：lists = []
//输出：[]
//
//
// 示例 3：
//
// 输入：lists = [[]]
//输出：[]
//
//
//
//
// 提示：
//
//
// k == lists.length
// 0 <= k <= 10^4
// 0 <= lists[i].length <= 500
// -10^4 <= lists[i][j] <= 10^4
// lists[i] 按 升序 排列
// lists[i].length 的总和不超过 10^4
//
// Related Topics 堆 链表 分治算法
// 👍 1306 👎 0

package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// To pick the minimum head node from lists in order
// O(k*len(lists))
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	for i, e := range lists {
		dummy := new(ListNode)
		dummy.Next = e
		lists[i] = dummy
	}

	var (
		dummy = new(ListNode)
		cur   = dummy
	)
	for {
		node := getMinHead(lists)
		if node == nil {
			break
		}

		cur.Next = node
		cur = node
	}

	return dummy.Next
}

func getMinHead(lists []*ListNode) *ListNode {
	n := len(lists)
	if n == 0 {
		return nil
	}

	var (
		min   *ListNode
		index = -1
	)
	for i := 0; i < n; i++ {
		dummy := lists[i]
		node := dummy.Next
		if node == nil {
			continue
		}

		if min == nil {
			min = node
			index = i
			continue
		}

		if min.Val > node.Val {
			min = node
			index = i
		}
	}

	if index == -1 {
		return nil
	}

	lists[index].Next = min.Next
	return min
}
