package maximum_frequency_stack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_reverseList(t *testing.T) {
	testCases := []struct {
		name  string
		input []int
		out   []int
	}{
		{
			name:  "[5,7,5,7,4,5]",
			input: []int{5, 7, 5, 7, 4, 5},
			out:   []int{5, 7, 5, 4},
		},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			stack := Constructor()
			for _, element := range c.input {
				stack.Push(element)
			}

			res := make([]int, 0, len(c.out))
			for i := 0; i < len(c.out); i++ {
				res = append(res, stack.Pop())
			}

			assert.Equal(t, c.out, res)
		})
	}
}
