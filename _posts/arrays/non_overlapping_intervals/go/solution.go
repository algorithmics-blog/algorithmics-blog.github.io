package main

import "sort"

func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	res := 0
	k := intervals[0][1]

	for idx := 1; idx < len(intervals); idx++ {
		end := intervals[idx][1]
		start := intervals[idx][0]

		if start >= k {
			k = end
		} else {
			res++
		}
	}

	return res
}
