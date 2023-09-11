package main

func main() {

}

func maxArea(height []int) int {
	res := 0
	i := 0
	j := len(height) - 1

	for i < j {
		length := j - i
		minHeight := height[i]
		if height[j] < minHeight {
			minHeight = height[j]
		}

		square := length * minHeight
		if square > res {
			res = square
		}

		if height[i] < height[j] {
			i++
			continue
		}

		if height[j] < height[i] {
			j--
			continue
		}

		i++
		j--
	}

	return res
}

func maxAreaBruteforce(height []int) int {
	res := 0
	for i := 0; i < len(height)-1; i++ {
		for j := len(height) - 1; j > i; j-- {
			length := j - i
			if length*height[i] < res {
				break
			}

			minHeight := height[i]
			if height[j] < minHeight {
				minHeight = height[j]
			}

			square := minHeight * length
			if square > res {
				res = square
			}
		}
	}

	return res
}
