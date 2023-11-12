export const maxProfit = (prices: number[]): number => {
    let res = 0
    let min = prices[0]
    let max = prices[0]

    for (let i = 1; i < prices.length; i++) {
        if (prices[i] > max) {
            max = prices[i]
        } else {
            res += max - min
            min = prices[i]
            max = prices[i]
        }
    }

    res += max - min

    return res
}

export const maxProfit2 = (prices: number[]): number => {
    let res = 0

    for (let i = 0; i < prices.length - 1; i++) {
        if (prices[i] < prices[i + 1]) {
            res += prices[i + 1] - prices[i]
        }
    }

    return res
}
