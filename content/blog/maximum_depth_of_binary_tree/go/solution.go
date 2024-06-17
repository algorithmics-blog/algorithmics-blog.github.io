package maximum_depth_of_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	return dfs(root, 0)
}

func dfs(root *TreeNode, level int) int {
	if root == nil {
		return level
	}

	leftDepth := dfs(root.Left, level+1)
	rightDepth := dfs(root.Right, level+1)

	if leftDepth > rightDepth {
		return leftDepth
	}

	return rightDepth
}
