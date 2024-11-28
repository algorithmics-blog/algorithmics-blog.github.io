type MyLinkedListNode = {
    val: number;
    next: MyLinkedListNode | null;
}

export class MyLinkedList {
    head: MyLinkedListNode | null;
    size: number;

    constructor() {
        this.head = null;
        this.size = 0;
    }

    private findPrevNode(index: number): MyLinkedListNode | null {
        let curr = this.head;
        for (let i = 0; i < index - 1; i++) {
            if (curr) {
                curr = curr.next;
            }
        }

        return curr
    }

    get(index: number): number {
        if (index >= this.size || index < 0) {
            return -1;
        }

        let curr = this.head;
        for (let i = 0; i < index; i++) {
            if (curr) {
                curr = curr.next;
            }
        }

        return curr ? curr.val : -1;
    }

    addAtHead(val: number): void {
        this.addAtIndex(0, val);
    }

    addAtTail(val: number): void {
        this.addAtIndex(this.size, val);
    }

    addAtIndex(index: number, val: number): void {
        if (index > this.size || index < 0) {
            return;
        }

        this.size++;

        if (index === 0) {
            this.head = { val, next: this.head };
            return;
        }

        const prev = this.findPrevNode(index)

        if (prev) {
            prev.next = { val, next: prev.next };
        }
    }

    deleteAtIndex(index: number): void {
        if (index < 0 || index >= this.size) {
            return;
        }

        this.size--;

        if (index === 0 && this.head) {
            this.head = this.head.next;
            return;
        }

        const prev = this.findPrevNode(index)

        if (prev && prev.next) {
            prev.next = prev.next.next;
        }
    }
}