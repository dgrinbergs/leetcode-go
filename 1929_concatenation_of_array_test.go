package leetcode_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// https://leetcode.com/problems/concatenation-of-array/description/

func Test_1929(t *testing.T) {
	assert.Equal(t, solve1929([]int{1, 2, 1}), []int{1, 2, 1, 1, 2, 1})
	assert.Equal(t, solve1929([]int{1, 3, 2, 1}), []int{1, 3, 2, 1, 1, 3, 2, 1})
}

func solve1929(nums []int) []int {
	size := len(nums)
	result := make([]int, size*2)

	for i, num := range nums {
		result[i] = num
		result[i+size] = num
	}

	return result
}
