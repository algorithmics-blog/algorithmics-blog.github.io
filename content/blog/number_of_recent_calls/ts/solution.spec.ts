import { RecentCounter } from "./solution";

describe('RecentCounter', () => {
	test('[1], [100], [3001], [3002]', () => {
		const counter = new RecentCounter()

		counter.ping(1)
		counter.ping(100)
		counter.ping(3001)
		counter.ping(3002)
		const res = counter.ping(3003)

		expect(res).toEqual(4);
	})

	test('[1], [100], [3001]', () => {
		const counter = new RecentCounter()

		counter.ping(1)
		counter.ping(100)
		const res = counter.ping(3001)

		expect(res).toEqual(3);
	})
})
