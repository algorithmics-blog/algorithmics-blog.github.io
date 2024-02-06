import { reverseWords } from './solution';

type Suit = {
    name: string
    in: string
    expected: string
}

const suits: Suit[] = [
    {
        name: "the sky is blue",
        in: "the sky is blue",
        expected: "blue is sky the",
    },
    {
        name: "  hello world  ",
        in: "  hello world  ",
        expected: "world hello",
    },
    {
        name: "a good   example",
        in: "a good   example",
        expected: "example good a",
    },
]

describe('reverseWords', () => {
    suits.forEach(suit => {
        test(suit.name, () => {
            const res = reverseWords(suit.in)
            expect(res).toEqual(suit.expected);
        })
    })
})