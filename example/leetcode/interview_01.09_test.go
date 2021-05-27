//字符串轮转。给定两个字符串s1和s2，请编写代码检查s2是否为s1旋转而成（比如，waterbottle是erbottlewat旋转后的字符串）。
//
// 示例1:
//
//  输入：s1 = "waterbottle", s2 = "erbottlewat"
// 输出：True
//
//
// 示例2:
//
//  输入：s1 = "aa", s2 = "aba"
// 输出：False
//
//
//
//
//
// 提示：
//
//
// 字符串长度在[0, 100000]范围内。
//
//
// 说明:
//
//
// 你能只调用一次检查子串的方法吗？
//
// Related Topics 字符串
// 👍 72 👎 0

package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsFlipedString(t *testing.T) {
	assert.True(t, isFlipedString("waterbottle", "erbottlewat"))
	assert.False(t, isFlipedString("aa", "aba"))
	assert.False(t, isFlipedString("", ""))
}

func isFlipedString(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	if len(s1) == 0 {
		return true
	}

	var (
		n = len(s1)
	)
	for i := 0; i < len(s1); i++ {
		if s1[0:i] == s2[n-i:n] && s1[i:] == s2[:n-i] {
			return true
		}
	}

	return false
}
