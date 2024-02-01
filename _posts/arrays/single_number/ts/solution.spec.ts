import { singleNumber } from "./solution";
import { singleNumberMap } from "./solution_map";

type RemoveDuplicatesSuit = {
    name: string
    nums: number[]
    expRes: number
}

const suits: RemoveDuplicatesSuit[] = [
    {
        name:   "[2,2,1]",
        nums:   [2, 2, 1],
        expRes: 1,
    },
    {
        name:   "[4,1,2,1,2]",
        nums:   [4, 1, 2, 1, 2],
        expRes: 4,
    },
    {
        name:   "[1]",
        nums:   [1],
        expRes: 1,
    },
]

describe('singleNumber', () => {
    suits.forEach(suit => {
        test(suit.name, () => {
            const res = singleNumber(suit.nums)
            expect(res).toEqual(suit.expRes);
        })
    })
})

describe('singleNumberMap', () => {
    suits.forEach(suit => {
        test(suit.name, () => {
            const res = singleNumberMap(suit.nums)
            expect(res).toEqual(suit.expRes);
        })
    })
})