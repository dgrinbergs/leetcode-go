package leetcode_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// https://leetcode.com/problems/two-sum/description/

func Test_1(t *testing.T) {
	assert.Contains(t, solve1([]int{2, 7, 11, 15}, 9), 0, 1)
	assert.Contains(t, solve1([]int{3, 2, 4}, 6), 1, 2)
	assert.Contains(t, solve1([]int{3, 3}, 6), 0, 1)
}

func solve1(nums []int, target int) []int {
	if len(nums) == 2 {
		return []int{0, 1}
	}

	remaining := make(map[int]int)

	for i, num := range nums {
		if j, ok := remaining[num]; ok {
			return []int{j, i}
		}

		remaining[target-num] = i
	}

	return make([]int, 0)
}
