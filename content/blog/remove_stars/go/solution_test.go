package remove_stars

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_removeStars(t *testing.T) {
	testCases := []struct {
		name string
		in   string
		out  string
	}{
		{
			name: "leet**cod*e",
			in:   "leet**cod*e",
			out:  "lecoe",
		},
		{
			name: "erase*****",
			in:   "erase*****",
			out:  "",
		},
		{
			name: "nostars",
			in:   "nostars",
			out:  "nostars",
		},
		{
			name: "[empty]",
			in:   "",
			out:  "",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res := removeStars(testCase.in)
			assert.Equal(t, res, testCase.out)
		})
	}
}
