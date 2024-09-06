package counting_bits

func hammingWeight(n int) int {
	count := 0

	for n != 0 {
		count++
		n &= n - 1
	}

	return count
}

func countBitsPopCount(n int) []int {
	res := make([]int, 0, n+1)

	for i := 0; i <= n; i++ {
		res = append(res, hammingWeight(i))
	}

	return res
}
