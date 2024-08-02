package leetcode_go

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_238(t *testing.T) {
	assert.Equal(t, []int{24, 12, 8, 6}, solve238([]int{1, 2, 3, 4}))
	assert.Equal(t, []int{0, 0, 9, 0, 0}, solve238([]int{-1, 1, 0, -3, 3}))
}

func solve238(nums []int) []int {
	result := make([]int, len(nums))

	left := make([]int, len(nums)-1)
	for i := 0; i < len(nums)-1; i++ {
		if i == 0 {
			left[i] = nums[i]
			continue
		}

		left[i] = left[i-1] * nums[i]
	}
	fmt.Println("left", left)

	right := make([]int, len(nums))
	for i := len(nums) - 1; i > 0; i-- {
		if i == len(nums)-1 {
			right[i] = nums[i]
			continue
		}

		right[i] = right[i+1] * nums[i]
	}
	fmt.Println("right", right)

	for i := 0; i < len(nums); i++ {
		if i == 0 {
			result[i] = right[i+1]
			continue
		}
		if i == len(nums)-1 {
			result[i] = left[len(left)-1]
			continue
		}
		result[i] = left[i-1] * right[i+1]
	}
	fmt.Println("result", result)

	return result
}
