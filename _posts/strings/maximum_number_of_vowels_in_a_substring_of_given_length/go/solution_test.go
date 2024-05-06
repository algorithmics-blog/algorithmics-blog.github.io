package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_maxVowels(t *testing.T) {
	testCases := []struct {
		name     string
		s        string
		k        int
		expected int
	}{
		{
			name:     "abciiidef",
			s:        "abciiidef",
			k:        3,
			expected: 3,
		},
		{
			name:     "aeiou",
			s:        "aeiou",
			k:        2,
			expected: 2,
		},
		{
			name:     "leetcode",
			s:        "leetcode",
			k:        3,
			expected: 2,
		},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			res := maxVowels(c.s, c.k)
			assert.Equal(t, c.expected, res)
		})
	}
}
