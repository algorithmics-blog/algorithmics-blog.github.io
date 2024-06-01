package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isSameAfterReversals(t *testing.T) {
	testCases := []struct {
		name string
		num  int
		out  bool
	}{
		{
			name: "526",
			num:  526,
			out:  true,
		},
		{
			name: "1800",
			num:  1800,
			out:  false,
		},
		{
			name: "0",
			num:  0,
			out:  true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res := isSameAfterReversals(testCase.num)
			assert.Equal(t, testCase.out, res)
		})
	}

}
