import { fizzBuzz } from "./solution";

type Suit = {
    name: string
    n: number
    expRes: string[]
}

const suits: Suit[] = [
    {
        name: "3",
        n: 3,
        expRes: ["1", "2", "Fizz"],
    },
    {
        name: "5",
        n: 5,
        expRes: ["1", "2", "Fizz", "4", "Buzz"],
    },
    {
        name: "15",
        n: 15,
        expRes: ["1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"],
    },
]

describe('fizzBuzz', () => {
    suits.forEach(suit => {
        test(suit.name, () => {
            const res = fizzBuzz(suit.n)
            expect(res).toEqual(suit.expRes);
        })
    })
})
