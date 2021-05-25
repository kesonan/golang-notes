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
