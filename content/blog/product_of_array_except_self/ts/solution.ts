// Хак, так как произведение отрицательного числа на 0 в JS дает -0
const multiply = (a: number, b: number): number => {
    const res = a * b
    return res === -0 ? 0 : res
}

export const productExceptSelf = (nums: number[]): number[] => {
    let res: number[] = []
    res[0] = 1

    for (let i = 1; i < nums.length; i++) {
        res[i] = multiply(res[i - 1], nums[i - 1])
    }

    let r = 1

    for (let i = nums.length - 1; i >= 0; i--) {
        res[i] = multiply(r, res[i])
        r *= nums[i]
    }

    return res
}
