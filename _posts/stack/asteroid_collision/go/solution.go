package main

func asteroidCollision(asteroids []int) []int {
	stack := make([]int, 0, len(asteroids))

	for i := 0; i < len(asteroids); i++ {
		stack = appendAsteroidToStack(stack, asteroids[i])
	}

	return stack
}

func appendAsteroidToStack(stack []int, asteroid int) []int {
	if len(stack) == 0 {
		return append(stack, asteroid)
	}

	current := stack[len(stack)-1]

	if current*asteroid > 0 || current < 0 {
		return append(stack, asteroid)
	}

	if current > -1*asteroid {
		return stack
	}

	if current < -1*asteroid {
		return appendAsteroidToStack(stack[:len(stack)-1], asteroid)
	}

	return stack[:len(stack)-1]
}
