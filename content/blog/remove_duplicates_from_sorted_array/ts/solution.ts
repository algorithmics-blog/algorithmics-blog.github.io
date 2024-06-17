export const removeDuplicates = (nums: number[]): number => {
    let pivot = 1

    for (let i = 1; i < nums.length; i++) {
        if (nums[i] != nums[i - 1]) {
            nums[pivot] = nums[i]
            pivot++
        }
    }

    return pivot
}