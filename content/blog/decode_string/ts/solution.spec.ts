import { decodeString } from "./solution";

type Suit = {
	name: string
	in:   string
	out:  string
}

const suits: Suit[] = [
	{
		name: "3[a]2[bc]",
		in:   "3[a]2[bc]",
		out:  "aaabcbc",
	},
	{
		name: "3[a2[c]]",
		in:   "3[a2[c]]",
		out:  "accaccacc",
	},
	{
		name: "2[abc]3[cd]ef",
		in:   "2[abc]3[cd]ef",
		out:  "abcabccdcdcdef",
	},
]

describe('decodeString', () => {
	suits.forEach(suit => {
		test(suit.name, () => {
			const res = decodeString(suit.in)
			expect(res).toEqual(suit.out);
		})
	})
})
