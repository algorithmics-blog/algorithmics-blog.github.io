package main

func main() {

}

func myPow(x float64, n int) float64 {
	if x == 1 {
		return 1
	}

	if x == -1 {
		if n%2 == 0 {
			return 1
		}

		return -1
	}

	if n < 0 {
		return 1 / powAbsN(x, -1*n)
	}

	return powAbsN(x, n)
}

func powAbsN(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	if n == 1 {
		return x
	}

	if n%2 != 0 {
		return powAbsN(x*x, n/2) * x
	}

	return powAbsN(x*x, n/2)
}

func myPowBruteforce(x float64, n int) float64 {
	if x == 1 {
		return 1
	}

	if x == -1 {
		if n%2 == 0 {
			return 1
		}

		return -1
	}

	if n < 0 {
		return 1 / powAbsNBruteforce(x, -1*n)
	}

	return powAbsNBruteforce(x, n)
}

func powAbsNBruteforce(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	if n == 1 {
		return x
	}

	res := x

	for ; n > 1; n-- {
		res *= x
	}

	return res
}
