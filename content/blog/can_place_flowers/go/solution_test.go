package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_canPlaceFlowers(t *testing.T) {
	testCases := []struct {
		name      string
		flowerbed []int
		n         int
		expected  bool
	}{
		{
			name:      "[1,0,0,0,1],n=1",
			flowerbed: []int{1, 0, 0, 0, 1},
			n:         1,
			expected:  true,
		},
		{
			name:      "[1,0,0,0,1],n=2",
			flowerbed: []int{1, 0, 0, 0, 1},
			n:         2,
			expected:  false,
		},
		{
			name:      "[1],n=0",
			flowerbed: []int{1},
			n:         0,
			expected:  true,
		},
		{
			name:      "[0],n=1",
			flowerbed: []int{0},
			n:         1,
			expected:  true,
		},
		{
			name:      "[1],n=1",
			flowerbed: []int{1},
			n:         1,
			expected:  false,
		},
		{
			name:      "[0,1,0],n=1",
			flowerbed: []int{0, 1, 0},
			n:         1,
			expected:  false,
		},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			res := canPlaceFlowers(c.flowerbed, c.n)

			assert.Equal(t, c.expected, res)
		})
	}
}
