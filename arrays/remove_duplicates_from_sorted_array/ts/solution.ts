export const removeDuplicates = (nums: number[]): number => {
    if (nums.length === 0) {
        return 0
    }

    let pivot = 1

    for (let i = pivot; i < nums.length; i++) {
        if (nums[i] != nums[i-1]) {
            nums[pivot] = nums[i]
            pivot++
        }
    }

    return pivot
}