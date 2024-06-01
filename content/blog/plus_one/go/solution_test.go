package plus_one

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_plusOnes(t *testing.T) {
	testCases := []struct {
		name   string
		digits []int
		out    []int
	}{
		{
			name:   "[1,2,3]",
			digits: []int{1, 2, 3},
			out:    []int{1, 2, 4},
		},
		{
			name:   "[2,2,9]",
			digits: []int{2, 2, 9},
			out:    []int{2, 3, 0},
		},
		{
			name:   "[9]",
			digits: []int{9},
			out:    []int{1, 0},
		},
		{
			name:   "[1]",
			digits: []int{1},
			out:    []int{2},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res := plusOne(testCase.digits)
			assert.Equal(t, testCase.out, res)
		})
	}

}
