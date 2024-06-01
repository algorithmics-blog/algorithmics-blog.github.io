import { reverse } from "./solution";

type Suit = {
	name: string
	num:  number
	out:  number
}

const suits: Suit[] = [
	{
		name: "123",
		num:  123,
		out:  321,
	},
	{
		name: "-123",
		num:  -123,
		out:  -321,
	},
	{
		name: "120",
		num:  120,
		out:  21,
	},
	{
		name: "2147483647",
		num:  2147483647,
		out:  0,
	},
	{
		name: "-2147483647",
		num:  -2147483647,
		out:  0,
	},
	{
		name: "-2147483412",
		num:  -2147483412,
		out:  -2143847412,
	},
]

describe('reverse', () => {
	suits.forEach(suit => {
		test(suit.name, () => {
			const res = reverse(suit.num)
			expect(res).toEqual(suit.out);
		})
	})
})
