import { findCircleNum } from './solution';

type Suit = {
    name: string
    isConnected: number[][]
    expected: number
}

const suits: Suit[] = [
    {
        name: "example_1",
        isConnected: [
            [1, 1, 0],
            [1, 1, 0],
            [0, 0, 1],
        ],
        expected: 2,
    },
    {
        name: "example_2",
        isConnected: [
            [1, 0, 0],
            [0, 1, 0],
            [0, 0, 1],
        ],
        expected: 3,
    },
]

describe('findCircleNum', () => {
    suits.forEach(suit => {
        test(suit.name, () => {
            const res = findCircleNum(suit.isConnected)
            expect(res).toEqual(suit.expected);
        })
    })
})
