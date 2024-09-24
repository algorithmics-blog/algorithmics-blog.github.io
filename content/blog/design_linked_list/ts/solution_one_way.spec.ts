import { MyLinkedList } from "./solution_one_way";

const buildValues = (list: MyLinkedList): number[] => {
    const values = []

    let current = list.head

    while (current) {
        values.push(current.val)
        current = current.next
    }

    return values
}

describe('MyLinkedList', () => {
    let list: MyLinkedList

    beforeEach(() => {
        list = new MyLinkedList()
    })

    test('scenario 1', () => {
        list.addAtHead(7)
        list.addAtHead(2)
        list.addAtHead(1)
        list.addAtIndex(3, 0)
        list.deleteAtIndex(2)
        list.addAtHead(6)
        list.addAtTail(4)
        list.addAtHead(4)
        list.addAtIndex(5, 0)
        list.addAtHead(6)

        const values = buildValues(list)

        expect(values).toEqual([6, 4, 6, 1, 2, 0, 0, 4])
        expect(list.size).toEqual(8)
    })

    test('scenario 2', () => {
        list.addAtHead(1)
        list.addAtTail(3)
        list.addAtIndex(1, 2)
        list.deleteAtIndex(0)


    })

    test('scenario 3', () => {
        list.addAtHead(1)
        list.addAtTail(3)
        list.addAtIndex(1, 2)

        expect(list.get(1)).toEqual(2)

        list.deleteAtIndex(1)

        expect(list.get(1)).toEqual(3)
        expect(list.get(3)).toEqual(-1)

        list.deleteAtIndex(3)
        list.deleteAtIndex(0)

        expect(list.get(0)).toEqual(3)

        list.deleteAtIndex(0)

        expect(list.get(0)).toEqual(-1)

        const values = buildValues(list)

        expect(values).toEqual([])
        expect(list.size).toEqual(0)
    })
})
