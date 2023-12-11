import { asteroidCollision } from "./solution";
import { asteroidCollisionTwoIdx } from "./solution_two_idx";

type Suit = {
    name: string
    in: number[]
    out: number[]
}

const suits: Suit[] = [
    {
        name: "[5,10,-5]",
        in: [5, 10, -5],
        out: [5, 10],
    },
    {
        name: "[8,-8]",
        in: [8, -8],
        out: [],
    },
    {
        name: "[10,2,-5]",
        in: [10, 2, -5],
        out: [10],
    },
    {
        name: "[-8,8]",
        in: [-8, 8],
        out: [-8, 8],
    },
    {
        name: "[-8,-9]",
        in: [-8, -9],
        out: [-8, -9],
    },
    {
        name: "[8,9]",
        in: [8, 9],
        out: [8, 9],
    },
]

describe('asteroidCollision', () => {
    suits.forEach(suit => {
        test(suit.name, () => {
            const res = asteroidCollision(suit.in)
            expect(res).toEqual(suit.out);
        })
    })
})

describe('asteroidCollisionTwoIdx', () => {
    suits.forEach(suit => {
        test(suit.name, () => {
            const res = asteroidCollisionTwoIdx(suit.in)
            expect(res).toEqual(suit.out);
        })
    })
})
