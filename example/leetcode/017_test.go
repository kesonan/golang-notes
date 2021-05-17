//给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。
//
// 给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
//
//
//
//
//
// 示例 1：
//
//
//输入：digits = "23"
//输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]
//
//
// 示例 2：
//
//
//输入：digits = ""
//输出：[]
//
//
// 示例 3：
//
//
//输入：digits = "2"
//输出：["a","b","c"]
//
//
//
//
// 提示：
//
//
// 0 <= digits.length <= 4
// digits[i] 是范围 ['2', '9'] 的一个数字。
//
// Related Topics 深度优先搜索 递归 字符串 回溯算法
// 👍 1313 👎 0

package leetcode

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLetterCombinations(t *testing.T) {
	assert.Equal(t, []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}, letterCombinations("23"))
	assert.Equal(t, []string(nil), letterCombinations(""))
	assert.Equal(t, []string{"a", "b", "c"}, letterCombinations("2"))
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}

	if len(digits) == 1 {
		return numberMap(digits)
	}

	var i = len(digits) - 2
	c := getLetter(digits[i:i+1], numberMap(digits[i+1:i+2]))
	for {
		i -= 1
		if i < 0 {
			return c
		}
		c = getLetter(digits[i:i+1], c)
	}
}

func getLetter(parent string, c []string) []string {
	p := numberMap(parent)
	var list []string
	for _, e := range p {
		for _, ie := range c {
			list = append(list, e+ie)
		}
	}
	return list
}

func numberMap(digit string) []string {
	switch digit {
	case "2":
		return strings.Split("abc", "")
	case "3":
		return strings.Split("def", "")
	case "4":
		return strings.Split("ghi", "")
	case "5":
		return strings.Split("jkl", "")
	case "6":
		return strings.Split("mno", "")
	case "7":
		return strings.Split("pqrs", "")
	case "8":
		return strings.Split("tuv", "")
	case "9":
		return strings.Split("wxyz", "")
	default:
		return nil
	}
}
