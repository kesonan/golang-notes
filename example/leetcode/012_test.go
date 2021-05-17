//罗马数字包含以下七种字符： I， V， X， L，C，D 和 M。
//
//
//字符          数值
//I             1
//V             5
//X             10
//L             50
//C             100
//D             500
//M             1000
//
// 例如， 罗马数字 2 写做 II ，即为两个并列的 1。12 写做 XII ，即为 X + II 。 27 写做 XXVII, 即为 XX + V + I
//I 。
//
// 通常情况下，罗马数字中小的数字在大的数字的右边。但也存在特例，例如 4 不写做 IIII，而是 IV。数字 1 在数字 5 的左边，所表示的数等于大数 5
// 减小数 1 得到的数值 4 。同样地，数字 9 表示为 IX。这个特殊的规则只适用于以下六种情况：
//
//
// I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
// X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。
// C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。
//
//
// 给你一个整数，将其转为罗马数字。
//
//
//
// 示例 1:
//
//
//输入: num = 3
//输出: "III"
//
// 示例 2:
//
//
//输入: num = 4
//输出: "IV"
//
// 示例 3:
//
//
//输入: num = 9
//输出: "IX"
//
// 示例 4:
//
//
//输入: num = 58
//输出: "LVIII"
//解释: L = 50, V = 5, III = 3.
//
//
// 示例 5:
//
//
//输入: num = 1994
//输出: "MCMXCIV"
//解释: M = 1000, CM = 900, XC = 90, IV = 4.
//
//
//
// 提示：
//
//
// 1 <= num <= 3999
//
// Related Topics 数学 字符串
// 👍 617 👎 0

package leetcode

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntToRoman(t *testing.T) {
	assert.Equal(t, "III", intToRoman(3))
	assert.Equal(t, "IV", intToRoman(4))
	assert.Equal(t, "IX", intToRoman(9))
	assert.Equal(t, "LVIII", intToRoman(58))
	assert.Equal(t, "MCMXCIV", intToRoman(1994))
	assert.Equal(t, "CI", intToRoman(101))
}

func intToRoman(num int) string {
	if num < 10 {
		return romanIn10(num)
	}

	if num < 100 {
		return romanIn10_100(num)
	}

	if num <= 1000 {
		return romanIn1000(num)
	}

	m := num / 1000
	ret := bytes.NewBufferString("")
	for i := 0; i < m; i++ {
		ret.WriteString("M")
	}

	ret.WriteString(romanIn1000(num - m*1000))
	return ret.String()
}

func romanIn10(num int) string {
	if num > 9 {
		return ""
	}

	switch num {
	case 1:
		return "I"
	case 2:
		return "II"
	case 3:
		return "III"
	case 4:
		return "IV"
	case 5:
		return "V"
	case 6:
		return "VI"
	case 7:
		return "VII"
	case 8:
		return "VIII"
	case 9:
		return "IX"
	default:
		return ""
	}
}

func romanIn10_100(num int) string {
	if num < 10 {
		return romanIn10(num)
	}

	v := romanIn10(num % 10)
	switch num / 10 {
	case 1:
		return "X" + v
	case 2:
		return "XX" + v
	case 3:
		return "XXX" + v
	case 4:
		return "XL" + v
	case 5:
		return "L" + v
	case 6:
		return "LX" + v
	case 7:
		return "LXX" + v
	case 8:
		return "LXXX" + v
	case 9:
		return "XC" + v
	default:
		return "" + v
	}
}

func romanIn1000(num int) string {
	if num > 1000 {
		return ""
	}

	v := romanIn10_100(num % 100)
	switch num / 100 {
	case 1:
		return "C" + v
	case 2:
		return "CC" + v
	case 3:
		return "CCC" + v
	case 4:
		return "CD" + v
	case 5:
		return "D" + v
	case 6:
		return "DC" + v
	case 7:
		return "DCC" + v
	case 8:
		return "DCCC" + v
	case 9:
		return "CM" + v
	case 10:
		return "M" + v
	default:
		return "" + v
	}
}
