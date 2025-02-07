package number_of_provinces

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type suit struct {
	name        string
	isConnected [][]int
	expected    int
}

func Test_findCircleNum(t *testing.T) {
	testCases := []suit{
		{
			name: "example_1",
			isConnected: [][]int{
				{1, 1, 0},
				{1, 1, 0},
				{0, 0, 1},
			},
			expected: 2,
		},
		{
			name: "example_2",
			isConnected: [][]int{
				{1, 0, 0},
				{0, 1, 0},
				{0, 0, 1},
			},
			expected: 3,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res := findCircleNum(testCase.isConnected)
			assert.Equal(t, testCase.expected, res)
		})
	}
}
