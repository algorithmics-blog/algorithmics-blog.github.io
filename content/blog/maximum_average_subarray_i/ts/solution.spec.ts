import {findMaxAverage} from "./solution";

type Suit = {
    name: string
    nums: number[]
    k: number
    out: number
}

const suits: Suit[] = [
    {
        name: "[1,12,-5,-6,50,3]",
        nums: [1, 12, -5, -6, 50, 3],
        k: 4,
        out: 12.75,
    },
    {
        name: "[5]",
        nums: [5],
        k: 1,
        out: 5,
    },
]

describe('findMaxAverage', () => {
    suits.forEach(suit => {
        test(suit.name, () => {
            const res = findMaxAverage(suit.nums, suit.k);
            expect(res).toEqual(suit.out);
        })
    })
})
