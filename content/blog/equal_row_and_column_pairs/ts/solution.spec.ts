import { equalPairs } from "./solution";

type Suit = {
    name: string
    in: number[][]
    expected: number
}

const suits: Suit[] = [
    {
        name: "[[3,2,1],[1,7,6],[2,7,7]]",
        in: [
            [3, 2, 1],
            [1, 7, 6],
            [2, 7, 7]
        ],
        expected: 1,
    },
    {
        name: "[[3,1,2,2],[1,4,4,5],[2,4,2,2],[2,4,2,2]]",
        in: [
            [3, 1, 2, 2],
            [1, 4, 4, 5],
            [2, 4, 2, 2],
            [2, 4, 2, 2]
        ],
        expected: 3,
    },
    {
        name: "[[5]]",
        in: [[5]],
        expected: 1,
    }
]

describe('equalPairs', () => {
    suits.forEach(suit => {
        test(suit.name, () => {
            const res = equalPairs(suit.in)
            expect(res).toEqual(suit.expected);
        })
    })
})
