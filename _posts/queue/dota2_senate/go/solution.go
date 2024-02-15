package main

func predictPartyVictory(senate string) string {
	rCount, dCount := 0, 0
	queue := make([]byte, 0, 2*len(senate))
	for i := 0; i < len(senate); i++ {
		if senate[i] == 'R' {
			rCount++
		} else {
			dCount++
		}

		queue = append(queue, senate[i])
	}

	for rCount > 0 && dCount > 0 {
		firstElem := queue[0]
		killCount := 1
		queue = append(queue, firstElem)
		queue = queue[1:]

		for killCount > 0 && rCount > 0 && dCount > 0 {
			if queue[0] == firstElem {
				queue = append(queue, queue[0])
				queue = queue[1:]
				killCount++
			} else {
				if queue[0] == 'R' {
					rCount--
				} else {
					dCount--
				}
				killCount--
				queue = queue[1:]
			}
		}
	}

	if rCount > 0 {
		return "Radiant"
	}

	return "Dire"
}
