const canBePlaced = (flowerbed: number[], i: number): boolean => {
    if (i == 0) {
        return flowerbed[i + 1] == 0
    }

    if (i == flowerbed.length - 1) {
        return flowerbed[i - 1] == 0
    }

    return flowerbed[i - 1] == 0 && flowerbed[i + 1] == 0
}

export const canPlaceFlowers = (flowerbed: number[], n: number): boolean => {
    if (n == 0) {
        return true
    }

    if (flowerbed.length == 0) {
        return false
    }

    if (flowerbed.length == 1) {
        if (n > 1) {
            return false
        }

        return flowerbed[0] == 0
    }

    for (let i = 0; i < flowerbed.length; i++) {
        if (flowerbed[i] == 1) {
            i++
            continue
        }

        if (canBePlaced(flowerbed, i)) {
            flowerbed[i] = 1
            n--
            i++
        }

        if (n == 0) {
            return true
        }
    }

    return n == 0
}


