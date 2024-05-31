package plus_one

func plusOne(digits []int) []int {
	res := make([]int, 0, len(digits)+1)
	temp := 0

	for i := len(digits) - 1; i >= 0; i-- {
		digit := digits[i]
		temp += digit

		if i == len(digits)-1 {
			temp += 1
		}

		if temp < 10 {
			res = append(res, temp)
			temp = 0
			continue
		}

		res = append(res, temp%10)
		temp = 1
	}

	if temp != 0 {
		res = append(res, temp)
	}

	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}

	return res
}
