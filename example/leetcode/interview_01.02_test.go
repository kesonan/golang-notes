//给定两个字符串 s1 和 s2，请编写一个程序，确定其中一个字符串的字符重新排列后，能否变成另一个字符串。
//
// 示例 1：
//
// 输入: s1 = "abc", s2 = "bca"
//输出: true
//
//
// 示例 2：
//
// 输入: s1 = "abc", s2 = "bad"
//输出: false
//
//
// 说明：
//
//
// 0 <= len(s1) <= 100
// 0 <= len(s2) <= 100
//
// Related Topics 数组 字符串
// 👍 34 👎 0
package leetcode

import (
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckPermutation(t *testing.T) {
	assert.True(t, CheckPermutation("asvnpzurz", "urzsapzvn"))
}

func CheckPermutation(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	return sortString(s1) == sortString(s2)
}

func sortString(s string) string {
	s2s := strings.Split(s, "")
	sort.Strings(s2s)
	s = strings.Join(s2s, "")
	return s
}
