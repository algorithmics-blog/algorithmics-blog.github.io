export class Store {
    private indexesMap: Map<number, number>
    private values: number[]

    constructor() {
        this.indexesMap = new Map()
        this.values = []
    }

    insert(value: number) {
        if (this.indexesMap.has(value)) {
            return
        }

        this.values.push(value)
        this.indexesMap.set(value, this.values.length - 1)
    }

    remove(value: number) {
        const index = this.indexesMap.get(value)
        if (!index) {
            return
        }

        const last = this.values.pop()
        if (!last) {
            return;
        }

        this.values[index] = last
        this.indexesMap.set(last, index)

        this.indexesMap.delete(value)
    }

    getRandom(): number {
        return this.values[Math.floor(Math.random() * this.values.length)]
    }

    // Method for tests
    getValues(): number[] {
        return this.values
    }
}
