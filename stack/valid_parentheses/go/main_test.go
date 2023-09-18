package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_generateParenthesis(t *testing.T) {
	testCases := []struct {
		name string
		in   string
		out  bool
	}{
		{
			name: "example_1",
			in:   "()",
			out:  true,
		},
		{
			name: "example_2",
			in:   "()[]{}",
			out:  true,
		},
		{
			name: "example_3",
			in:   "(]",
			out:  false,
		},
		{
			name: "example_4",
			in:   "((((",
			out:  false,
		},
		{
			name: "example_5",
			in:   "))))",
			out:  false,
		},
		{
			name: "example_6",
			in:   "(){",
			out:  false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res := isValid(testCase.in)

			assert.Equal(t, testCase.out, res)
		})
	}

}
