package kids_with_the_greatest_number_of_candies

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_kidsWithCandies(t *testing.T) {
	testCases := []struct {
		name         string
		candies      []int
		extraCandies int
		expected     []bool
	}{
		{
			name:         "[2,3,5,1,3], extraCandies=3",
			candies:      []int{2, 3, 5, 1, 3},
			extraCandies: 3,
			expected:     []bool{true, true, true, false, true},
		},
		{
			name:         "[4,2,1,1,2], extraCandies=1",
			candies:      []int{4, 2, 1, 1, 2},
			extraCandies: 1,
			expected:     []bool{true, false, false, false, false},
		},
		{
			name:         "[12,1,12], extraCandies=10",
			candies:      []int{12, 1, 12},
			extraCandies: 10,
			expected:     []bool{true, false, true},
		},
		{
			name:         "[2, 2], extraCandies=1",
			candies:      []int{2, 2},
			extraCandies: 3,
			expected:     []bool{true, true},
		},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			res := kidsWithCandies(c.candies, c.extraCandies)

			assert.Equal(t, c.expected, res)
		})
	}
}
