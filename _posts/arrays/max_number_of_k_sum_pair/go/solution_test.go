package max_number_of_k_sum_pair

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxOperations(t *testing.T) {
	testCases := []struct {
		name string
		nums []int
		k    int
		out  int
	}{
		{
			name: "[1,2,3,4]",
			nums: []int{1, 2, 3, 4},
			k:    5,
			out:  2,
		},
		{
			name: "[3,1,3,4,3]",
			nums: []int{3, 1, 3, 4, 3},
			k:    6,
			out:  1,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res := maxOperations(testCase.nums, testCase.k)
			assert.Equal(t, res, testCase.out)
		})
	}
}
