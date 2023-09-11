package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_romanToInt(t *testing.T) {
	testCases := []struct {
		name     string
		in       string
		expected int
	}{
		{
			name:     "III",
			in:       "III",
			expected: 3,
		},
		{
			name:     "LVIII",
			in:       "LVIII",
			expected: 58,
		},
		{
			name:     "MCMXCIV",
			in:       "MCMXCIV",
			expected: 1994,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res := romanToInt(testCase.in)
			assert.Equal(t, testCase.expected, res)
		})
	}
}
