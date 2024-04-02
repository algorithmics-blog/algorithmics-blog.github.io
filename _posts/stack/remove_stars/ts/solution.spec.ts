import { removeStars } from "./solution";

type Suit = {
    name: string
    in: string
    out: string
}

const suits: Suit[] = [
    {
        name: "leet**cod*e",
        in:   "leet**cod*e",
        out:  "lecoe",
    },
    {
        name: "erase*****",
        in:   "erase*****",
        out:  "",
    },
    {
        name: "nostars",
        in:   "nostars",
        out:  "nostars",
    },
    {
        name: "[empty]",
        in:   "",
        out:  "",
    },
]

describe('removeStars', () => {
    suits.forEach(suit => {
        test(suit.name, () => {
            const res = removeStars(suit.in)
            expect(res).toEqual(suit.out);
        })
    })
})
