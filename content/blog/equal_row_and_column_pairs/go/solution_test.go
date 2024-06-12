package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_equalPairs(t *testing.T) {
	testCases := []struct {
		name     string
		in       [][]int
		expected int
	}{
		{
			name: "[[3,2,1],[1,7,6],[2,7,7]]",
			in: [][]int{
				{3, 2, 1},
				{1, 7, 6},
				{2, 7, 7},
			},
			expected: 1,
		},
		{
			name: "[[3,1,2,2],[1,4,4,5],[2,4,2,2],[2,4,2,2]]",
			in: [][]int{
				{3, 1, 2, 2},
				{1, 4, 4, 5},
				{2, 4, 2, 2},
				{2, 4, 2, 2},
			},
			expected: 3,
		},
		{
			name: "[[5]]",
			in: [][]int{
				{5},
			},
			expected: 1,
		},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			res := equalPairs(c.in)
			assert.Equal(t, c.expected, res)
		})
	}
}
