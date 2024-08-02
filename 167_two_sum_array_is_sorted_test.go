package leetcode_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/description/

func Test167(t *testing.T) {
	// for some bizarre reason indexes need to be 1-indexed...
	assert.ElementsMatch(t, []int{1, 2}, solve167([]int{2, 7, 11, 15}, 9))
	assert.ElementsMatch(t, []int{1, 3}, solve167([]int{2, 3, 4}, 6))
	assert.ElementsMatch(t, []int{1, 2}, solve167([]int{-1, 0}, -1))
	assert.ElementsMatch(t, []int{1, 2}, solve167([]int{-3, 3, 4, 90}, 0))
	assert.ElementsMatch(t, []int{2, 3}, solve167([]int{5, 25, 75}, 100))
}

func solve167(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1

	for left < right {
		sum := numbers[left] + numbers[right]
		if sum < target {
			left++
			continue
		} else if sum > target {
			right--
			continue
		}

		break
	}

	return []int{left + 1, right + 1}
}
