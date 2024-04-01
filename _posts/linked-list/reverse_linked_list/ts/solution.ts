export class ListNode {
    val: number
    next: ListNode | null

    constructor(val?: number, next?: ListNode | null) {
        this.val = (val === undefined ? 0 : val)
        this.next = (next === undefined ? null : next)
    }
}

export const reverseList = (head: ListNode | null): ListNode | null => {
    let prev: ListNode | null = null
    let curr = head

    while (curr) {
        const next = curr.next
        curr.next = prev
        prev = curr
        curr = next
    }

    return prev
}
