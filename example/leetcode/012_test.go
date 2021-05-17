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
