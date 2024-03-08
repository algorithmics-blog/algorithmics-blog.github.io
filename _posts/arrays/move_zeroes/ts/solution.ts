export const moveZeroes = (nums: number[]): void => {
	for (let firstZeroIndex = 0, i = 0; i < nums.length; i++) {
		if (nums[i] !== 0) {
			const temp = nums[i]
			nums[i] = nums[firstZeroIndex]
			nums[firstZeroIndex] = temp
			firstZeroIndex++
		}
	}
};
