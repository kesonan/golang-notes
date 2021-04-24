package tree

import "math"

type Node struct {
	value interface{}
	left  *Node
	right *Node
}

func (n *Node) AddLeft(value interface{}) *Node {
	if n.left != nil {
		n.left.value = value
	}

	n.left = &Node{value: value}
	return n.left
}

func (n *Node) AddRight(value interface{}) *Node {
	if n.right != nil {
		n.right.value = value
	}

	n.right = &Node{value: value}
	return n.right
}

func (n *Node) Depth() int {
	if n == nil {
		return 0
	}

	return int(math.Max(float64(n.left.Depth()), float64(n.right.Depth()))) + 1
}

func (n *Node) traversalPreOrder() []*Node {
	if n == nil {
		return nil
	}

	var list []*Node
	list = append(list, n)
	list = append(list, n.left.traversalPreOrder()...)
	list = append(list, n.right.traversalPreOrder()...)
	return list
}

func (n *Node) traversalInOrder() []*Node {
	if n == nil {
		return nil
	}

	var list []*Node
	list = append(list, n.left.traversalInOrder()...)
	list = append(list, n)
	list = append(list, n.right.traversalInOrder()...)
	return list
}

func (n *Node) traversalPostOrder() []*Node {
	if n == nil {
		return nil
	}

	var list []*Node
	list = append(list, n.left.traversalPostOrder()...)
	list = append(list, n.right.traversalPostOrder()...)
	list = append(list, n)
	return list
}

func (n *Node) traversalLevel() []*Node {
	if n == nil {
		return nil
	}

	var list []*Node
	if n.left != nil {
		list = append(list, n.left)
	}

	if n.right != nil {
		list = append(list, n.right)
	}
	return list
}

func (n *Node) Count() int {
	if n == nil {
		return 0
	}

	var count = 0
	if n.left != nil {
		count += n.left.Count()
	}

	if n.right != nil {
		count += n.right.Count()
	}

	return count + 1
}
