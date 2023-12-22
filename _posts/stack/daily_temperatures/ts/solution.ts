export const dailyTemperatures = (temperatures: number[]): number[] => {
	const stack: number[] = []
	const res: number[] = new Array(temperatures.length).fill(0)
	
	temperatures.forEach((dailyTemp, j) => {
		while (stack.length > 0 && dailyTemp > temperatures[stack[stack.length - 1]]) {
			const i = stack[stack.length-1]
			stack.pop()
			res[i] = j - i
		}

		stack.push(j)
	})

	return res
}
