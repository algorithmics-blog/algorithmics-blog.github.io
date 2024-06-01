import { dailyTemperatures } from "./solution";

type Suit = {
	name: string
	in: number[]
	out: number[]
}

const suits: Suit[] = [
	{
		name: "[73,74,75,71,69,72,76,73]",
		in:   [73, 74, 75, 71, 69, 72, 76, 73],
		out:  [1, 1, 4, 2, 1, 1, 0, 0],
	},
	{
		name: "[30,40,50,60]",
		in:   [30, 40, 50, 60],
		out:  [1, 1, 1, 0],
	},
	{
		name: "[30,60,90]",
		in:   [30, 60, 90],
		out:  [1, 1, 0],
	},
]


describe('dailyTemperatures', () => {
	suits.forEach(suit => {
		test(suit.name, () => {
			const res = dailyTemperatures(suit.in)
			console.log(res)
			expect(res).toEqual(suit.out);
		})
	})
})