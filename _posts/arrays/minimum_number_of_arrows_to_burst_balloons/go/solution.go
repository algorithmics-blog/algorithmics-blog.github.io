package minimum_number_of_arrows_to_burst_balloons

import (
	"math"
	"sort"
)

func findMinArrowShots(points [][]int) int {
	// Сортировка точек по второй координате каждой пары
	sort.Slice(points, func(i int, j int) bool {
		return points[i][1] < points[j][1]
	})

	count := 0
	lastShot := math.MinInt

	for _, cords := range points {
		if cords[0] > lastShot {
			count++
			lastShot = cords[1]
		}
	}

	return count
}
