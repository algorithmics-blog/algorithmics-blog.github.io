export const findMaxAverage = (nums: number[], k: number): number => {
	let sum = 0

	for (let i = 0; i < k; i++) {
		sum += nums[i]
	}

	let maxSum = sum

	for (let i = k; i < nums.length; i++) {
		sum = sum + nums[i] - nums[i-k]
		if (sum > maxSum) {
			maxSum = sum
		}
	}

	return maxSum / k
}
