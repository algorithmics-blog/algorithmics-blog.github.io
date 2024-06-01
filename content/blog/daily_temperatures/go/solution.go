package main

func dailyTemperatures(temperatures []int) []int {
	stack := make([]int, 0, len(temperatures))
	res := make([]int, len(temperatures))
	for j, dailyTemp := range temperatures {
		for len(stack) > 0 && dailyTemp > temperatures[stack[len(stack)-1]] {
			i := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[i] = j - i
		}

		stack = append(stack, j)
	}

	return res
}
