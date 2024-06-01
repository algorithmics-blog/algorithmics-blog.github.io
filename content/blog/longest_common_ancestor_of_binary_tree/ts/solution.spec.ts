import { lowestCommonAncestor, TreeNode } from "./solution";

type Suit = {
    name: string;
    in: TreeNode
    expected?: TreeNode
    p: TreeNode
    q: TreeNode
}

const tree: TreeNode = {
    val: 3,
    left: {
        val: 5,
        left: {
            val: 6,
        },
        right: {
            val: 2,
            left: {
                val: 7,
            },
            right: {
                val: 4,
            },
        },
    },
    right: {
        val: 1,
        left: {
            val: 0,
        },
        right: {
            val: 8,
        },
    },
}

const suits: Suit[] = [
    {
        name: "p = 5, q = 1",
        in: tree,
        p: {
            val: 5,
        },
        q: {
            val: 1,
        },
        expected: tree,
    },
    {
        name: "p = 5, q = 4",
        in: tree,
        p: {
            val: 5,
        },
        q: {
            val: 4,
        },
        expected: {
            val: 5,
            left: {
                val: 6,
            },
            right: {
                val: 2,
                left: {
                    val: 7,
                },
                right: {
                    val: 4,
                },
            },
        },
    },
]


describe('lowestCommonAncestor', () => {
    suits.forEach(suit => {
        test(suit.name, () => {
            const res = lowestCommonAncestor(suit.in, suit.p, suit.q)
            expect(res).toEqual(suit.expected);
        })
    })
})
