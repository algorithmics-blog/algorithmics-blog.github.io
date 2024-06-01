export const asteroidCollisionTwoIdx = (asteroids: number[]): number[] => {
    let i = 0;
    let j = 1;

    while (i >= 0 && j < asteroids.length) {
        const first = asteroids[i]
        const second = asteroids[j]

        if (first * second > 0 || first < 0) {
            i++
            j++
            continue
        }

        if (first > -1 * second) {
            asteroids = [...asteroids.slice(0, i + 1), ...asteroids.slice(j + 1)]
            continue
        }

        if (first < -1 * second) {
            asteroids = [...asteroids.slice(0, i), ...asteroids.slice(j)]
            j--
            i--
            if (i < 0) {
                i = 0
                j = 1
            }

            continue
        }

        if (first == -1 * second) {
            asteroids = [...asteroids.slice(0, i), ...asteroids.slice(j + 1)]
            j--
            i--
            if (i < 0) {
                i = 0
                j = 1
            }
        }
    }

    return asteroids
}
