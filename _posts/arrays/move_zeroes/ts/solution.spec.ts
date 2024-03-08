import { moveZeroes } from "./solution";

type Suit = {
	name: string
	in: number[]
	out: number[]
}

const suits: Suit[] = [
	{
		name: "[0,1,0,3,12]",
		in:   [0, 1, 0, 3, 12],
		out:  [1, 3, 12, 0, 0],
	},
	{
		name: "[1,2,3,4,5]",
		in:   [1, 2, 3, 4, 5],
		out:  [1, 2, 3, 4, 5],
	},
	{
		name: "[0,0,0,0,0]",
		in:   [0, 0, 0, 0, 0],
		out:  [0, 0, 0, 0, 0],
	},
	{
		name: "[0]",
		in:   [0],
		out:  [0],
	},
	{
		name: "[1]",
		in:   [1],
		out:  [1],
	},
]

describe('predictPartyVictory', () => {
	suits.forEach(suit => {
		test(suit.name, () => {
			moveZeroes(suit.in)
			expect(suit.in).toEqual(suit.out);
		})
	})
})
