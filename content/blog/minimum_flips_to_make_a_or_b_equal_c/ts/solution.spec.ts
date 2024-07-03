import {minFlips} from "./solution";

type Suit = {
    name: string
    a: number
    b: number
    c: number
    expected: number
}

const suits: Suit[] = [
    {
        name: "a = 2, b = 6, c = 5",
        a: 2,
        b: 6,
        c: 5,
        expected: 3,
    },
    {
        name: "a = 4, b = 2, c = 7",
        a: 4,
        b: 2,
        c: 7,
        expected: 1,
    },
    {
        name: "a = 1, b = 2, c = 3",
        a: 1,
        b: 2,
        c: 3,
        expected: 0,
    }
]

describe('minFlips', () => {
    suits.forEach(suit => {
        test(suit.name, () => {
            const res = minFlips(suit.a, suit.b, suit.c)
            expect(res).toEqual(suit.expected);
        })
    })
})
