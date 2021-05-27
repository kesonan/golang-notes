//字符串压缩。利用字符重复出现的次数，编写一种方法，实现基本的字符串压缩功能。比如，字符串aabcccccaaa会变为a2b1c5a3。若“压缩”后的字符串没
//有变短，则返回原先的字符串。你可以假设字符串中只包含大小写英文字母（a至z）。
//
// 示例1:
//
//
// 输入："aabcccccaaa"
// 输出："a2b1c5a3"
//
//
// 示例2:
//
//
// 输入："abbccd"
// 输出："abbccd"
// 解释："abbccd"压缩后为"a1b2c2d1"，比原字符串长度更长。
//
//
// 提示：
//
//
// 字符串长度在[0, 50000]范围内。
//
// Related Topics 字符串
// 👍 84 👎 0

package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompressString(t *testing.T) {
	assert.Equal(t, "a2b1c5a3", compressString("aabcccccaaa"))
	assert.Equal(t, "abbccd", compressString("abbccd"))
	assert.Equal(t, "r4L31v11K14i28Z19I38o36b31v23l40B7K14f35D27s39N5Z26N10D15T5", compressString("rrrrLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLvvvvvvvvvvvKKKKKKKKKKKKKKiiiiiiiiiiiiiiiiiiiiiiiiiiiiZZZZZZZZZZZZZZZZZZZIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIoooooooooooooooooooooooooooooooooooobbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbvvvvvvvvvvvvvvvvvvvvvvvllllllllllllllllllllllllllllllllllllllllBBBBBBBKKKKKKKKKKKKKKfffffffffffffffffffffffffffffffffffDDDDDDDDDDDDDDDDDDDDDDDDDDDsssssssssssssssssssssssssssssssssssssssNNNNNZZZZZZZZZZZZZZZZZZZZZZZZZZNNNNNNNNNNDDDDDDDDDDDDDDDTTTTT"))
}

func compressString(S string) string {
	var (
		ret         string
		cursorCount = 0
	)
	for i, e := range S {
		if i > 0 && uint8(e) == S[i-1] {
			cursorCount += 1
			ret = fmt.Sprintf("%s%d", ret[:len(ret)-len(fmt.Sprintf("%d", cursorCount-1))], cursorCount)
			continue
		}
		x := string(e)
		cursorCount = 1
		ret = fmt.Sprintf("%s%s%d", ret, x, cursorCount)
	}

	if len(ret) < len(S) {
		return ret
	}

	return S
}
