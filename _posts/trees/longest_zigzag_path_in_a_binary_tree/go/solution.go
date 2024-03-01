package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func longestZigZag(root *TreeNode) int {
	_, _, m := dfs(root)

	return m
}

func dfs(root *TreeNode) (int, int, int) {
	var leftPathLength, rightPathLength, leftMaxLength, rightMaxLength int
	if root == nil {
		return leftPathLength, rightPathLength, 0
	}

	if root.Left != nil {
		var childRightZigZag int
		_, childRightZigZag, leftMaxLength = dfs(root.Left)
		leftPathLength = childRightZigZag + 1
	}

	if root.Right != nil {
		var childLeftZigZag int
		childLeftZigZag, _, rightMaxLength = dfs(root.Right)

		rightPathLength = childLeftZigZag + 1
	}

	maxLength := max(max(leftMaxLength, rightMaxLength), max(leftPathLength, rightPathLength))
	return leftPathLength, rightPathLength, maxLength
}
