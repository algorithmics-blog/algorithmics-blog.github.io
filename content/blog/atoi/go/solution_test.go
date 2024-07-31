package atoi

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_myAtoi(t *testing.T) {
	testCases := []struct {
		name string
		num  string
		out  int
	}{
		{
			name: "42",
			num:  "42",
			out:  42,
		},
		{
			name: "+42",
			num:  "+42",
			out:  42,
		},
		{
			name: "-42",
			num:  "-42",
			out:  -42,
		},
		{
			name: "    42",
			num:  "    42",
			out:  42,
		},
		{
			name: "    -42",
			num:  "    -42",
			out:  -42,
		},
		{
			name: "193 with words",
			num:  "193",
			out:  193,
		},
		{
			name: "+-12",
			num:  "+-12",
			out:  0,
		},
		{
			name: "+ 123",
			num:  "+ 123",
			out:  0,
		},
		{
			name: "00123",
			num:  "00123",
			out:  123,
		},
		{
			name: "empty string",
			num:  "",
			out:  0,
		},
		{
			name: "empty string with white spaces",
			num:  "     ",
			out:  0,
		},
		{
			name: "+",
			num:  "+",
			out:  0,
		},
		{
			name: "-",
			num:  "-",
			out:  0,
		},
		{
			name: "00000-42a1234",
			num:  "00000-42a1234",
			out:  0,
		},
		{
			name: "123-",
			num:  "123-",
			out:  123,
		},
		{
			name: "21474836++",
			num:  "21474836++",
			out:  21474836,
		},
		{
			name: "Overflow 32-bit 91283472332",
			num:  "91283472332",
			out:  2147483647,
		},
		{
			name: "Overflow 32-bit -91283472332",
			num:  "-91283472332",
			out:  -2147483648,
		},
		{
			name: "Overflow 32-bit 9223372036854775808",
			num:  "9223372036854775808",
			out:  2147483647,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res := myAtoi(testCase.num)
			assert.Equal(t, testCase.out, res)
		})
	}
}
