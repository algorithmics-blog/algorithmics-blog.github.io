package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_twoSum(t *testing.T) {
	testCases := []struct {
		name     string
		chars    []byte
		expected int
	}{
		{
			name:     `["a","a","b","b","c","c","c"]`,
			chars:    []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'},
			expected: 6,
		},
		{
			name:     `["a"]`,
			chars:    []byte{'a'},
			expected: 1,
		},
		{
			name:     `["a","b","b","b","b","b","b","b","b","b","b","b","b"]`,
			chars:    []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'},
			expected: 4,
		},
		{
			name:     `100 a`,
			chars:    []byte{'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a'},
			expected: 4,
		},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			res := compress(c.chars)
			assert.Equal(t, c.expected, res)
		})
	}
}
