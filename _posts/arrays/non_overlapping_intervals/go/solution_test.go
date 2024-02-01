package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_eraseOverlapIntervals(t *testing.T) {
	testCases := []struct {
		name     string
		in       [][]int
		expected int
	}{
		{
			name:     "[[1,2],[2,3],[3,4],[1,3]]",
			in:       [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}},
			expected: 1,
		},
		{
			name:     "[[1,2],[1,2],[1,2]]",
			in:       [][]int{{1, 2}, {1, 2}, {1, 2}},
			expected: 2,
		},
		{
			name:     "[[1,2],[2,3]]",
			in:       [][]int{{1, 2}, {2, 3}},
			expected: 0,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res := eraseOverlapIntervals(testCase.in)

			assert.Equal(t, testCase.expected, res)
		})
	}
}
