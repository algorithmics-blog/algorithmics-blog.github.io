package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_asteroidCollisionTwoIdx(t *testing.T) {
	testCases := []struct {
		name string
		in   []int
		out  []int
	}{
		{
			name: "[5,10,-5]",
			in:   []int{5, 10, -5},
			out:  []int{5, 10},
		},
		{
			name: "[8,-8]",
			in:   []int{8, -8},
			out:  []int{},
		},
		{
			name: "[10,2,-5]",
			in:   []int{10, 2, -5},
			out:  []int{10},
		},
		{
			name: "[-8,8]",
			in:   []int{-8, 8},
			out:  []int{-8, 8},
		},
		{
			name: "[-8,-9]",
			in:   []int{-8, -9},
			out:  []int{-8, -9},
		},
		{
			name: "[8,9]",
			in:   []int{8, 9},
			out:  []int{8, 9},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res := asteroidCollisionTwoIdx(testCase.in)

			assert.Equal(t, testCase.out, res)
		})
	}
}
