package maximum_depth_of_binary_tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestSuit struct {
	name  string
	input *TreeNode
	out   int
}

func Test_maxDepth(t *testing.T) {
	testCases := []TestSuit{
		{
			name: "[3,9,20,null,null,15,7]",
			input: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 9,
				},
				Right: &TreeNode{
					Val: 20,
					Left: &TreeNode{
						Val: 15,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
			out: 3,
		},
		{
			name: "[1,null,2]",
			input: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
				},
			},
			out: 2,
		},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			res := maxDepth(c.input)
			assert.Equal(t, c.out, res)
		})
	}
}
