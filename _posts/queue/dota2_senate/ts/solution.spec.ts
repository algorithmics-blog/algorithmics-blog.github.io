import { predictPartyVictory } from "./solution";

type Suit = {
	name: string
	in: string
	out: 'Radiant' | 'Dire'
}

const suits: Suit[] = [
	{
		name: 'RD',
		in:   'RD',
		out:  'Radiant',
	},
	{
		name: 'RDD',
		in:   'RDD',
		out:  'Dire',
	},
]

describe('predictPartyVictory', () => {
	suits.forEach(suit => {
		test(suit.name, () => {
			const winner = predictPartyVictory(suit.in)
			expect(winner).toEqual(suit.out);
		})
	})
})
