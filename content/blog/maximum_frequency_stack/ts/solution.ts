class Item {
    freq: number;
    seq: number;
    val: number;

    constructor(freq: number, seq: number, val: number) {
        this.freq = freq;
        this.seq = seq;
        this.val = val;
    }

    more(second: Item): boolean {
        if (this.freq !== second.freq) {
            return this.freq > second.freq;
        }
        return this.seq > second.seq;
    }
}

class Heap {
    items: Item[] = [];

    add(newItem: Item): void {
        this.items.push(newItem);
        this.upElement(this.items.length - 1);
    }

    pop(): Item | null {
        if (this.items.length === 0) {
            return null;
        }

        const res = this.items[0];
        if (this.items.length > 1) {
            this.items[0] = this.items.pop()!;
            this.heapify(0);
        } else {
            this.items.pop();
        }

        return res;
    }

    private heapify(index: number): void {
        const left = index * 2 + 1;
        const right = left + 1;
        let largest = index;

        if (left < this.items.length && this.items[left].more(this.items[largest])) {
            largest = left;
        }
        if (right < this.items.length && this.items[right].more(this.items[largest])) {
            largest = right;
        }

        if (largest !== index) {
            [this.items[index], this.items[largest]] = [this.items[largest], this.items[index]];
            this.heapify(largest);
        }
    }

    private upElement(index: number): void {
        const parent = Math.floor((index - 1) / 2);

        if (parent >= 0 && this.items[index].more(this.items[parent])) {
            [this.items[index], this.items[parent]] = [this.items[parent], this.items[index]];
            this.upElement(parent);
        }
    }
}

export class FreqStack {
    private h: Heap;
    private freqElem: Map<number, number>;
    private curSeq: number;

    constructor() {
        this.h = new Heap();
        this.freqElem = new Map();
        this.curSeq = 0;
    }

    push(val: number): void {
        this.curSeq++;
        this.freqElem.set(val, (this.freqElem.get(val) || 0) + 1);
        this.h.add(new Item(this.freqElem.get(val)!, this.curSeq, val));
    }

    pop(): number {
        const item = this.h.pop();
        if (item) {
            this.freqElem.set(item.val, this.freqElem.get(item.val)! - 1);
            if (this.freqElem.get(item.val) === 0) {
                this.freqElem.delete(item.val);
            }
            return item.val;
        }
        return -1; // or throw an error
    }
}