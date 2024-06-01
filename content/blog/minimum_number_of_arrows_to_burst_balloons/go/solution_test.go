package minimum_number_of_arrows_to_burst_balloons

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findMinArrowShots(t *testing.T) {
	testCases := []struct {
		name   string
		points [][]int
		out    int
	}{
		{
			name: "[[10,16],[2,8],[1,6],[7,12]]",
			points: [][]int{
				{10, 16}, {2, 8}, {1, 6}, {7, 12},
			},
			out: 2,
		},
		{
			name: "[[1,2],[3,4],[5,6],[7,8]]",
			points: [][]int{
				{1, 2}, {3, 4}, {51, 6}, {7, 8},
			},
			out: 4,
		},
		{
			name: "[[1,2],[2,3],[3,4],[4,5]]",
			points: [][]int{
				{1, 2}, {2, 3}, {3, 4}, {4, 5},
			},
			out: 2,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res := findMinArrowShots(testCase.points)
			assert.Equal(t, res, testCase.out)
		})
	}
}
