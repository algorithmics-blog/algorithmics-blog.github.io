package is_subsequence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isSubsequences(t *testing.T) {
	testCases := []struct {
		name string
		s    string
		t    string
		out  bool
	}{
		{
			name: "s = \"abc\", t = \"ahbgdc\"",
			s:    "abc",
			t:    "ahbgdc",
			out:  true,
		},
		{
			name: "s = \"axc\", t = \"ahbgdc\"",
			s:    "axc",
			t:    "ahbgdc",
			out:  false,
		},
		{
			name: "s = \"\", t = \"ahbgdc\"",
			s:    "",
			t:    "ahbgdc",
			out:  true,
		},
		{
			name: "s = \"n\", t = \"c\"",
			s:    "b",
			t:    "c",
			out:  false,
		},
		{
			name: "s = \"b\", t = \"abc\"",
			s:    "b",
			t:    "abc",
			out:  true,
		},
		{
			name: "s = \"abc\", t = \"b\"",
			s:    "abc",
			t:    "b",
			out:  false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res := isSubsequence(testCase.s, testCase.t)
			assert.Equal(t, testCase.out, res)
		})
	}
}
