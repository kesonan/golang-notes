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
