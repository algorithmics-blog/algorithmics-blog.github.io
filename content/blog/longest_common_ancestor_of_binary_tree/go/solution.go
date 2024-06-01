package longest_common_ancestor_of_binary_tree

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var lca *TreeNode
	var findNode func(node *TreeNode) bool

	findNode = func(node *TreeNode) bool {
		if node == nil || q == nil || p == nil {
			return false
		}

		left := findNode(node.Left)
		right := findNode(node.Right)

		if left && right {
			lca = node
			return true
		}

		if node.Val == p.Val || node.Val == q.Val {
			if left || right {
				lca = node
			}
			return true
		}

		return left || right
	}

	findNode(root)

	return lca
}
