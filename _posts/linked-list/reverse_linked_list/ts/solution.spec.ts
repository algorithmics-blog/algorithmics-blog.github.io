import { ListNode, reverseList } from "./solution";

type Suit = {
    name: string
    list: ListNode | null
    out: ListNode | null
}

const suits: Suit[] = [
    {
        name: "[1, 2, 3, 4, 5]",
        list: new ListNode(1, new ListNode(2, new ListNode(3, new ListNode(4, new ListNode(5))))),
        out: new ListNode(5, new ListNode(4, new ListNode(3, new ListNode(2, new ListNode(1))))),
    },
    {
        name: "[1]",
        list: new ListNode(1),
        out: new ListNode(1),
    },
    {
        name: "[]",
        list: new ListNode(),
        out: new ListNode(),
    },
    {
        name: "null",
        list: null,
        out: null,
    }
]

describe('reverseList', () => {
    suits.forEach(suit => {
        test(suit.name, () => {
            const reversed = reverseList(suit.list)
            expect(reversed).toEqual(suit.out);
        })
    })
})
