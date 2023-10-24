package main

import (
	"fmt"
	"math"
)

func main() {
	//fmt.Println(coinChange([]int{1, 2, 5}, 11))
	//fmt.Println(coinChangeDynamic([]int{186, 419, 83, 408}, 6249))
	fmt.Println(coinChange([]int{186, 419, 83, 408}, 6249))
}

func coinChange(coins []int, amount int) int {
	vector := make([]int, amount+1)

	for i := 0; i < amount+1; i++ {
		idxs := getPrevAmounts(coins, i)
		if len(idxs) == 0 {
			vector[i] = -1
		}

		min := math.MaxInt

		for _, idx := range idxs {
			if idx == 0 {
				min = 0
				break
			}

			if vector[idx] == -1 {
				continue
			}

			if vector[idx] < min {
				min = vector[idx]
			}
		}

		if min == math.MaxInt {
			vector[i] = -1
		} else {
			vector[i] = min + 1
		}
	}

	return vector[amount]
}

func getPrevAmounts(coins []int, target int) []int {
	res := make([]int, 0, len(coins))
	for _, coin := range coins {
		if coin > target {
			continue
		}

		res = append(res, target-coin)
	}

	return res
}

//func coinChange(coins []int, amount int) int {
//	if amount == 0 {
//		return 0
//	}
//
//	sort.Ints(coins)
//
//	return coinChangeRecursive(coins, amount, len(coins)-1)
//}
//
//func coinChangeRecursive(coins []int, target, currentIdx int) int {
//	for i := currentIdx; i >= 0; i-- {
//		if coins[i] > target {
//			continue
//		}
//
//		if target%coins[i] == 0 {
//			fmt.Println(fmt.Sprintf("target - %d, coin - %d, count - %d", target, coins[i], target/coins[i]))
//			return target / coins[i]
//		}
//
//		coinsAmount := coinChangeRecursive(coins, target-coins[i], i)
//		if coinsAmount != -1 {
//			fmt.Println(fmt.Sprintf("target - %d, coin - %d, count - %d", target, coins[i], 1))
//			return coinsAmount + 1
//		}
//	}
//
//	return -1
//}
