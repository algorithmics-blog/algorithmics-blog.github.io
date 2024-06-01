const moveFirst = (arr: string[]) => {
	const first = arr.shift()

	if (first) {
		arr.push(first)
	}
}

export const predictPartyVictory = (senate: string): string => {
	let rCount = 0
	let dCount = 0

	const queue: string[] = []

	for (let i = 0; i < senate.length; i++) {
		if (senate[i] === 'R') {
			rCount++
		} else {
			dCount++
		}

		queue.push(senate[i])
	}

	while (rCount > 0 && dCount > 0) {
		let firstElem = queue[0]
		let killCount = 1

		moveFirst(queue)

		while (killCount > 0 && rCount > 0 && dCount > 0) {
			if (queue[0] == firstElem) {
				moveFirst(queue)
				killCount++
			} else {
				if (queue[0] == 'R') {
					rCount--
				} else {
					dCount--
				}
				killCount--
				queue.shift()
			}
		}
	}

	if (rCount > 0) {
		return "Radiant"
	}

	return "Dire"
}
