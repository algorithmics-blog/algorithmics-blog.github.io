package maximum_average_subarray_i

func findMaxAverage(nums []int, k int) float64 {
	sum := 0

	for i := 0; i < k; i++ {
		sum += nums[i]
	}

	maxSum := sum

	for i := k; i < len(nums); i++ {
		sum = sum + nums[i] - nums[i-k]
		if sum > maxSum {
			maxSum = sum
		}
	}

	return float64(maxSum) / float64(k)
}
