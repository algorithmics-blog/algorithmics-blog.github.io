package guess_number_higher_or_lower

var guess func(num int) int

func guessNumber(n int) int {
	left := 1
	right := n

	for left <= right {
		mid := (left + right) / 2
		res := guess(mid)

		if res == 0 {
			return mid
		} else if res == 1 {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return 0
}
