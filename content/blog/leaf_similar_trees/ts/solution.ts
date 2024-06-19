export type TreeNode = {
    val: number
    left: TreeNode | null
    right: TreeNode | null
}

export function leafSimilar(root1: TreeNode | null, root2: TreeNode | null): boolean {
    if (!root1 || !root2) {
        return false
    }

    const first = getSequence(root1)
    const second = getSequence(root2)

    return first === second
}

function getSequence(root: TreeNode): string {
    let sequence = ""

    const dfs = (node: TreeNode): void => {
        if (!node.left && !node.right) {
            sequence = `${sequence}${node.val};`
        }

        if (node.left) {
            dfs(node.left)
        }

        if (node.right) {
            dfs(node.right)
        }
    }

    dfs(root)

    return sequence
}