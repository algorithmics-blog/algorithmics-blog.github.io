package number_of_1_bits

func hammingWeightPopCount(n int) int {
	mask := 1
	count := 0

	for i := 0; i < 32; i++ {
		if n&mask != 0 {
			count++
		}
		mask <<= 1
	}

	return count
}
