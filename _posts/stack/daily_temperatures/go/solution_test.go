package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_dailyTemperatures(t *testing.T) {
	testCases := []struct {
		name string
		in   []int
		out  []int
	}{
		{
			name: "[73,74,75,71,69,72,76,73]",
			in:   []int{73, 74, 75, 71, 69, 72, 76, 73},
			out:  []int{1, 1, 4, 2, 1, 1, 0, 0},
		},
		{
			name: "[30,40,50,60]",
			in:   []int{30, 40, 50, 60},
			out:  []int{1, 1, 1, 0},
		},
		{
			name: "[30,60,90]",
			in:   []int{30, 60, 90},
			out:  []int{1, 1, 0},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res := dailyTemperatures(testCase.in)

			assert.Equal(t, testCase.out, res)
		})
	}
}
