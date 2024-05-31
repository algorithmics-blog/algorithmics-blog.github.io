package main

func canPlaceFlowers(flowerbed []int, n int) bool {
	if n == 0 {
		return true
	}

	if len(flowerbed) == 0 {
		return false
	}

	if len(flowerbed) == 1 {
		if n > 1 {
			return false
		}

		return flowerbed[0] == 0
	}

	for i := 0; i < len(flowerbed); i++ {
		if flowerbed[i] == 1 {
			i++
			continue
		}

		if canBePlaced(flowerbed, i) {
			flowerbed[i] = 1
			n--
			i++
		}

		if n == 0 {
			return true
		}
	}

	return n == 0
}

func canBePlaced(flowerbed []int, i int) bool {
	if i == 0 {
		return flowerbed[i+1] == 0
	}

	if i == len(flowerbed)-1 {
		return flowerbed[i-1] == 0
	}

	return flowerbed[i-1] == 0 && flowerbed[i+1] == 0
}
