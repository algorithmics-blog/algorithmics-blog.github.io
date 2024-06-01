import { productExceptSelf } from "./solution";

type Suit = {
    name: string
    in: number[]
    expected: number[]
}

const suits: Suit[] = [
    {
        name: "[1,2,3,4]",
        in: [1, 2, 3, 4],
        expected: [24, 12, 8, 6],
    },
    {
        name: "[-1,1,0,-3,3]",
        in: [-1, 1, 0, -3, 3],
        expected: [0, 0, 9, 0, 0],
    },
    {
        name: "[1, 2]",
        in: [1, 2],
        expected: [2, 1],
    },
]

describe('productExceptSelf', () => {
    suits.forEach(suit => {
        test(suit.name, () => {
            const res = productExceptSelf(suit.in)
            expect(res).toEqual(suit.expected);
        })
    })
})
