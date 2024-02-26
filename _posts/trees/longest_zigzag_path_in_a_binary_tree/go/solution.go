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
	var leftPathLength, rightPathLength, maxLength, maxPathFromLeft, maxPathFromRight int

	if root.Left != nil {
		var childRightZigZag int
		_, childRightZigZag, maxPathFromLeft = dfs(root.Left)
		leftPathLength = childRightZigZag + 1
	}

	if root.Right != nil {
		var childLeftZigZag int
		childLeftZigZag, _, maxPathFromRight = dfs(root.Right)

		rightPathLength = childLeftZigZag + 1
	}

	maxLength = max(max(maxPathFromLeft, maxPathFromRight), max(leftPathLength, rightPathLength))

	return leftPathLength, rightPathLength, maxLength
}
