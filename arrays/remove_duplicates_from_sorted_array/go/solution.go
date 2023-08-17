package main

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	pivot := 1

	for i := pivot; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[pivot] = nums[i]
			pivot++
		}
	}

	return pivot
}
