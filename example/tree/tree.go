package tree

type BinaryTree interface {
	Traversal(tp TraversalType) ([]*Node, error)
	Depth() int
	IsEmpty() bool
	NodeCount() int
}

var _ BinaryTree = (*binaryTree)(nil)

type TraversalType int

const (
	PreOrder TraversalType = iota
	InOrder
	PostOrder
	Level
)
