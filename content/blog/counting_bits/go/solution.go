package counting_bits

func countBits(n int) []int {
	response := make([]int, n+1)

	for i := 0; i <= n; i++ {
		response[i] = response[i>>1] + (i & 1)
	}

	return response
}
