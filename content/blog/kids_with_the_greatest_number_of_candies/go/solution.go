package kids_with_the_greatest_number_of_candies

import "slices"

func kidsWithCandies(candies []int, extraCandies int) []bool {
	res := make([]bool, 0, len(candies))

	maxCandies := slices.Max(candies)

	for _, count := range candies {
		res = append(res, count+extraCandies >= maxCandies)
	}

	return res
}
