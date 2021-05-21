//实现获取 下一个排列 的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。
//
// 如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。
//
// 必须 原地 修改，只允许使用额外常数空间。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,2,3]
//输出：[1,3,2]
//
//
// 示例 2：
//
//
//输入：nums = [3,2,1]
//输出：[1,2,3]
//
//
// 示例 3：
//
//
//输入：nums = [1,1,5]
//输出：[1,5,1]
//
//
// 示例 4：
//
//
//输入：nums = [1]
//输出：[1]
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 100
// 0 <= nums[i] <= 100
//
// Related Topics 数组
// 👍 1125 👎 0

package leetcode

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNextPermutation(t *testing.T) {
	nums := []int{1, 2, 3}
	nextPermutation(nums)
	assert.Equal(t, []int{1, 3, 2}, nums)
	nextPermutation(nums)
	assert.Equal(t, []int{2, 1, 3}, nums)
	nextPermutation(nums)
	assert.Equal(t, []int{2, 3, 1}, nums)
	nextPermutation(nums)
	assert.Equal(t, []int{3, 1, 2}, nums)
	nextPermutation(nums)
	assert.Equal(t, []int{3, 2, 1}, nums)
	nextPermutation(nums)
	assert.Equal(t, []int{1, 2, 3}, nums)
}

func nextPermutation(nums []int) {
	length := len(nums)
	if length < 2 {
		return
	}

	var (
		leftIndex  = -1
		rightIndex = -1
	)
	for i := length - 1; i >= 0; i-- {
		if i > 0 && nums[i] > nums[i-1] {
			leftIndex = i - 1
			break
		}
	}

	if leftIndex == -1 {
		sort.Ints(nums)
		return
	}

	for i := length - 1; i >= 0; i-- {
		if nums[i] > nums[leftIndex] {
			rightIndex = i
			break
		}
	}

	if rightIndex == -1 {
		sort.Ints(nums)
		return
	}

	tmp := nums[leftIndex]
	nums[leftIndex] = nums[rightIndex]
	nums[rightIndex] = tmp
	sub := nums[leftIndex+1:]
	sort.Ints(sub)
}
