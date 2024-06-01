import { reverseVowels } from "./solution";

type Suit = {
	name: string
	in:   string
	out:  string
}

const suits: Suit[] = [
	{
		name: "algorithmics",
		in:   "algorithmics",
		out:  "ilgirothmacs",
	},
	{
		name: "AlgorithmIcs",
		in:   "AlgorithmIcs",
		out:  "IlgirothmAcs",
	},
	{
		name: "bae",
		in:   "bae",
		out:  "bea",
	},
	{
		name: "aeb",
		in:   "aeb",
		out:  "eab",
	},
	{
		name: "b",
		in:   "b",
		out:  "b",
	},
	{
		name: "ab",
		in:   "ab",
		out:  "ab",
	},
	{
		name: "ba",
		in:   "ba",
		out:  "ba",
	},
]

describe('generateParenthesis', () => {
	suits.forEach(suit => {
		test(suit.name, () => {
			const res = reverseVowels(suit.in)
			expect(res).toEqual(suit.out);
		})
	})
})
