export const singleNumber = (nums: number[]): number => {
    let res = 0

    nums.forEach((num) => {
        res ^= num
    })

    return res
}