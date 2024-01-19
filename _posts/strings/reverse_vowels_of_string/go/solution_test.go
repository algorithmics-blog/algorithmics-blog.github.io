package reverse_vowels_of_string

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Suit struct {
	name string
	in   string
	out  string
}

func Test_generateParenthesis(t *testing.T) {
	testCases := []Suit{
		{
			name: "algorithmics",
			in:   "algorithmics",
			out:  "ilgirothmacs",
		},
		{
			name: "AlgorithmIcs",
			in:   "AlgorithmIcs",
			out:  "IlgirothmAcs",
		},
		{
			name: "bae",
			in:   "bae",
			out:  "bea",
		},
		{
			name: "aeb",
			in:   "aeb",
			out:  "eab",
		},
		{
			name: "b",
			in:   "b",
			out:  "b",
		},
		{
			name: "ab",
			in:   "ab",
			out:  "ab",
		},
		{
			name: "ba",
			in:   "ba",
			out:  "ba",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res := reverseVowels(testCase.in)
			assert.Equal(t, testCase.out, res)
		})
	}
}
