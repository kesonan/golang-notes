package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanPermutePalindrome(t *testing.T) {
	assert.True(t, canPermutePalindrome("tactcoa"))
	assert.True(t, canPermutePalindrome("abccba"))
	assert.False(t, canPermutePalindrome("abc"))
	assert.False(t, canPermutePalindrome("aabc"))
}

func canPermutePalindrome(s string) bool {
	var m = make(map[int32]int)
	for _, r := range s {
		if c, ok := m[r]; ok {
			c += 1
			m[r] = c
		} else {
			m[r] = 1
		}
	}

	var ret int
	for _, c := range m {
		if c%2 != 0 {
			ret += 1
		}
	}

	return ret < 2
}
