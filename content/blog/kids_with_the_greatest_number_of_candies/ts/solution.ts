export const kidsWithCandies = (candies: number[], extraCandies: number): boolean[] => {
    let res: boolean[] = []

    const maxCandies = Math.max(...candies)

    candies.forEach(count => {
        res.push(count + extraCandies >= maxCandies)
    })

    return res
}
