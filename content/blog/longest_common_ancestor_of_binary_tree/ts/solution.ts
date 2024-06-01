export type TreeNode = {
    val: number
    left?: TreeNode
    right?: TreeNode
}

export const lowestCommonAncestor = (root?: TreeNode, p?: TreeNode, q?: TreeNode): TreeNode | null => {
    let lca: TreeNode | null = null;

    const findNode = (node?: TreeNode): boolean => {
        if (!node || !p || !q) {
            return false
        }

        const left = findNode(node.left)
        const right = findNode(node.right)

        if (left && right) {
            lca = node
            return true
        }

        if (node.val == p.val || node.val == q.val) {
            if (left || right) {
                lca = node
            }
            return true
        }

        return left || right;
    }

    findNode(root)

    return lca
}
