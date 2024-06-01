package two_close_strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_closeStrings(t *testing.T) {
	testCases := []struct {
		name     string
		str1     string
		str2     string
		expected bool
	}{
		{
			name:     "abc, bca",
			str1:     "abc",
			str2:     "bca",
			expected: true,
		},
		{
			name:     "a, aa",
			str1:     "a",
			str2:     "aa",
			expected: false,
		},
		{
			name:     "cabbba, abbccc",
			str1:     "cabbba",
			str2:     "abbccc",
			expected: true,
		},
		{
			name:     "cabbba, bbcccd",
			str1:     "cabbba",
			str2:     "bbcccd",
			expected: false,
		},
		{
			name:     "abbzzca, babzzcz",
			str1:     "abbzzca",
			str2:     "babzzcz",
			expected: false,
		},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			res := closeStrings(c.str1, c.str2)
			assert.Equal(t, c.expected, res)
		})
	}
}
