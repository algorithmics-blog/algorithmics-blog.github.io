import { plusOne } from "./solution";

type Suit = {
    name: string
    digits: number[]
    out: number[]
}

const suits: Suit[] = [
    {
        name: "[1,2,3]",
        digits: [1, 2, 3],
        out: [1, 2, 4],
    },
    {
        name: "[2,2,9]",
        digits: [2, 2, 9],
        out: [2, 3, 0],
    },
    {
        name: "[9]",
        digits: [9],
        out: [1, 0],
    },
    {
        name: "[1]",
        out: [2],
        digits: [1],
    },
]

describe('isSameAfterReversals', () => {
    suits.forEach(suit => {
        test(suit.name, () => {
            const res = plusOne(suit.digits)
            expect(res).toEqual(suit.out);
        })
    })
})
