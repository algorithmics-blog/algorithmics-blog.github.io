import { TreeNode, maxDepth } from './solution';

type Suit = {
    name: string
    input: TreeNode
    out: number
}

const testCases: Suit[] = [
    {
        name: '[3,9,20,null,null,15,7]',
        input: {
            val: 3,
            left: {
                val: 9,
                left: null,
                right: null,
            },
            right: {
                val: 20,
                left: {
                    val: 15,
                    left: null,
                    right: null,
                },
                right: {
                    val: 7,
                    left: null,
                    right: null,
                },
            }
        },
        out: 3
    },
    {
        name: '[1,null,2]',
        input: {
            val: 1,
            left: null,
            right: {
                val: 2,
                left: null,
                right: null,
            }
        },
        out: 2
    },
];

describe('maxDepth', () => {
    testCases.forEach(c => {
        test(c.name, () => {
            const res = maxDepth(c.input)
            expect(res).toEqual(c.out);
        });
    });
});
