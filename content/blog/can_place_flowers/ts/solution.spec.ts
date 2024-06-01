import { canPlaceFlowers } from "./solution";

type TestSuit = {
    name: string
    flowerbed: number[]
    n: number
    expected: boolean
}

const suits: TestSuit[] = [
    {
        name: "[1,0,0,0,1],n=1",
        flowerbed: [1, 0, 0, 0, 1],
        n: 1,
        expected: true,
    },
    {
        name: "[1,0,0,0,1],n=2",
        flowerbed: [1, 0, 0, 0, 1],
        n: 2,
        expected: false,
    },
    {
        name: "[1],n=0",
        flowerbed: [1],
        n: 0,
        expected: true,
    },
    {
        name: "[0],n=1",
        flowerbed: [0],
        n: 1,
        expected: true,
    },
    {
        name: "[1],n=1",
        flowerbed: [1],
        n: 1,
        expected: false,
    },
    {
        name: "[0,1,0],n=1",
        flowerbed: [0, 1, 0],
        n: 1,
        expected: false,
    },
]

describe('maxProfitMinMax', () => {
    suits.forEach(suit => {
        test(suit.name, () => {
            const res = canPlaceFlowers(suit.flowerbed, suit.n)
            expect(res).toEqual(suit.expected);
        })
    })
})
