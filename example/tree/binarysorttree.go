package tree

type BinarySortTree struct {
	root    *Node
	compare func(parent, node interface{}) int
}

func NewBinarySortTree(compare func(parent, node interface{}) int) *BinarySortTree {
	return &BinarySortTree{
		compare: compare,
	}
}

func (b *BinarySortTree) Append(value interface{}) *Node {
	if b.root == nil {
		b.root = &Node{value: value}
		return b.root
	}

	return b.root.Append(b.root, value, b.compare)
}

func (b *BinarySortTree) Search(value interface{}) *Node {
	if b.root == nil {
		return nil
	}

	return b.root.Search(value, b.compare)
}

func (b *BinarySortTree) Delete(value interface{}) {
	if b.root == nil {
		return
	}

	node := b.Search(value)
	if node == nil {
		return
	}

	var nextNode *Node
	if node.left != nil && node.right != nil {
		list := b.traversalInOrder()
		var i int
		for index, e := range list {
			if e.value == value {
				i = index
				break
			}
		}

		if i == 0 {
			nextNode = list[1]
		} else if i == len(list)-1 {
			nextNode = list[len(list)-2]
		} else {
			nextNode = list[i-1]
		}

		if node.right != nextNode {
			nextNode.right = node.right
		}
		if node.left != nextNode {
			nextNode.left = node.left
		}
	} else {
		if node.left != nil {
			nextNode = node.left
		} else if node.right != nil {
			nextNode = node.right
		} else {
			nextNode = nil
		}
	}

	if node.parent != nil {
		if nextNode != nil {
			nextNode.parent = node.parent
		}

		if node == node.parent.left {
			node.parent.left = nextNode
		} else {
			node.parent.right = nextNode
		}
	}

}

func (b *BinarySortTree) Traversal(tp TraversalType) ([]*Node, error) {
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

func (b *BinarySortTree) Depth() int {
	return b.root.Depth()
}

func (b *BinarySortTree) IsEmpty() bool {
	return b.root == nil
}

func (b *BinarySortTree) NodeCount() int {
	if b.root == nil {
		return 0
	}

	return b.root.Count()
}

func (b *BinarySortTree) Clean() {
	b.root = nil
}

// traversalPreOrder visit from left to right recursively, and output the value of the node visited for the first time
func (b *BinarySortTree) traversalPreOrder() []*Node {
	if b.root == nil {
		return nil
	}

	return b.root.traversalPreOrder()
}

func (b *BinarySortTree) traversalInOrder() []*Node {
	if b.root == nil {
		return nil
	}

	return b.root.traversalInOrder()
}

func (b *BinarySortTree) traversalPostOrder() []*Node {
	if b.root == nil {
		return nil
	}

	return b.root.traversalPostOrder()
}

func (b *BinarySortTree) traversalLevel() []*Node {
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
