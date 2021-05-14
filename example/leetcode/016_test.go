//给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和
//。假定每组输入只存在唯一答案。
//
//
//
// 示例：
//
// 输入：nums = [-1,2,1,-4], target = 1
//输出：2
//解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。
//
//
//
//
// 提示：
//
//
// 3 <= nums.length <= 10^3
// -10^3 <= nums[i] <= 10^3
// -10^4 <= target <= 10^4
//
// Related Topics 数组 双指针
// 👍 775 👎 0
package leetcode

import (
	"math"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThreeSumClosest(t *testing.T) {
	assert.Equal(t, 2, threeSumClosest([]int{-1, 2, 1, 4}, 1))
	assert.Equal(t, 53, threeSumClosest([]int{-1, 2, 4, 4, 5, 6, -9, -4, -5, 3, 5, 14, 67, -89, 100, -200, 1, 1, 1, 1, 1, 1}, 50))
}

func threeSumClosest(nums []int, target int) int {
	n := len(nums)
	sort.Ints(nums)
	var sum = math.MaxInt32
	for i, e := range nums {
		left, right := i+1, n-1
		for left < right {
			tmp := e + nums[left] + nums[right]
			if tmp < target {
				left += 1
			} else if tmp > target {
				right -= 1
			} else {
				return tmp
			}
			if math.Abs(float64(tmp-target)) < math.Abs(float64(target-sum)) {
				sum = tmp
			}
		}
	}

	return sum
}
