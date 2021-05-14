//给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i,
//ai) 和 (i, 0) 。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
//
// 说明：你不能倾斜容器。
//
//
//
// 示例 1：
//
//
//
//
//输入：[1,8,6,2,5,4,8,3,7]
//输出：49
//解释：图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。
//
// 示例 2：
//
//
//输入：height = [1,1]
//输出：1
//
//
// 示例 3：
//
//
//输入：height = [4,3,2,1,4]
//输出：16
//
//
// 示例 4：
//
//
//输入：height = [1,2,1]
//输出：2
//
//
//
//
// 提示：
//
//
// n = height.length
// 2 <= n <= 3 * 104
// 0 <= height[i] <= 3 * 104
//
// Related Topics 数组 双指针
// 👍 2455 👎 0

package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxArea(t *testing.T) {
	assert.Equal(t, 49, maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	assert.Equal(t, 1, maxArea([]int{1, 1}))
	assert.Equal(t, 2, maxArea([]int{1, 2, 1}))
	assert.Equal(t, 0, maxArea([]int{}))
}

// 暴力破解
//func maxArea(height []int) int {
//	var ret int
//	for i := 0; i < len(height); i++ {
//		for j := i + 1; j < len(height); j++ {
//			w := j - i
//			h := int(math.Min(float64(height[i]), float64(height[j])))
//			tm := w * h
//			if tm > ret {
//				ret = tm
//			}
//		}
//	}
//	return ret
//}

// 双指针
func maxArea(height []int) int {
	var ret int
	var left, right = 0, len(height) - 1
	for right > left {
		w := right - left
		h := height[left]
		if height[left] < height[right] {
			left++
		} else {
			h = height[right]
			right--
		}

		tmp := w * h
		if tmp > ret {
			ret = tmp
		}
	}

	return ret
}
