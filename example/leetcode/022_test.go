package leetcode

import (
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateParenthesis(t *testing.T) {
	expected3 := []string{"((()))", "(()())", "(())()", "()(())", "()()()"}
	actual3 := generateParenthesis2(3)
	sort.Strings(expected3)
	sort.Strings(actual3)
	assert.Equal(t, expected3, actual3)

	expected4 := []string{"(((())))", "((()()))", "((())())", "((()))()", "(()(()))", "(()()())", "(()())()", "(())(())", "(())()()", "()((()))", "()(()())", "()(())()", "()()(())", "()()()()"}
	actual4 := generateParenthesis2(4)
	sort.Strings(expected4)
	sort.Strings(actual4)
	assert.Equal(t, expected4, actual4)

	assert.Equal(t, []string(nil), generateParenthesis2(0))
	assert.Equal(t, []string{"()"}, generateParenthesis2(1))
}

func generateParenthesis2(n int) []string {
	if n == 0 {
		return nil
	}

	return generateParen(0, 0, n, "")
}

func generateParen(left, right, n int, paren string) []string {
	var (
		ret []string
	)
	if left == n && right == n && paren != "" {
		ret = append(ret, paren)
	}

	if left < n {
		ret = append(ret, generateParen(left+1, right, n, paren+"(")...)
	}
	if right < left {
		ret = append(ret, generateParen(left, right+1, n, paren+")")...)
	}

	return ret
}

// answer1: tree
func generateParenthesis(n int) []string {
	if n == 0 {
		return nil
	}

	var ret []string
	terminalNodes := generateBinaryTree(n)
	for _, e := range terminalNodes {
		list, ok := e.traversal()
		if !ok {
			continue
		}
		ret = append(ret, strings.Join(list, ""))
	}
	return ret
}

type Node struct {
	value  string
	parent *Node
	left   *Node
	right  *Node
}

// traversal returns the result list from root node to  terminal node,
// remove the path which join with empty string are not balance
func (n *Node) traversal() ([]string, bool) {
	var (
		p       = n
		list    []string
		balance = 0
		paren   string
	)

	for {
		if p.value == "(" {
			balance += 1
			paren += "("
			if balance == 0 {
				return nil, false
			}
			p = p.parent
		} else if p.value == ")" {
			balance -= 1
			paren += ")"
			if balance == 0 {
				list = append(list, paren)
				paren = ""
			}
			p = p.parent
		} else {
			break
		}
	}

	return list, balance == 0
}

func (n *Node) isTerminal() bool {
	return n.left == nil && n.right == nil
}

func (n *Node) createChildren(depth int) []*Node {
	var ret []*Node
	if depth <= 0 {
		return ret
	}

	n.left = new(Node)
	n.right = new(Node)
	n.left.value = "("
	n.left.parent = n
	n.right.value = ")"
	n.right.parent = n
	if depth == 1 {
		ret = append(ret, n.left, n.right)
	}

	depth -= 1
	l1 := n.left.createChildren(depth)
	l2 := n.right.createChildren(depth)
	ret = append(ret, l1...)
	ret = append(ret, l2...)
	return ret
}

func generateBinaryTree(depth int) []*Node {
	dummy := new(Node)
	dummy.value = "*"
	return dummy.createChildren(depth * 2)
}
