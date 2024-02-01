import {eraseOverlapIntervals} from "./solution";

type TestSuit = {
    name: string
    in: number[][]
    expected: number
}

const suits: TestSuit[] = [
    {
        name: "[[1,2],[2,3],[3,4],[1,3]]",
        in: [[1, 2], [2, 3], [3, 4], [1, 3]],
        expected: 1,
    },
    {
        name: "[[1,2],[1,2],[1,2]]",
        in: [[1, 2], [1, 2], [1, 2]],
        expected: 2,
    },
    {
        name: "[[1,2],[2,3]]",
        in: [[1, 2], [2, 3]],
        expected: 0,
    },
]

describe('maxProfitMinMax', () => {
    suits.forEach(suit => {
        test(suit.name, () => {
            const res = eraseOverlapIntervals(suit.in)
            expect(res).toEqual(suit.expected);
        })
    })
})
