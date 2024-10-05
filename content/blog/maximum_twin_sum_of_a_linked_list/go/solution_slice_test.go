package maximum_twin_sum_of_a_linked_list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_pairSumSlice(t *testing.T) {
	testCases := []struct {
		name   string
		values []int
		out    int
	}{
		{
			name:   "[5,4,2,1]",
			values: []int{5, 4, 2, 1},
			out:    6,
		},
		{
			name:   "[4,2,2,3]",
			values: []int{4, 2, 2, 3},
			out:    7,
		},
		{
			name:   "[1,100000]",
			values: []int{1, 100000},
			out:    100001,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			head := valuesSliceToList(testCase.values)

			res := pairSumSlice(head)
			assert.Equal(t, res, testCase.out)
		})
	}
}
