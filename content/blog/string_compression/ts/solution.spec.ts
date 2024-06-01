import {compress} from "./solution";

type Suit = {
    name: string
    chars: string[]
    expected: number
}

const suits: Suit[] = [
    {
        name: `["a","a","b","b","c","c","c"]`,
        chars: ['a', 'a', 'b', 'b', 'c', 'c', 'c'],
        expected: 6,
    },
    {
        name: `["a"]`,
        chars: ['a'],
        expected: 1,
    },
    {
        name: `["a","b","b","b","b","b","b","b","b","b","b","b","b"]`,
        chars: ['a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'],
        expected: 4,
    },
    {
        name: `100 a`,
        chars: ['a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a'],
        expected: 4,
    },
]

describe('compress', () => {
    suits.forEach(suit => {
        test(suit.name, () => {
            const res = compress(suit.chars)
            expect(res).toEqual(suit.expected);
        })
    })
})