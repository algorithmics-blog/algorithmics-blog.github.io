package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_generateParenthesis(t *testing.T) {
	testCases := []struct {
		name string
		in   int
		out  []string
	}{
		{
			name: "example_1",
			in:   3,
			out:  []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
		{
			name: "example_2",
			in:   1,
			out:  []string{"()"},
		},
		{
			name: "example_8",
			in:   8,
			out:  []string{"()"},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res := generateParenthesis(testCase.in)
			sort.Strings(res)
			sort.Strings(testCase.out)

			assert.Equal(t, testCase.out, res)
		})
	}

}
