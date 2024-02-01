package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_romanToInt(t *testing.T) {
	testCases := []struct {
		name     string
		in       string
		expected string
	}{
		{
			name:     "the sky is blue",
			in:       "the sky is blue",
			expected: "blue is sky the",
		},
		{
			name:     "  hello world  ",
			in:       "  hello world  ",
			expected: "world hello",
		},
		{
			name:     "a good   example",
			in:       "a good   example",
			expected: "example good a",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res := reverseWords(testCase.in)
			assert.Equal(t, testCase.expected, res)
		})
	}
}
