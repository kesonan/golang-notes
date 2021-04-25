package tree

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinaryTree(t *testing.T) {
	root := &Node{
		value: "A",
	}
	b := root.AddLeft("B")
	d := b.AddLeft("D")
	e := b.AddRight("E")
	e.AddLeft("J")
	d.AddLeft("H")
	d.AddRight("I")

	c := root.AddRight("C")
	c.AddLeft("F")
	c.AddRight("G")

	binaryTree := NewBinaryTree(root)
	t.Run("traversalPreOrder", func(t *testing.T) {
		list, err := binaryTree.Traversal(PreOrder)
		assert.Nil(t, err)
		assert.Equal(t, "ABDHIEJCFG", convertNode2ArrayString(list, ""))
	})

	t.Run("traversalInOrder", func(t *testing.T) {
		list, err := binaryTree.Traversal(InOrder)
		assert.Nil(t, err)
		assert.Equal(t, "HDIBJEAFCG", convertNode2ArrayString(list, ""))
	})

	t.Run("traversalPostOrder", func(t *testing.T) {
		list, err := binaryTree.Traversal(PostOrder)
		assert.Nil(t, err)
		assert.Equal(t, "HIDJEBFGCA", convertNode2ArrayString(list, ""))
	})

	t.Run("traversalLevel", func(t *testing.T) {
		list, err := binaryTree.Traversal(Level)
		assert.Nil(t, err)
		assert.Equal(t, "ABCDEFGHIJ", convertNode2ArrayString(list, ""))
	})

	t.Run("Depth", func(t *testing.T) {
		assert.Equal(t, 4, binaryTree.Depth())
	})

	t.Run("Count", func(t *testing.T) {
		r := &Node{value: "A"}
		assert.Equal(t, 1, r.Count())
		r.AddLeft("B")
		assert.Equal(t, 2, r.Count())
		r.AddRight("C")
		assert.Equal(t, 3, r.Count())
		assert.Equal(t, 10, binaryTree.NodeCount())
	})
}

func TestBinarySortTree(t *testing.T) {
	list := []int{61, 87, 59, 47, 35, 73, 51, 98, 37, 93, 60}
	sortTree := NewBinarySortTree(func(parent, node interface{}) int {
		return parent.(int) - node.(int)
	})

	for _, e := range list {
		sortTree.Append(e)
	}

	nodes := sortTree.traversalInOrder()
	assert.Equal(t, "35,37,47,51,59,60,61,73,87,93,98", convertNode2ArrayString(nodes, ","))
	for _, e := range list {
		assert.True(t, sortTree.Search(e).value.(int) == e)
	}
	assert.True(t, sortTree.Search(85) == nil)

	sortTree.Delete(37)
	nodes = sortTree.traversalInOrder()
	assert.Equal(t, "35,47,51,59,60,61,73,87,93,98", convertNode2ArrayString(nodes, ","))

	sortTree.Delete(47)
	nodes = sortTree.traversalInOrder()
	assert.Equal(t, "35,51,59,60,61,73,87,93,98", convertNode2ArrayString(nodes, ","))

	sortTree.Delete(35)
	nodes = sortTree.traversalInOrder()
	assert.Equal(t, "51,59,60,61,73,87,93,98", convertNode2ArrayString(nodes, ","))

}

func convertNode2ArrayString(list []*Node, sep string) string {
	var join []string
	for _, item := range list {
		join = append(join, fmt.Sprintf("%v", item.value))
	}
	return strings.Join(join, sep)
}
