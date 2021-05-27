//编写一种算法，若M × N矩阵中某个元素为0，则将其所在的行与列清零。
//
//
//
// 示例 1：
//
// 输入：
//[
//  [1,1,1],
//  [1,0,1],
//  [1,1,1]
//]
//输出：
//[
//  [1,0,1],
//  [0,0,0],
//  [1,0,1]
//]
//
//
// 示例 2：
//
// 输入：
//[
//  [0,1,2,0],
//  [3,4,5,2],
//  [1,3,1,5]
//]
//输出：
//[
//  [0,0,0,0],
//  [0,4,5,0],
//  [0,3,1,0]
//]
//
// Related Topics 数组
// 👍 30 👎 0

package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetZeroes(t *testing.T) {
	matrix := [][]int{
		{0, 1, 2, 0},
		{3, 4, 5, 2},
		{1, 3, 1, 5},
	}
	setZeroes(matrix)
	assert.Equal(t, [][]int{
		{0, 0, 0, 0},
		{0, 4, 5, 0},
		{0, 3, 1, 0},
	}, matrix)

	matrix = [][]int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	}
	setZeroes(matrix)
	assert.Equal(t, [][]int{
		{1, 0, 1},
		{0, 0, 0},
		{1, 0, 1},
	}, matrix)
}

func setZeroes(matrix [][]int) {
	var (
		n    = len(matrix)
		m    = len(matrix[0])
		list = make([][]int, n)
	)

	copy(list, matrix)

	if m == 0 || n == 0 {
		return
	}

	var xs = make(map[int]struct{})
	for y := 0; y < n; y++ {
		ixs := make(map[int]struct{})
		for x := 0; x < m; x++ {
			if matrix[y][x] == 0 {
				ixs[x] = struct{}{}
			}
		}
		if len(ixs) == 0 {
			continue
		}
		for k := range ixs {
			xs[k] = struct{}{}
		}

		list[y] = make([]int, m)
	}

	for x := range xs {
		for y := 0; y < n; y++ {
			list[y][x] = 0
		}
	}

	copy(matrix, list)
}
