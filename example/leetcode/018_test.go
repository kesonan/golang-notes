//给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c +
// d 的值与 target 相等？找出所有满足条件且不重复的四元组。
//
// 注意：答案中不可以包含重复的四元组。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,0,-1,0,-2,2], target = 0
//输出：[[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
//
//
// 示例 2：
//
//
//输入：nums = [], target = 0
//输出：[]
//
//
//
//
// 提示：
//
//
// 0 <= nums.length <= 200
// -109 <= nums[i] <= 109
// -109 <= target <= 109
//
// Related Topics 数组 哈希表 双指针
// 👍 850 👎 0

package leetcode

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFourSum(t *testing.T) {
	// -2,-1,0,0,1,2
	assert.Equal(t, [][]int{
		{-2, -1, 1, 2},
		{-2, 0, 0, 2},
		{-1, 0, 0, 1},
	}, fourSum([]int{1, 0, -1, 0, -2, 2}, 0))

	assert.Equal(t, [][]int{
		{0, 0, 1, 2},
	}, fourSum([]int{1, 0, -1, 0, -2, 2}, 3))

	assert.Equal(t, [][]int{
		{-2, -1, 1, 2},
		{-1, -1, 1, 1},
	}, fourSum([]int{-2, -1, -1, 1, 1, 2, 2}, 0))

	// -5,-4,-3,-2,1,3,3,5
	assert.Equal(t, [][]int{
		{-5, -4, -3, 1},
	}, fourSum([]int{1, -2, -5, -4, -3, 3, 3, 5}, -11))

	assert.Equal(t, [][]int(nil), fourSum([]int{}, 0))
}

func fourSum(nums []int, target int) [][]int {
	var ret [][]int
	if len(nums) < 4 {
		return ret
	}

	n := len(nums)
	sort.Ints(nums)
	for i, e := range nums {
		if i > 0 && nums[i-1] == e {
			continue
		}

		j := i + 1
		for j < n {
			if j > i+1 && nums[j] == nums[j-1] {
				j += 1
				continue
			}

			var left, right = j + 1, n - 1
			var flag = -2
			for left < right {
				sum := e + nums[j] + nums[left] + nums[right]
				if sum < target {
					left += 1
					flag = -1
				} else if sum > target {
					right -= 1
					flag = 1
				} else {
					switch flag {
					case -1:
						if left > j+1 && nums[left] == nums[left-1] {
							left += 1
							continue
						}
					case 0:
						if left > j+1 && nums[left] == nums[left-1] && right < n-1 && nums[right] == nums[right+1] {
							right -= 1
							continue
						}
					case 1:
						if right < n-1 && nums[right] == nums[right+1] {
							left += 1
							right -= 1
							continue
						}
					}
					ret = append(ret, []int{e, nums[j], nums[left], nums[right]})
					left += 1
					right -= 1
					flag = 0
				}
			}

			j += 1
		}
	}

	return ret
}
