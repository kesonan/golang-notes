package leetcode

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	assert.Equal(t, 3, lengthOfLongestSubstring("abcabcbb"))
	assert.Equal(t, 2, lengthOfLongestSubstring("au"))
	assert.Equal(t, 1, lengthOfLongestSubstring("c"))
	assert.Equal(t, 0, lengthOfLongestSubstring(""))
	assert.Equal(t, 1, lengthOfLongestSubstring("bbbbb"))
	assert.Equal(t, 3, lengthOfLongestSubstring("pwwkew"))
	assert.Equal(t, 2, lengthOfLongestSubstring("cdd"))
	assert.Equal(t, 3, lengthOfLongestSubstring("ohomm"))
}

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	if n < 2 {
		return len(s)
	}
	var (
		ret         int
		left, right int
		sub         string
	)
	for left < n && right < n {
		cursor := string(s[right])
		if !strings.Contains(sub, cursor) {
			right += 1
			sub = s[left:right]
			if right == n {
				if len(sub) > ret {
					ret = len(sub)
				}
			}
		} else {
			sub = s[left:right]
			if len(sub) > ret {
				ret = len(sub)
			}
			sub = ""
			left += 1
			right = left
		}
	}

	return ret
}
