export const reverse = (x: number): number => {
    const maxInt32 = Math.pow(2, 31) - 1;
    const minInt32 = maxInt32 * -1;

    let left = x
    let res = 0

    while (left !== 0) {
        const right = left % 10
        left = left > 0 ? Math.floor(left / 10) : Math.ceil(left / 10)

        const max = Math.floor(maxInt32 / 10)
        const min = Math.ceil(minInt32 / 10)

        if (res > max || res == max && right > 7) {
            return 0
        }
        if (res < min || res == min && right < -7) {
            return 0
        }

        res = res * 10 + right
    }

    return res
}
