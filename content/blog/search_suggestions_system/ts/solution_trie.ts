class Trie {
	word: string = '';
	sort: string[] = [];
	children: Map<string, Trie> = new Map();

	insert(word: string): void {
		this._insert(word, [...word]);
	}

	private _insert(origin: string, word: string[]): void {
		if (word.length === 0) {
			this.word = origin;
			return;
		}

		const char = word[0];
		let child = this.children.get(char);

		if (!child) {
			child = new Trie();
			this.children.set(char, child);
			this.sort.push(char);
		}

		child._insert(origin, word.slice(1));
	}

	closest(prefix: string[], count: number): string[][] {
		const res: string[][] = [];
		let excludeChar: string | null = null;

		if (prefix.length > 0) {
			excludeChar = prefix[0];
		}

		res.push(this.exactWords(excludeChar, count));

		if (prefix.length === 0) {
			return res;
		}

		const child = this.children.get(prefix[0]);
		if (!child) {
			res[0] = this.exactWords(null, count);
			return res.concat(this.gd(prefix));
		}

		return res.concat(child.closest(prefix.slice(1), count));
	}

	private exactWords(excludeChar: string | null, n: number): string[] {
		let res: string[] = [];
		if (this.word.length > 0) {
			res.push(this.word);
			n--;
		}

		if (n === 0) {
			return res;
		}

		for (const childChar of this.sort) {
			if (excludeChar !== null && childChar === excludeChar) {
				continue;
			}

			const childWords = this.children.get(childChar)!.exactWords(null, n);
			if (childWords.length >= n) {
				return res.concat(childWords.slice(0, n));
			}

			res = res.concat(childWords);
			n -= childWords.length;

			if (n === 0) {
				break;
			}
		}

		return res;
	}

	gd(prefix: string[]): string[][] {
		const res: string[][] = [];
		for (let i = 0; i < prefix.length; i++) {
			res.push([]);
		}
		return res;
	}
}

export function suggestedProductsTrie(products: string[], searchWord: string): string[][] {
	const root = new Trie();

	products.sort();

	for (const p of products) {
		root.insert(p);
	}

	const runeSearch = [...searchWord];
	const firstChar = runeSearch[0];
	const worldStart = root.children.get(firstChar);

	if (!worldStart) {
		return root.gd(runeSearch);
	}

	const suggestions = worldStart.closest(runeSearch.slice(1), 3);
	const res: string[][] = new Array(suggestions.length);
	let prev: string[] = [];
	let i = suggestions.length - 1;

	if (i >= 0) {
		prev = suggestions[i].slice();
		res[i] = prev;
		i--;
	}

	for (; i >= 0; i--) {
		let newItem = suggestions[i].slice();

		if (prev.length === 0) {
			res[i] = newItem;
			prev = newItem.slice();
			continue;
		} else if (newItem.length === 0) {
			res[i] = prev.slice();
			continue;
		}

		newItem = newItem.concat(prev);
		newItem.sort();

		if (newItem.length > 3) {
			newItem = newItem.slice(0, 3);
		}

		res[i] = newItem;
		prev = newItem.slice();
	}

	return res;
}
