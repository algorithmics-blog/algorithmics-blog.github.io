import { Store } from "./solution";

describe('Store', () => {
	test("Insert duplicates", () => {
		const store = new Store()

		store.insert(1)
		store.insert(2)
		store.insert(1)

		expect(store.getValues()).toEqual([1, 2]);
	})

	test("Remove value", () => {
		const store = new Store()

		store.insert(1)
		store.insert(2)
		store.insert(3)
		store.remove(2)

		expect(store.getValues()).toEqual([1, 3]);
	})

	test("GetRandom value", () => {
		const store = new Store()

		store.insert(1)
		store.insert(2)
		store.insert(3)

		const value = store.getRandom()

		expect(store.getValues()).toEqual(expect.arrayContaining([value]));
	})
})
