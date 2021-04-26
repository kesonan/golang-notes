package tree

type AVLTree struct {
	*BinarySortTree
}

func NewAVLTree(compare func(parent, node interface{}) int) *AVLTree {
	return &AVLTree{
		NewBinarySortTree(compare),
	}
}

func (t *AVLTree) Balance() bool {
	return t.root.depthSub() >= -1 && t.root.depthSub() <= 1
}

func (t *AVLTree) Append(v interface{}) *Node {
	n := t.BinarySortTree.Append(v)
	vInt := t.root.depthSub()
	if vInt < -1 {
		t.lRotate(n)
	} else if vInt > 1 {
		t.rRotate(n)
	}

	return n
}

// lRotate:
// 1、原根节点右孩子成为新的根节点
// 2、原根节点右孩子的左子树变为原根节点的右子树
// 3、原根节点变为原根节点右孩子的左子树
func (t *AVLTree) lRotate(n *Node) {
	originalRoot := t.root
	newRoot := originalRoot.right
	originalRoot.right = newRoot.left
	newRoot.left.parent = originalRoot
	newRoot.left = originalRoot
	originalRoot.parent = newRoot
	newRoot.parent = nil
	t.root = newRoot
}

// rRotate:
// 1、原根节点的左孩子成为新的根节点
// 2、原根节点的左孩子的右子树变为原根节点的左子树
// 3、原根节点变为原根节点左孩子的右子树
func (t *AVLTree) rRotate(n *Node) {
	originalRoot := t.root
	newRoot := originalRoot.left
	originalRoot.left = newRoot.right
	newRoot.right.parent = originalRoot
	newRoot.right = originalRoot
	originalRoot.parent = newRoot
	newRoot.parent = nil
	t.root = newRoot
}
