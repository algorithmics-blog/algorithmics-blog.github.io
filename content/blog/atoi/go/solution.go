package atoi

func myAtoi(s string) int {
	const (
		maxInt32 = 1<<31 - 1
		minInt32 = -1 << 31
	)

	var (
		res   = 0
		sign  = 1
		index = 0
	)

	runes := []rune(s)

	for index < len(s) && runes[index] == ' ' {
		index++
	}

	if index < len(runes) && runes[index] == '-' {
		sign = -1
		index++
	} else if index < len(runes) && runes[index] == '+' {
		index++
	}

	for i := index; i < len(runes); i++ {
		char := runes[i]

		if char < '0' || char > '9' {
			break
		}

		res = res*10 + int(char-'0')
		if sign*res > maxInt32 {
			return maxInt32
		}
		if sign*res < minInt32 {
			return minInt32
		}
	}

	return sign * res
}
