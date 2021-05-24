//给定一个二叉树的根节点 root ，返回它的 中序 遍历。
//
//
//
// 示例 1：
//
//
//输入：root = [1,null,2,3]
//输出：[1,3,2]
//
//
// 示例 2：
//
//
//输入：root = []
//输出：[]
//
//
// 示例 3：
//
//
//输入：root = [1]
//输出：[1]
//
//
// 示例 4：
//
//
//输入：root = [1,2]
//输出：[2,1]
//
//
// 示例 5：
//
//
//输入：root = [1,null,2]
//输出：[1,2]
//
//
//
//
// 提示：
//
//
// 树中节点数目在范围 [0, 100] 内
// -100 <= Node.val <= 100
//
//
//
//
// 进阶: 递归算法很简单，你可以通过迭代算法完成吗？
// Related Topics 栈 树 哈希表
// 👍 958 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInorderTraversal(t *testing.T) {
	root := new(TreeNode)
	root.Val = 1
	r := new(TreeNode)
	r.Val = 2
	l := new(TreeNode)
	l.Val = 3
	r.Left = l
	root.Right = r
	assert.Equal(t, []int{1, 3, 2}, inorderTraversal(root))
	assert.Equal(t, inorderTraversal2(root), inorderTraversal(root))
}

func inorderTraversal(root *TreeNode) []int {
	return root.inorderTraversal()
}

func inorderTraversal2(root *TreeNode) []int {
	stack := make([]*TreeNode, 0)
	var res []int
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack)-1]
		// 出栈
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)
		root = root.Right
	}
	return res
}
