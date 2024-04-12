import {maxOperations} from "./solution";

type Suit = {
    name: string
    nums: number[]
    k: number
    out: number
}

const suits: Suit[] = [
    {
        name: "[1,2,3,4]",
        nums: [1, 2, 3, 4],
        k: 5,
        out: 2,
    },
    {
        name: "[3,1,3,4,3]",
        nums: [3, 1, 3, 4, 3],
        k: 6,
        out: 1,
    },
]

describe('maxOperations', () => {
    suits.forEach(suit => {
        test(suit.name, () => {
            const res = maxOperations(suit.nums, suit.k)
            expect(res).toEqual(suit.out);
        })
    })
})
