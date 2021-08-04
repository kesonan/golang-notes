package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPowerOfFour(t *testing.T) {
	assert.True(t, isPowerOfFour(1))
	assert.True(t, isPowerOfFour(4))
	assert.True(t, isPowerOfFour(16))
	assert.True(t, isPowerOfFour(64))
}

func isPowerOfFour(n int) bool {
	if n == 1 || n == 4 {
		return true
	}

	var v = 4
	for v < n {
		v = v * 4
	}

	return v == n
}
