//将一个给定字符串 s 根据给定的行数 numRows ，以从上往下、从左到右进行 Z 字形排列。
//
// 比如输入字符串为 "PAYPALISHIRING" 行数为 3 时，排列如下：
//
//
//P   A   H   N
//A P L S I I G
//Y   I   R
//
// 之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："PAHNAPLSIIGYIR"。
//
// 请你实现这个将字符串进行指定行数变换的函数：
//
//
//string convert(string s, int numRows);
//
//
//
// 示例 1：
//
//
//输入：s = "PAYPALISHIRING", numRows = 3
//输出："PAHNAPLSIIGYIR"
//
//示例 2：
//
//
//输入：s = "PAYPALISHIRING", numRows = 4
//输出："PINALSIGYAHRPI"
//解释：
//P     I    N
//A   L S  I G
//Y A   H R
//P     I
//
//
// 示例 3：
//
//
//输入：s = "A", numRows = 1
//输出："A"
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 1000
// s 由英文字母（小写和大写）、',' 和 '.' 组成
// 1 <= numRows <= 1000
//
// Related Topics 字符串
// 👍 1137 👎 0

package leetcode

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvert(t *testing.T) {
	assert.Equal(t, "PAHNAPLSIIGYIR", convert("PAYPALISHIRING", 3))
	assert.Equal(t, "PINALSIGYAHRPI", convert("PAYPALISHIRING", 4))
	assert.Equal(t, "A", convert("A", 1))
}

func convert(s string, numRows int) string {
	turnCount := 0
	turnFactor := numRows - 1

	if turnFactor == 0 {
		return s
	}

	var ret = make([][]string, numRows)
	ret[0] = append(ret[0], s[0:1])
	var cur = 0
	for cur < len(s) {
		next := cur + turnFactor
		if turnCount%2 == 0 {
			ok := down(cur, next, s, ret)
			if !ok {
				break
			}
		} else {
			ok := up(cur, next, s, ret)
			if !ok {
				break
			}
		}

		cur = next
		turnCount += 1
	}
	writer := bytes.NewBufferString("")
	for _, e := range ret {
		writer.WriteString(strings.Join(e, ""))
	}

	return writer.String()
}

// dot not contains start
func down(start, end int, s string, ret [][]string) bool {
	for i := 0; i < end-start; i++ {
		if start+i+2 > len(s) {
			break
		}
		c := s[start+i+1 : start+i+2]
		if len(c) == 0 {
			continue
		}

		ret[i+1] = append(ret[i+1], c)
	}

	return end < len(s)
}

// dot not contains start
func up(start, end int, s string, ret [][]string) bool {
	divider := end - start
	for i := 0; i < divider; i++ {
		if start+i+2 > len(s) {
			break
		}

		c := s[start+i+1 : start+i+2]
		if len(c) == 0 {
			continue
		}

		ret[divider-i-1] = append(ret[divider-i-1], c)
	}

	return end < len(s)
}
