package maximum_average_subarray_i

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findMaxAverage(t *testing.T) {
	testCases := []struct {
		name string
		nums []int
		k    int
		out  float64
	}{
		{
			name: "[1,12,-5,-6,50,3]",
			nums: []int{1, 12, -5, -6, 50, 3},
			k:    4,
			out:  12.75,
		},
		{
			name: "[5]",
			nums: []int{5},
			k:    1,
			out:  5,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res := findMaxAverage(testCase.nums, testCase.k)
			assert.Equal(t, res, testCase.out)
		})
	}
}
