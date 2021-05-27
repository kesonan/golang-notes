//给定一个字符串，编写一个函数判定其是否为某个回文串的排列之一。
//
// 回文串是指正反两个方向都一样的单词或短语。排列是指字母的重新排列。
//
// 回文串不一定是字典当中的单词。
//
//
//
// 示例1：
//
// 输入："tactcoa"
//输出：true（排列有"tacocat"、"atcocta"，等等）
//
//
//
// Related Topics 哈希表 字符串
// 👍 51 👎 0

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
