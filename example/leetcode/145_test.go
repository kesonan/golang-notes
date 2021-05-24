//给定一个二叉树，返回它的 后序 遍历。
//
// 示例:
//
// 输入: [1,null,2,3]
//   1
//    \
//     2
//    /
//   3
//
//输出: [3,2,1]
//
// 进阶: 递归算法很简单，你可以通过迭代算法完成吗？
// Related Topics 栈 树
// 👍 588 👎 0

package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostorderTraversal(t *testing.T) {
	root := new(TreeNode)
	root.Val = 1
	r := new(TreeNode)
	r.Val = 2
	l := new(TreeNode)
	l.Val = 3
	r.Left = l
	root.Right = r
	assert.Equal(t, []int{3, 2, 1}, postorderTraversal(root))
}

func postorderTraversal(root *TreeNode) []int {
	return root.postorderTraversal()
}
