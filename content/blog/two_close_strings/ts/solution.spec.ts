import {closeStrings} from "./solution";

type TestSuit = {
    name: string
    str1: string
    str2: string
    expected: boolean
}

const suits: TestSuit[] = [
    {
        name: "abc, bca",
        str1: "abc",
        str2: "bca",
        expected: true,
    },
    {
        name: "a, aa",
        str1: "a",
        str2: "aa",
        expected: false,
    },
    {
        name: "cabbba, abbccc",
        str1: "cabbba",
        str2: "abbccc",
        expected: true,
    },
    {
        name: "cabbba, bbcccd",
        str1: "cabbba",
        str2: "bbcccd",
        expected: false,
    },
    {
        name: "abbzzca, babzzcz",
        str1: "abbzzca",
        str2: "babzzcz",
        expected: false,
    },
]

describe('kidsWithCandies', () => {
    suits.forEach(suit => {
        test(suit.name, () => {
            const res = closeStrings(suit.str1, suit.str2)
            expect(res).toEqual(suit.expected);
        })
    })
})
