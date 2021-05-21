//给定一个字符串 s 和一些 长度相同 的单词 words 。找出 s 中恰好可以由 words 中所有单词串联形成的子串的起始位置。
//
// 注意子串要与 words 中的单词完全匹配，中间不能有其他字符 ，但不需要考虑 words 中单词串联的顺序。
//
//
//
// 示例 1：
//
//
//输入：s = "barfoothefoobarman", words = ["foo","bar"]
//输出：[0,9]
//解释：
//从索引 0 和 9 开始的子串分别是 "barfoo" 和 "foobar" 。
//输出的顺序不重要, [9,0] 也是有效答案。
//
//
// 示例 2：
//
//
//输入：s = "wordgoodgoodgoodbestword", words = ["word","good","best","word"]
//输出：[]
//
//
// 示例 3：
//
//
//输入：s = "barfoofoobarthefoobarman", words = ["bar","foo","the"]
//输出：[6,9,12]
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 104
// s 由小写英文字母组成
// 1 <= words.length <= 5000
// 1 <= words[i].length <= 30
// words[i] 由小写英文字母组成
//
// Related Topics 哈希表 双指针 字符串
// 👍 476 👎 0

package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindSubstring(t *testing.T) {
	assert.Equal(t, []int{0, 9}, findSubstring("barfoothefoobarman", []string{"foo", "bar"}))
	assert.Equal(t, []int(nil), findSubstring("wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"}))
	assert.Equal(t, []int{6, 9, 12}, findSubstring("barfoofoobarthefoobarman", []string{"bar", "foo", "the"}))
	assert.Equal(t, []int{8}, findSubstring("wordgoodgoodgoodbestword", []string{"word", "good", "best", "good"}))
}

func findSubstring(s string, words []string) []int {
	m := make(map[string]int)
	flag := make(map[int]struct{})
	for _, e := range words {
		if count, ok := m[e]; ok {
			m[e] = count + 1
		} else {
			m[e] = 1
		}
	}

	stepSize := len(words[0])
	var ret []int
	for i := 0; i < stepSize; i++ {
		l := slideWindow(i, stepSize, len(words)*stepSize, s, m)
		for _, x := range l {
			if _, ok := flag[x]; !ok {
				ret = append(ret, x)
				flag[x] = struct{}{}
			}
		}
	}

	return ret
}

func slideWindow(start, stepSize, L int, s string, m map[string]int) []int {
	var (
		ret []int
		i   = 0
	)
	for {
		st := i + start*stepSize
		ed := st + L
		if st > len(s) || ed > len(s) {
			return ret
		}

		sub := s[st:ed]
		if exists(sub, stepSize, m) {
			ret = append(ret, st)
		}

		i += 1
	}
}

func exists(sub string, stepSize int, m map[string]int) bool {
	tmp := make(map[string]int)
	var i = 0
	for {
		st := i * stepSize
		ed := st + stepSize
		if st > len(sub) || ed > len(sub) {
			if len(tmp) != len(m) {
				return false
			}
			for k, v := range m {
				if tmp[k] != v {
					return false
				}
			}
			return true
		}

		word := sub[st:ed]
		if _, ok := m[word]; ok {
			if count, ok := tmp[word]; ok {
				tmp[word] = count + 1
			} else {
				tmp[word] = 1
			}
		} else {
			return false
		}
		i += 1
	}
}
