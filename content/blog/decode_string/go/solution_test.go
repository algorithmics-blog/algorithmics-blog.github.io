package decode_string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_generateParenthesis(t *testing.T) {
	testCases := []struct {
		name string
		in   string
		out  string
	}{
		{
			name: "3[a]2[bc]",
			in:   "3[a]2[bc]",
			out:  "aaabcbc",
		},
		{
			name: "3[a2[c]]",
			in:   "3[a2[c]]",
			out:  "accaccacc",
		},
		{
			name: "2[abc]3[cd]ef",
			in:   "2[abc]3[cd]ef",
			out:  "abcabccdcdcdef",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res := decodeString(testCase.in)
			assert.Equal(t, testCase.out, res)
		})
	}
}
