package leaf_similar_trees

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	first := getSequence(root1)
	second := getSequence(root2)

	return first == second
}

func getSequence(root *TreeNode) string {
	var dfs func(node *TreeNode)

	sequence := ""

	dfs = func(node *TreeNode) {
		if node.Left == nil && node.Right == nil {
			fmt.Println(node.Val)
			sequence += strconv.Itoa(node.Val) + ";"
		}

		if node.Left != nil {
			dfs(node.Left)
		}

		if node.Right != nil {
			dfs(node.Right)
		}
	}

	dfs(root)

	return sequence
}
