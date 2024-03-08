package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_moveZeroes(t *testing.T) {
	testCases := []struct {
		name string
		in   []int
		out  []int
	}{
		{
			name: "[0,1,0,3,12]",
			in:   []int{0, 1, 0, 3, 12},
			out:  []int{1, 3, 12, 0, 0},
		},
		{
			name: "[1,2,3,4,5]",
			in:   []int{1, 2, 3, 4, 5},
			out:  []int{1, 2, 3, 4, 5},
		},
		{
			name: "[0,0,0,0,0]",
			in:   []int{0, 0, 0, 0, 0},
			out:  []int{0, 0, 0, 0, 0},
		},
		{
			name: "[0]",
			in:   []int{0},
			out:  []int{0},
		},
		{
			name: "[1]",
			in:   []int{1},
			out:  []int{1},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			moveZeroes(testCase.in)
			assert.Equal(t, testCase.in, testCase.out)
		})
	}
}
