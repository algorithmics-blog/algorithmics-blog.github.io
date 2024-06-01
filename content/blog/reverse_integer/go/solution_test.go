package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_reverse(t *testing.T) {
	testCases := []struct {
		name string
		num  int
		out  int
	}{
		{
			name: "123",
			num:  123,
			out:  321,
		},
		{
			name: "-123",
			num:  -123,
			out:  -321,
		},
		{
			name: "120",
			num:  120,
			out:  21,
		},
		{
			name: "2147483647",
			num:  2147483647,
			out:  0,
		},
		{
			name: "-2147483647",
			num:  -2147483647,
			out:  0,
		},
		{
			name: "-2147483412",
			num:  -2147483412,
			out:  -2143847412,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res := reverse(testCase.num)
			assert.Equal(t, testCase.out, res)
		})
	}

}
