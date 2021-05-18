package leetcode

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsMatch(t *testing.T) {
	assert.False(t, isMatch("aa", "a"))
	assert.True(t, isMatch("aa", "a*"))
	assert.True(t, isMatch("ab", ".*"))
	assert.True(t, isMatch("aab", "c*a*b"))
}

func isMatch(s string, p string) bool {
	if p == ".*" {
		return true
	}
	list := strings.Split(s, "")
	index := match(list, 0, p)
	return index == 0
}

func match(list []string, startIndex int, p string) int {
	var preChar = list[0]
	pl := strings.Split(p, "")
	index := startIndex
	for _, r := range pl {
		switch r {
		case "*":
			index = matchStar(list[startIndex:], index, preChar)
			if index == -1 {
				return index
			}
		case ".":
			index = matchPoint(list[startIndex:], index, preChar)
			if index == -1 {
				return index
			}
		default:
			index = matchChar(list[startIndex:], index, preChar)
			if index == -1 {
				return index
			}

		}
		preChar = r
		index += 1
		if index == len(list) {
			return 0
		}
	}

	return -1
}

// 0-n
func matchStar(list []string, startIndex int, base string) int {
	var index = -1
	for i := startIndex; i < len(list); i++ {
		if i > startIndex && list[i] == list[i-1] && list[i] == base {
			index = i
			continue
		}

		if list[i] == base {
			index = i
			continue
		}

		return startIndex
	}

	return index
}

// 1
func matchPoint(list []string, startIndex int, _ string) int {
	return startIndex
}

// equal
func matchChar(list []string, startIndex int, base string) int {
	if list[startIndex] == base {
		return startIndex
	}

	return -1
}
