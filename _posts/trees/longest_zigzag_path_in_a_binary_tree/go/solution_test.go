package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrie_Insert(t *testing.T) {
	testCases := []struct {
		name     string
		in       *TreeNode
		expected int
	}{
		{
			name: "3",
			in: &TreeNode{
				Right: &TreeNode{
					Left: &TreeNode{},
					Right: &TreeNode{
						Left: &TreeNode{
							Right: &TreeNode{
								Right: &TreeNode{},
							},
						},
						Right: &TreeNode{},
					},
				},
			},
			expected: 3,
		},
		{
			name: "4",
			in: &TreeNode{
				Left: &TreeNode{
					Right: &TreeNode{
						Left: &TreeNode{
							Right: &TreeNode{},
						},
						Right: &TreeNode{},
					},
				},
				Right: &TreeNode{},
			},
			expected: 4,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res := longestZigZag(testCase.in)

			assert.Equal(t, testCase.expected, res)
		})
	}
}
