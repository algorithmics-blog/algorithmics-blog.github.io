package main

func maxProfit(prices []int) int {
	res := 0
	min := prices[0]
	max := prices[0]

	for i := 1; i < len(prices); i++ {
		if prices[i] > max {
			max = prices[i]
		} else {
			res += max - min
			min = prices[i]
			max = prices[i]
		}
	}

	res += max - min

	return res
}

func maxProfit2(prices []int) int {
	res := 0

	for i := 0; i < len(prices)-1; i++ {
		if prices[i] < prices[i+1] {
			res += prices[i+1] - prices[i]
		}
	}

	return res
}
