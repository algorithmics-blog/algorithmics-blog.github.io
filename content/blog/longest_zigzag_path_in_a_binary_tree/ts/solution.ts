export type TreeNode = {
    val: number
    left?: TreeNode
    right?: TreeNode
}

type DFSResult = {
    leftPathLength: number
    rightPathLength: number
    maxLength: number
}

const dfs = (root?: TreeNode): DFSResult => {
    if (!root) {
        return {
            leftPathLength: 0,
            rightPathLength: 0,
            maxLength: 0,
        }
    }

    let leftPathLength = 0
    let rightPathLength = 0
    let leftMaxLength = 0
    let rightMaxLength = 0

    if (root.left) {
        const leftDFSRes = dfs(root.left)
        leftPathLength = leftDFSRes.rightPathLength + 1
        leftMaxLength = leftDFSRes.maxLength
    }

    if (root.right) {
        const rightDFSRes = dfs(root.right)
        rightPathLength = rightDFSRes.leftPathLength + 1
        rightMaxLength = rightDFSRes.maxLength
    }

    const maxLength = Math.max(Math.max(leftMaxLength, rightMaxLength), Math.max(leftPathLength, rightPathLength))

    return {
        leftPathLength,
        rightPathLength,
        maxLength,
    }
}

export const longestZigZag = (root?: TreeNode): number => {
    const { maxLength} = dfs(root)
    return maxLength
}