//给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重
//复的三元组。
//
// 注意：答案中不可以包含重复的三元组。
//
//
//
// 示例 1：
//
//
//输入：nums = [-1,0,1,2,-1,-4]
//输出：[[-1,-1,2],[-1,0,1]]
//
//
// 示例 2：
//
//
//输入：nums = []
//输出：[]
//
//
// 示例 3：
//
//
//输入：nums = [0]
//输出：[]
//
//
//
//
// 提示：
//
//
// 0 <= nums.length <= 3000
// -105 <= nums[i] <= 105
//
// Related Topics 数组 双指针
// 👍 3339 👎 0
package leetcode

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThreeSum(t *testing.T) {
	assert.Equal(t, [][]int{
		{-1, -1, 2},
		{-1, 0, 1},
	}, threeSum([]int{-1, 0, 1, 2, -1, -4}))

	assert.Equal(t, [][]int(nil), threeSum([]int{0}))
	assert.Equal(t, [][]int{{0, 0, 0}}, threeSum([]int{0, 0, 0, 0, 0, 1}))
}

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}

	var ret [][]int
	sort.Ints(nums)
	for i, e := range nums {
		if e > 0 {
			break
		}

		if i > 0 && e == nums[i-1] {
			continue
		}

		var left, right = i + 1, len(nums) - 1
		for left < right {
			if left != i+1 && nums[left-1] == nums[left] {
				left += 1
				continue
			}

			if nums[left]+nums[right] < -e {
				left += 1
			} else if nums[left]+nums[right] > -e {
				right -= 1
			} else {
				ret = append(ret, []int{e, nums[left], nums[right]})
				left += 1
				right -= 1
			}

		}
	}

	return ret
}
