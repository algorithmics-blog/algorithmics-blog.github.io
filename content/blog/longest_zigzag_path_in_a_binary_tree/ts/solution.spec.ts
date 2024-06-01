import {longestZigZag, TreeNode} from "./solution";

type Suit = {
    name: string;
    in: TreeNode
    expected: number
}

const suits: Suit[] = [
    {
        name: "3",
        in: {
            right: {
                left: {
					val: 1
				},
                right: {
                    left: {
                        right: {
                            right: {
								val: 2
							},
							val: 3
                        },
						val: 4
                    },
                    right: {
						val: 5
					},
					val: 6
                },
				val: 7
            },
			val: 8
        },
        expected: 3,
    },
    {
        name: "4",
        in: {
            left: {
                right: {
                    left: {
                        right: {
							val: 6
						},
						val: 5
                    },
                    right: {
						val: 17
					},
					val: 4
                },
				val: 2
            },
            right: {
				val: 3
			},
			val: 1
        },
        expected: 4,
    },
]

describe('longestZigZag', () => {
    suits.forEach(suit => {
        test(suit.name, () => {
            const res = longestZigZag(suit.in)
            expect(res).toEqual(suit.expected);
        })
    })
})
