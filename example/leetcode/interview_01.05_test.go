//字符串有三种编辑操作:插入一个字符、删除一个字符或者替换一个字符。 给定两个字符串，编写一个函数判定它们是否只需要一次(或者零次)编辑。
//
//
//
// 示例 1:
//
// 输入:
//first = "pale"
//second = "ple"
//输出: True
//
//
//
// 示例 2:
//
// 输入:
//first = "pales"
//second = "pal"
//输出: False
//
// Related Topics 字符串 动态规划
// 👍 72 👎 0
package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOneEditAway(t *testing.T) {
	assert.True(t, oneEditAway("pale", "ple"))
	assert.False(t, oneEditAway("pales", "pal"))
	assert.True(t, oneEditAway("pales", "pale"))
	assert.True(t, oneEditAway("pales", "paled"))
}

func oneEditAway(first string, second string) bool {
	lf := len(first)
	ls := len(second)
	sub := lf - ls
	if sub > 1 || sub < -1 {
		return false
	}

	if lf < 2 && ls < 2 {
		return true
	}

	if lf == ls {
		count := 0
		for i := 0; i < lf; i++ {
			if first[i] != second[i] {
				count += 1
			}
		}
		return count < 2
	} else {
		var (
			min, max string
		)
		if lf > ls {
			max = first
			min = second
		} else {
			max = second
			min = first
		}

		for i := 0; i < len(min); i++ {
			if max[i] != min[i] {
				return max[0:i]+max[i+1:] == min
			}
		}
		return true
	}
}
