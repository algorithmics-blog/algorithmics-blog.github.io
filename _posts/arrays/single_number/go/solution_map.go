package single_number

func singleNumberMap(nums []int) int {
	countMap := make(map[int]int)

	for _, num := range nums {
		countMap[num] += 1
	}

	for num, count := range countMap {
		if count == 1 {
			return num
		}
	}

	return 0
}
