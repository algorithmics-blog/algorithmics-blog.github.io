import { hammingWeightPopCount } from "./pop_count";
import { hammingWeight } from "./solution";

type Suit = {
    name: string
    num: number
    out: number
}

const suits: Suit[] = [
    {
        name: "0",
        num:  0,
        out:  0,
    },
    {
        name: "1",
        num:  1,
        out:  1,
    },
    {
        name: "11",
        num:  11,
        out:  3,
    },
    {
        name: "128",
        num:  128,
        out:  1,
    },
    {
        name: "2147483645",
        num:  2147483645,
        out:  30,
    },
]

describe('hammingWeightPopCount', () => {
    suits.forEach(suit => {
        test(suit.name, () => {
            const res = hammingWeightPopCount(suit.num)
            expect(res).toEqual(suit.out);
        })
    })
})

describe('hammingWeight', () => {
    suits.forEach(suit => {
        test(suit.name, () => {
            const res = hammingWeight(suit.num)
            expect(res).toEqual(suit.out);
        })
    })
})