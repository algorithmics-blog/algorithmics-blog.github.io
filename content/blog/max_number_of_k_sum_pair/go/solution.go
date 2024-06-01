package max_number_of_k_sum_pair

func maxOperations(nums []int, k int) int {
	digits := make(map[int]int)
	res := 0

	for _, num := range nums {
		diff := k - num

		if val := digits[diff]; val > 0 {
			digits[diff]--
			res++
		} else {
			digits[num]++
		}
	}

	return res
}
