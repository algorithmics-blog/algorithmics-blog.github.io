export class RecentCounter {
    calls: number[]

    constructor() {
        this.calls = []
    }

    ping(t: number): number {
        this.calls.push(t)

        while (this.calls[0] < t - 3000) {
            this.calls.shift()
        }

        return this.calls.length;
    }
}