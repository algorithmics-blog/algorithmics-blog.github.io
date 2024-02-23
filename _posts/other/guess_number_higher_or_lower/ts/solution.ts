// Для реализации примера функция принимает в качестве параметра функцию guess.
// В задаче на leetcode это функция предоставляется как API.
export const guessNumber = (n: number, guess: (num: number) => number): number => {
    let left = 1
    let right = n

    while (left <= right) {
        const mid = Math.floor((left + right) / 2)
        const res = guess(mid)

        if (res === 0) {
            return mid
        } else if (res === 1) {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }

    return 0
}