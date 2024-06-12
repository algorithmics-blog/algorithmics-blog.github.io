package main

import (
	"strconv"
	"strings"
)

func equalPairs(grid [][]int) int {
	countMap := make(map[string]int)
	builder := strings.Builder{}
	counter := 0

	for _, row := range grid {
		for _, col := range row {
			builder.WriteString(strconv.Itoa(col))
			builder.WriteString(";")
		}

		countMap[builder.String()]++
		builder.Reset()
	}

	for j := 0; j < len(grid); j++ {
		for i := 0; i < len(grid); i++ {
			builder.WriteString(strconv.Itoa(grid[i][j]))
			builder.WriteString(";")
		}

		counter += countMap[builder.String()]
		builder.Reset()
	}

	return counter
}
