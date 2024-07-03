package main

func removeDuplicates(nums []int) int {
	pivot := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[pivot] = nums[i]
			pivot++
		}
	}

	return pivot
}
