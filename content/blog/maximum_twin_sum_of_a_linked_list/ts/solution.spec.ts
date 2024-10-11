import { ListNode } from './ListNode';
import { pairSum } from './solution';
import { pairSumStack } from './solution_half_stack';
import { pairSumSlice } from './solution_slice';

type Suit = {
    name: string
    values: number[]
    out: number
}

const testCases: Suit[] = [
    {
        name: '[5,4,2,1]',
        values: [5, 4, 2, 1],
        out: 6,
    },
    {
        name: '[4,2,2,3]',
        values: [4, 2, 2, 3],
        out: 7,
    },
    {
        name: '[1,100000]',
        values: [1, 100000],
        out: 100001,
    },
]

function valuesSliceToList(values: number[]): ListNode {
    let cur = new ListNode(values[0])
    const head = cur

    for (let i = 1; i < values.length; i++) {
        const next = new ListNode(values[i])
        cur.next = next
        cur = next
    }

    return head
}

describe('pairSum', () => {
    testCases.forEach(testCase => {
        test(testCase.name, () => {
            const head = valuesSliceToList(testCase.values)

            const res = pairSum(head)
            expect(res).toEqual(testCase.out);
        })
    })
})

describe('pairSumStack', () => {
    testCases.forEach(testCase => {
        test(testCase.name, () => {
            const head = valuesSliceToList(testCase.values)

            const res = pairSumStack(head)
            expect(res).toEqual(testCase.out);
        })
    })
})

describe('pairSumSlice', () => {
    testCases.forEach(testCase => {
        test(testCase.name, () => {
            const head = valuesSliceToList(testCase.values)

            const res = pairSumSlice(head)
            expect(res).toEqual(testCase.out);
        })
    })
})