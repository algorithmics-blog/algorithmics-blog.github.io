import { longestValidParenthesesWindowing } from "./solution_window";

type Suit = {
    name: string
    in: string
    expected: number
}

const suits: Suit[] = [
    {
        name: "(()",
        in: "(()",
        expected: 2,
    },
    {
        name: ")()())",
        in: ")()())",
        expected: 4,
    },
    {
        name: "",
        in: "",
        expected: 0,
    },
]

describe('longestValidParenthesesWindowing', () => {
    suits.forEach(suit => {
        test(suit.name, () => {
            const res = longestValidParenthesesWindowing(suit.in)
            expect(res).toEqual(suit.expected);
        })
    })
})
