package leetcode_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// https://leetcode.com/problems/valid-anagram/description/

func Test_242(t *testing.T) {
	assert.True(t, solve242("anagram", "nagaram"))
	assert.False(t, solve242("rat", "car"))
}

func solve242(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	seen := make(map[byte]int)

	for i := 0; i < len(a); i++ {
		seen[a[i]]++
		seen[b[i]]--
	}

	for _, count := range seen {
		if count != 0 {
			return false
		}
	}

	return true
}
