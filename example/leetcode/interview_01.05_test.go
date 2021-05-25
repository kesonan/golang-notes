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
