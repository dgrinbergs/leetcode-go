package leetcode_go

import (
	"github.com/dgrinbergs/leetcode-git/datastructures"
	"github.com/stretchr/testify/assert"
	"testing"
)

// https://leetcode.com/problems/valid-parentheses/description/

func Test_20(t *testing.T) {
	assert.True(t, solve20("()"))
	assert.True(t, solve20("()[]{}"))
	assert.True(t, solve20("(([]))"))
	assert.False(t, solve20("(]"))
}

var parens = map[rune]rune{
	')': '(',
	'}': '{',
	']': '[',
}

func solve20(s string) bool {
	stack := datastructures.Stack[rune]{}

	for _, anyParen := range s {
		if openingParen, ok := parens[anyParen]; ok {
			last, ok := stack.Pop()
			if !ok {
				return false
			}

			if *last != openingParen {
				return false
			}

			continue
		}

		stack.Push(&anyParen)
	}

	return stack.Size() == 0
}
