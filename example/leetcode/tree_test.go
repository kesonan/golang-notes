package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) inorderTraversal() []int {
	if t == nil {
		return nil
	}
	var ret []int
	if t.Left != nil {
		ret = append(ret, t.Left.inorderTraversal()...)
	}
	ret = append(ret, t.Val)
	if t.Right != nil {
		ret = append(ret, t.Right.inorderTraversal()...)
	}
	return ret
}

func (t *TreeNode) postorderTraversal() []int {
	if t == nil {
		return nil
	}
	var ret []int
	if t.Left != nil {
		ret = append(ret, t.Left.inorderTraversal()...)
	}
	if t.Right != nil {
		ret = append(ret, t.Right.inorderTraversal()...)
	}
	ret = append(ret, t.Val)
	return ret
}
