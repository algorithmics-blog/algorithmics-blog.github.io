export const compress = (chars: string[]): number => {
	let newPositionIdx = 0;

	for (let firstElemIdx = 0; firstElemIdx < chars.length; firstElemIdx++) {
		const lastChar = chars[firstElemIdx];

		let lastElemIdx = firstElemIdx + 1;
		while (lastElemIdx < chars.length && chars[lastElemIdx] === lastChar) {
			lastElemIdx++;
		}

		chars[newPositionIdx] = chars[lastElemIdx - 1];
		newPositionIdx++;

		const counter = lastElemIdx - firstElemIdx;
		if (counter > 1) {
			for (const countChar of counter.toString()) {
				chars[newPositionIdx] = countChar;
				newPositionIdx++;
			}
		}

		firstElemIdx = lastElemIdx - 1;
	}

	chars.splice(newPositionIdx, chars.length - newPositionIdx);
	return chars.length;
}
