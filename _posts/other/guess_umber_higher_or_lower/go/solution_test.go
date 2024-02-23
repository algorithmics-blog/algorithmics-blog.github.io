package guess_umber_higher_or_lower

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Suit struct {
	name   string
	n      int
	secret int
}

func Test_guessNumber(t *testing.T) {
	testCases := []Suit{
		{
			name:   "n = 10, secret = 3",
			n:      10,
			secret: 3,
		},
		{
			name:   "n = 1, secret = 1",
			n:      1,
			secret: 1,
		},
		{
			name:   "n = 10, secret = 10",
			n:      10,
			secret: 10,
		},
		{
			name:   "n = 10, secret = 5",
			n:      10,
			secret: 5,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			guess = func(num int) int {
				if num > testCase.secret {
					return -1
				} else if num < testCase.secret {
					return 1
				}

				return 0
			}

			res := guessNumber(testCase.n)
			assert.Equal(t, testCase.secret, res)
		})
	}
}
