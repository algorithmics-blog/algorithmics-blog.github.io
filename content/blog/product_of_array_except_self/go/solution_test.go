package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_productExceptSelf(t *testing.T) {
	testCases := []struct {
		name     string
		in       []int
		expected []int
	}{
		{
			name:     "[1,2,3,4]",
			in:       []int{1, 2, 3, 4},
			expected: []int{24, 12, 8, 6},
		},
		{
			name:     "[-1,1,0,-3,3]",
			in:       []int{-1, 1, 0, -3, 3},
			expected: []int{0, 0, 9, 0, 0},
		},
		{
			name:     "[1, 2]",
			in:       []int{1, 2},
			expected: []int{2, 1},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res := productExceptSelf(testCase.in)
			assert.Equal(t, testCase.expected, res)
		})
	}
}
