package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_predictPartyVictory(t *testing.T) {
	testCases := []struct {
		name string
		in   string
		out  string
	}{
		{
			name: "RD",
			in:   "RD",
			out:  "Radiant",
		},
		{
			name: "RDD",
			in:   "RDD",
			out:  "Dire",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res := predictPartyVictory(testCase.in)

			assert.Equal(t, testCase.out, res)
		})
	}
}
