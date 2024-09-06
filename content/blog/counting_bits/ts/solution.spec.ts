import { countBitsPopCount } from "./pop_count";
import { countBits } from "./solution";

type Suit = {
    name: string
    num: number
    out: number[]
}

const suits: Suit[] = [
    {
        name: "0",
        num:  0,
        out:  [0],
    },
    {
        name: "1",
        num:  1,
        out:  [0, 1],
    },
    {
        name: "2",
        num:  2,
        out:  [0, 1, 1],
    },
    {
        name: "5",
        num:  5,
        out:  [0, 1, 1, 2, 1, 2],
    },
    {
        name: "7",
        num:  7,
        out:  [0, 1, 1, 2, 1, 2, 2, 3],
    },
]

describe('countBitsPopCount', () => {
    suits.forEach(suit => {
        test(suit.name, () => {
            const res = countBitsPopCount(suit.num)
            expect(res).toEqual(suit.out);
        })
    })
})

describe('countBits', () => {
    suits.forEach(suit => {
        test(suit.name, () => {
            const res = countBits(suit.num)
            expect(res).toEqual(suit.out);
        })
    })
})