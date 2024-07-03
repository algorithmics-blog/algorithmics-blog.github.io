export type TreeNode = {
    val: number
    left: TreeNode | null
    right: TreeNode | null
}

export function maxDepth(root: TreeNode | null): number {
    return dfs(root, 0)
}

function dfs(node: TreeNode | null, level: number): number {
    if (!node) {
        return level
    }

    const leftDepth = dfs(node.left, level + 1)
    const rightDepth = dfs(node.right, level + 1)

    return Math.max(leftDepth, rightDepth)
}