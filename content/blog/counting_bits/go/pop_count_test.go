package counting_bits

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_countBitsPopCount(t *testing.T) {
	testCases := []struct {
		name string
		num  int
		out  []int
	}{
		{
			name: "0",
			num:  0,
			out:  []int{0},
		},
		{
			name: "1",
			num:  1,
			out:  []int{0, 1},
		},
		{
			name: "2",
			num:  2,
			out:  []int{0, 1, 1},
		},
		{
			name: "5",
			num:  5,
			out:  []int{0, 1, 1, 2, 1, 2},
		},
		{
			name: "7",
			num:  7,
			out:  []int{0, 1, 1, 2, 1, 2, 2, 3},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res := countBitsPopCount(testCase.num)
			assert.Equal(t, testCase.out, res)
		})
	}

}
