import { maxVowels } from "./solution";

type Suit = {
    name: string
    s: string
    k: number
    expected: number
}

const suits: Suit[] = [
    {
        name: "abciiidef",
        s: "abciiidef",
        k: 3,
        expected: 3,
    },
    {
        name: "aeiou",
        s: "aeiou",
        k: 2,
        expected: 2,
    },
    {
        name: "leetcode",
        s: "leetcode",
        k: 3,
        expected: 2,
    },
]

describe('maxVowels', () => {
    suits.forEach(suit => {
        test(suit.name, () => {
            const res = maxVowels(suit.s, suit.k)
            expect(res).toEqual(suit.expected);
        })
    })
})
