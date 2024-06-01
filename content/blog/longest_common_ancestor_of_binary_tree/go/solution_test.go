package longest_common_ancestor_of_binary_tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Suit struct {
	name     string
	in       *TreeNode
	expected *TreeNode
	p        *TreeNode
	q        *TreeNode
}

func TestTrie_lowestCommonAncestor(t *testing.T) {
	tree := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 7,
				},
				Right: &TreeNode{
					Val: 4,
				},
			},
		},
		Right: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 0,
			},
			Right: &TreeNode{
				Val: 8,
			},
		},
	}

	testCases := []Suit{
		{
			name: "p = 5, q = 1",
			in:   tree,
			p: &TreeNode{
				Val: 5,
			},
			q: &TreeNode{
				Val: 1,
			},
			expected: tree,
		},
		{
			name: "p = 5, q = 4",
			in:   tree,
			p: &TreeNode{
				Val: 5,
			},
			q: &TreeNode{
				Val: 4,
			},
			expected: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 6,
				},
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 7,
					},
					Right: &TreeNode{
						Val: 4,
					},
				},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res := lowestCommonAncestor(testCase.in, testCase.p, testCase.q)
			assert.Equal(t, testCase.expected, res)
		})
	}
}
