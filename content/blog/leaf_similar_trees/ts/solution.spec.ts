import { leafSimilar, TreeNode } from './solution';

type Suit = {
    name: string
    firstTree: TreeNode
    secondTree: TreeNode
    out: boolean
}

const testCases: Suit[] = [
    {
        name: 'root1 = [3,5,1,6,2,9,8,null,null,7,4], root2 = [3,5,1,6,7,4,2,null,null,null,null,null,null,9,8]',
        firstTree: {
            val: 3,
            left: {
                val: 5,
                left: {
                    val: 6,
                    left: null,
                    right: null,
                },
                right: {
                    val: 2,
                    left: {
                        val: 7,
                        left: null,
                        right: null,
                    },
                    right: {
                        val: 4,
                        left: null,
                        right: null,
                    }
                },
            },
            right: {
                val: 1,
                left: {
                    val: 9,
                    left: null,
                    right: null,
                },
                right: {
                    val: 8,
                    left: null,
                    right: null,
                },
            }
        },
        secondTree: {
            val: 3,
            left: {
                val: 5,
                left: {
                    val: 6,
                    left: null,
                    right: null,
                },
                right: {
                    val: 7,
                    left: null,
                    right: null,
                },
            },
            right: {
                val: 1,
                left: {
                    val: 4,
                    left: null,
                    right: null,
                },
                right: {
                    val: 2,
                    left: {
                        val: 9,
                        left: null,
                        right: null,
                    },
                    right: {
                        val: 8,
                        left: null,
                        right: null,
                    },
                },
            }
        },
        out: true
    },
    {
        name: 'root1 = [1,2,3], root2 = [1,3,2]',
        firstTree: {
            val: 1,
            left: {
                val: 2,
                left: null,
                right: null,
            },
            right: {
                val: 3,
                left: null,
                right: null,
            },
        },
        secondTree: {
            val: 1,
            left: {
                val: 3,
                left: null,
                right: null,
            },
            right: {
                val: 2,
                left: null,
                right: null,
            },
        },
        out: false
    },
];

describe('maxDepth', () => {
    testCases.forEach(c => {
        test(c.name, () => {
            const res = leafSimilar(c.firstTree, c.secondTree)
            expect(res).toEqual(c.out);
        });
    });
});
