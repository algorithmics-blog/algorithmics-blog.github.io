package number_of_1_bits

func hammingWeight(n int) int {
	count := 0

	for n != 0 {
		count++
		n &= n - 1
	}

	return count
}
