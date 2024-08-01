package leetcode_go

import (
	"github.com/stretchr/testify/assert"
	"slices"
	"testing"
)

// https://leetcode.com/problems/top-k-frequent-elements/description/

func Test_347(t *testing.T) {
	assert.Equal(t, []int{1, 2}, solve437([]int{1, 1, 1, 2, 2, 3}, 2))
	assert.Equal(t, []int{1}, solve437([]int{1}, 1))
}

func solve437(nums []int, k int) []int {
	frequencies := make(map[int]int)
	for _, num := range nums {
		frequencies[num]++
	}

	result := make([]int, 0)
	for num := range frequencies {
		result = append(result, num)
	}

	slices.SortFunc(result, func(a, b int) int {
		return frequencies[b] - frequencies[a]
	})

	return result[:k]
}
