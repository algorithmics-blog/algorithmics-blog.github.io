import {guessNumber} from "./solution";

type Suit = {
    name: string
    n: number
    secret: number
}

const suits: Suit[] = [
    {
        name: "n = 10, secret = 3",
        n: 10,
        secret: 3,
    },
    {
        name: "n = 1, secret = 1",
        n: 1,
        secret: 1,
    },
    {
        name: "n = 10, secret = 10",
        n: 10,
        secret: 10,
    },
    {
        name: "n = 10, secret = 5",
        n: 10,
        secret: 5,
    },
]

describe('guessNumber', () => {
    suits.forEach(suit => {
        test(suit.name, () => {
            const guess = (num: number): number => {
                if (num > suit.secret) {
                    return -1
                } else if (num < suit.secret) {
                    return 1
                }

                return 0
            }

            const res = guessNumber(suit.n, guess)
            expect(res).toEqual(suit.secret);
        })
    })
})
