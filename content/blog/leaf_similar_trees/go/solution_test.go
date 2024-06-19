package leaf_similar_trees

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestSuit struct {
	name       string
	firstTree  *TreeNode
	secondTree *TreeNode
	out        bool
}

func Test_leafSimilar(t *testing.T) {
	testCases := []TestSuit{
		{
			name: "root1 = [3,5,1,6,2,9,8,null,null,7,4], root2 = [3,5,1,6,7,4,2,null,null,null,null,null,null,9,8]",
			firstTree: &TreeNode{
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
						Val: 9,
					},
					Right: &TreeNode{
						Val: 8,
					},
				},
			},
			secondTree: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 6,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
				Right: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 9,
						},
						Right: &TreeNode{
							Val: 8,
						},
					},
				},
			},
			out: true,
		},
		{
			name: "root1 = [1,2,3], root2 = [1,3,2]",
			firstTree: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			secondTree: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 3,
				},
				Right: &TreeNode{
					Val: 2,
				},
			},
			out: false,
		},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			res := leafSimilar(c.firstTree, c.secondTree)
			assert.Equal(t, c.out, res)
		})
	}
}
