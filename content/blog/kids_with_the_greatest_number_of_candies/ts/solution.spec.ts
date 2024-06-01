import { kidsWithCandies } from "./solution";

type TestSuit = {
    name: string
    candies: number[]
    extraCandies: number
    expected: boolean[]
}

const suits: TestSuit[] = [
    {
        name: "[2,3,5,1,3], extraCandies=3",
        candies: [2, 3, 5, 1, 3],
        extraCandies: 3,
        expected: [true, true, true, false, true],
    },
    {
        name: "[4,2,1,1,2], extraCandies=1",
        candies: [4, 2, 1, 1, 2],
        extraCandies: 1,
        expected: [true, false, false, false, false],
    },
    {
        name: "[12,1,12], extraCandies=10",
        candies: [12, 1, 12],
        extraCandies: 10,
        expected: [true, false, true],
    },
    {
        name: "[2, 2], extraCandies=1",
        candies: [2, 2],
        extraCandies: 3,
        expected: [true, true],
    },
]

describe('kidsWithCandies', () => {
    suits.forEach(suit => {
        test(suit.name, () => {
            const res = kidsWithCandies(suit.candies, suit.extraCandies)
            expect(res).toEqual(suit.expected);
        })
    })
})
