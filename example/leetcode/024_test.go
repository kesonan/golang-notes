//给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
//
// 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
//
//
//
// 示例 1：
//
//
//输入：head = [1,2,3,4]
//输出：[2,1,4,3]
//
//
// 示例 2：
//
//
//输入：head = []
//输出：[]
//
//
// 示例 3：
//
//
//输入：head = [1]
//输出：[1]
//
//
//
//
// 提示：
//
//
// 链表中节点的数目在范围 [0, 100] 内
// 0 <= Node.val <= 100
//
//
//
//
// 进阶：你能在不修改链表节点值的情况下解决这个问题吗?（也就是说，仅修改节点本身。）
// Related Topics 递归 链表
// 👍 917 👎 0

package leetcode

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var (
		remain             = head
		first, second, pre *ListNode
		dummy              = new(ListNode)
	)
	for {
		if remain == nil {
			break
		}

		remain, first, second = pickNodes(remain)
		if second == nil {
			if pre == nil {
				pre = first
				dummy.Next = pre
			} else {
				pre.Next = first
			}
			break
		}

		nf, ns := swapPair(first, second)
		if pre == nil {
			pre = ns
			dummy.Next = nf
		} else {
			pre.Next = nf
			pre = ns
		}
	}

	return dummy.Next
}

func pickNodes(head *ListNode) (*ListNode, *ListNode, *ListNode) {
	var remain *ListNode
	first := head
	second := head.Next
	if second != nil {
		remain = second.Next
	}

	return remain, first, second
}

func swapPair(first, second *ListNode) (*ListNode, *ListNode) {
	first.Next = nil
	second.Next = first
	return second, first
}
