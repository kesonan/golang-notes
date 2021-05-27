//实现一个算法，确定一个字符串 s 的所有字符是否全都不同。
//
// 示例 1：
//
// 输入: s = "leetcode"
//输出: false
//
//
// 示例 2：
//
// 输入: s = "abc"
//输出: true
//
//
// 限制：
//
// 0 <= len(s) <= 100
// 如果你不使用额外的数据结构，会很加分。
//
// Related Topics 数组
// 👍 115 👎 0

package leetcode

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsUnique(t *testing.T) {
	assert.False(t, isUnique("aa"))
	assert.True(t, isUnique("abc"))
	assert.False(t, isUnique("leetcode"))
}

func isUnique(astr string) bool {
	var sub string
	for i, e := range astr {
		if strings.Contains(sub, string(e)) {
			return false
		}
		sub = astr[0 : i+1]
	}

	return true
}
