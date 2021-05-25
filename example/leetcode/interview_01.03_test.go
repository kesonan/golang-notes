package leetcode

import (
	"testing"
	"unicode"

	"github.com/stretchr/testify/assert"
)

func TestReplaceSpaces(t *testing.T) {
	assert.Equal(t, "Mr%20John%20Smith", replaceSpaces("Mr John Smith", 13))
	assert.Equal(t, "%20%20%20%20%20", replaceSpaces("     ", 5))
}

func replaceSpaces(s string, length int) string {
	bytes := []byte{}
	for i, e := range s {
		if i < length {
			if unicode.IsSpace(e) {
				bytes = append(bytes, '%', '2', '0')
			} else {
				bytes = append(bytes, byte(e))
			}
		}
	}

	return string(bytes)
}
