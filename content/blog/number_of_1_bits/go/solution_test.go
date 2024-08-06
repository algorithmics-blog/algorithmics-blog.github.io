package number_of_1_bits

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_hammingWeight(t *testing.T) {
	testCases := []struct {
		name string
		num  int
		out  int
	}{
		{
			name: "0",
			num:  0,
			out:  0,
		},
		{
			name: "1",
			num:  1,
			out:  1,
		},
		{
			name: "11",
			num:  11,
			out:  3,
		},
		{
			name: "128",
			num:  128,
			out:  1,
		},
		{
			name: "2147483645",
			num:  2147483645,
			out:  30,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res := hammingWeight(testCase.num)
			assert.Equal(t, testCase.out, res)
		})
	}
}
