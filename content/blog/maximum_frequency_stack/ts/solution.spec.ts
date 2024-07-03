import { FreqStack } from './solution';

type Suit = {
    name: string
    input: number[]
    out: number[]
}

const testCases: Suit[] = [
    {
        name: '[5,7,5,7,4,5]',
        input: [5, 7, 5, 7, 4, 5],
        out: [5, 7, 5, 4]
    },
];

describe('FreqStack', () => {
    testCases.forEach(c => {
        test(c.name, () => {
            const stack = new FreqStack();
            c.input.forEach(element => stack.push(element));

            const res: number[] = [];
            for (let i = 0; i < c.out.length; i++) {
                res.push(stack.pop());
            }

            expect(res).toEqual(c.out);
        });
    });
});
