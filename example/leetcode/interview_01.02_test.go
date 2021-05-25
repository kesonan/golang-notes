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
