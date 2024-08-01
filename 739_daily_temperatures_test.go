package leetcode_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// https://leetcode.com/problems/daily-temperatures/description/

func Test_739(t *testing.T) {
	assert.Equal(t, []int{1, 1, 0}, solve739([]int{30, 60, 90}))
	assert.Equal(t, []int{1, 1, 1, 0}, solve739([]int{30, 40, 50, 60}))
	assert.Equal(t, []int{1, 1, 4, 2, 1, 1, 0, 0}, solve739([]int{73, 74, 75, 71, 69, 72, 76, 73}))
}

func solve739(temperatures []int) []int {
	result := make([]int, len(temperatures))
	stack := make([]int, 0) // stores temperature indexes

	for index, temperature := range temperatures {
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] < temperature {
			topIndex := stack[len(stack)-1]
			result[topIndex] = index - topIndex
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, index)
	}

	return result
}
