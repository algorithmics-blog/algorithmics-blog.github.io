package single_number

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type singleNumberSuit struct {
	name   string
	nums   []int
	expRes int
}

func Test_singleNumber(t *testing.T) {
	suits := []singleNumberSuit{
		{
			name:   "[2,2,1]",
			nums:   []int{2, 2, 1},
			expRes: 1,
		},
		{
			name:   "[4,1,2,1,2]",
			nums:   []int{4, 1, 2, 1, 2},
			expRes: 4,
		},
		{
			name:   "[1]",
			nums:   []int{1},
			expRes: 1,
		},
	}

	for _, suit := range suits {
		t.Run(suit.name, func(t *testing.T) {
			res := singleNumber(suit.nums)
			assert.Equal(t, suit.expRes, res)
		})
	}
}
