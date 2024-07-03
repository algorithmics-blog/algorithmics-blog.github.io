package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_minFlips(t *testing.T) {
	testCases := []struct {
		name     string
		a        int
		b        int
		c        int
		expected int
	}{
		{
			name:     "a = 2, b = 6, c = 5",
			a:        2,
			b:        6,
			c:        5,
			expected: 3,
		},
		{
			name:     "a = 4, b = 2, c = 7",
			a:        4,
			b:        2,
			c:        7,
			expected: 1,
		},
		{
			name:     "a = 1, b = 2, c = 3",
			a:        1,
			b:        2,
			c:        3,
			expected: 0,
		},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			res := minFlips(c.a, c.b, c.c)
			assert.Equal(t, c.expected, res)
		})
	}
}
