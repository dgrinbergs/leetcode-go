package leetcode_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// https://leetcode.com/problems/contains-duplicate/description/

func Test_219(t *testing.T) {
	assert.Equal(t, solve217([]int{1, 2, 3, 1}), true)
	assert.Equal(t, solve217([]int{1, 2, 3, 4}), false)
	assert.Equal(t, solve217([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}), true)
}

func solve217(nums []int) bool {
	seen := make(map[int]bool)

	for _, num := range nums {
		if seen[num] {
			return true
		}

		seen[num] = true
	}

	return false
}
