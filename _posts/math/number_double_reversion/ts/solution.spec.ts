import { isSameAfterReversals } from "./solution";

type Suit = {
    name: string
    num: number
    out: boolean
}

const suits: Suit[] = [
    {
        name: "526",
        num: 526,
        out: true,
    },
    {
        name: "1800",
        num: 1800,
        out: false,
    },
    {
        name: "0",
        num: 0,
        out: true,
    },
]

describe('isSameAfterReversals', () => {
    suits.forEach(suit => {
        test(suit.name, () => {
            const res = isSameAfterReversals(suit.num)
            expect(res).toEqual(suit.out);
        })
    })
})
