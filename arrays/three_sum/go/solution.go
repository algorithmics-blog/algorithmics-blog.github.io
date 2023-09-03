package main

import "sort"

func threeSumTwoPointers(nums []int) [][]int {
	result := make([][]int, 0)

	sort.Ints(nums)

	for i, num := range nums {
		left := i + 1
		right := len(nums) - 1

		if i != 0 && nums[i-1] == num {
			continue
		}

		for left < right {
			sum := num + nums[left] + nums[right]

			if sum == 0 {
				result = append(result, []int{num, nums[left], nums[right]})
				right -= 1
				left += 1

				for left < right && nums[left] == nums[left-1] {
					left += 1
				}
			} else if sum > 0 {
				right -= 1
			} else {
				left = left + 1
			}
		}
	}

	return result
}
