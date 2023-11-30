package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type maxProfitMinMaxSuit struct {
	name   string
	n      int
	expRes []string
}

func Test_fizzBuzz(t *testing.T) {
	var suits = []maxProfitMinMaxSuit{
		{
			name:   "3",
			n:      3,
			expRes: []string{"1", "2", "Fizz"},
		},
		{
			name:   "5",
			n:      5,
			expRes: []string{"1", "2", "Fizz", "4", "Buzz"},
		},
		{
			name:   "15",
			n:      15,
			expRes: []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"},
		},
	}

	for _, suit := range suits {
		t.Run(suit.name, func(t *testing.T) {
			res := fizzBuzz(suit.n)
			assert.Equal(t, suit.expRes, res)
		})
	}
}
