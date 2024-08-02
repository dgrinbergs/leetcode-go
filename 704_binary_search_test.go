package leetcode_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// https://leetcode.com/problems/binary-search/description/

func Test704(t *testing.T) {
	assert.Equal(t, 4, solve704([]int{-1, 0, 3, 5, 9, 12}, 9))
	assert.Equal(t, -1, solve704([]int{-1, 0, 3, 5, 9, 12}, 2))
}

func solve704(nums []int, target int) int {
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		}
		return -1
	}

	mid := len(nums) / 2

	if nums[mid] > target {
		return solve704(nums[:mid], target)
	}

	if index := solve704(nums[mid:], target); index == -1 {
		return -1
	} else {
		return index + mid
	}
}
