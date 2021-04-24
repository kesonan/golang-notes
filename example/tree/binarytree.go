package tree

import (
	"errors"
)

var errTraversalType = errors.New("invalid traversal type")

type binaryTree struct {
	root *Node
}

func NewBinaryTree(v interface{}) BinaryTree {
	return newBinaryTree(&Node{value: v})
}

func (b *binaryTree) Traversal(tp TraversalType) ([]*Node, error) {
	switch tp {
	case PreOrder:
		return b.traversalPreOrder(), nil
	case InOrder:
		return b.traversalInOrder(), nil
	case PostOrder:
		return b.traversalPostOrder(), nil
	case Level:
		return b.traversalLevel(), nil
	default:
		return nil, errTraversalType
	}
}

func (b *binaryTree) Depth() int {
	return b.root.Depth()
}

func (b *binaryTree) IsEmpty() bool {
	return b.root == nil
}

func (b *binaryTree) NodeCount() int {
	if b.root == nil {
		return 0
	}

	return b.root.Count()
}

func (b *binaryTree) Clean() {
	b.root = nil
}

func newBinaryTree(root *Node) BinaryTree {
	return &binaryTree{root: root}
}

// traversalPreOrder visit from left to right recursively, and output the value of the node visited for the first time
func (b *binaryTree) traversalPreOrder() []*Node {
	if b.root == nil {
		return nil
	}

	return b.root.traversalPreOrder()
}

func (b *binaryTree) traversalInOrder() []*Node {
	if b.root == nil {
		return nil
	}

	return b.root.traversalInOrder()
}

func (b *binaryTree) traversalPostOrder() []*Node {
	if b.root == nil {
		return nil
	}

	return b.root.traversalPostOrder()
}

func (b *binaryTree) traversalLevel() []*Node {
	if b.root == nil {
		return nil
	}

	var list []*Node
	list = append(list, b.root)
	var ret = list
	for len(ret) != 0 {
		ret = forRange(ret)
		list = append(list, ret...)
	}

	return list
}

func forRange(list []*Node) []*Node {
	var ret []*Node
	for _, e := range list {
		ret = append(ret, e.traversalLevel()...)
	}
	return ret
}
