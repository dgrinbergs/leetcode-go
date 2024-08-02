package leetcode_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// https://leetcode.com/problems/container-with-most-water/submissions/

func Test_11(t *testing.T) {
	assert.Equal(t, 49, solve11([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	assert.Equal(t, 1, solve11([]int{1, 1}))
}

func solve11(heights []int) int {
	left := 0
	right := len(heights) - 1
	area := 0

	for left < right {
		width := right - left
		height := min(heights[left], heights[right])
		area = max(area, width*height)

		if heights[left] > heights[right] {
			right--
			continue
		}

		if heights[left] < heights[right] {
			left++
			continue
		}

		left++
		right--
	}

	return area
}
