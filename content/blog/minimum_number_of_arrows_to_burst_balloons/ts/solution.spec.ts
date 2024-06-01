import { findMinArrowShots } from "./solution";

type Suit = {
    name: string
    points: number[][]
    out: number
}

const suits: Suit[] = [
    {
        name: "[[10,16],[2,8],[1,6],[7,12]]",
        points: [
            [10, 16], [2, 8], [1, 6], [7, 12],
        ],
        out: 2,
    },
    {
        name: "[[1,2],[3,4],[5,6],[7,8]]",
        points: [
            [1, 2], [3, 4], [51, 6], [7, 8],
        ],
        out: 4,
    },
    {
        name: "[[1,2],[2,3],[3,4],[4,5]]",
        points: [
            [1, 2], [2, 3], [3, 4], [4, 5],
        ],
        out: 2,
    },
]

describe('maxOperations', () => {
    suits.forEach(suit => {
        test(suit.name, () => {
            const res = findMinArrowShots(suit.points)
            expect(res).toEqual(suit.out);
        })
    })
})
