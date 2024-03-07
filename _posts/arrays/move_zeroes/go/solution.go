package main

func moveZeroes(nums []int) {
	for firstZeroIndex, i := 0, 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[firstZeroIndex] = nums[firstZeroIndex], nums[i]
			firstZeroIndex++
		}
	}
}
