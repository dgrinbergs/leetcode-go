package leetcode_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_238(t *testing.T) {
	assert.Equal(t, []int{24, 12, 8, 6}, solve238([]int{1, 2, 3, 4}))
	assert.Equal(t, []int{0, 0, 9, 0, 0}, solve238([]int{-1, 1, 0, -3, 3}))
}

func solve238(nums []int) []int {
	result := make([]int, len(nums))

	left := 1
	for i := 0; i < len(nums); i++ {
		result[i] = left
		left *= nums[i]
	}

	right := 1
	for i := len(nums) - 1; i >= 0; i-- {
		result[i] *= right
		right *= nums[i]
	}

	return result
}
