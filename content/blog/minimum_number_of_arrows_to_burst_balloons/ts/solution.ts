export const findMinArrowShots = (points: number[][]): number => {
    // Сортировка точек по второй координате каждой пары
    points.sort((a, b) => a[1] - b[1]);

    let count = 0;
    let lastShot = -Infinity;

    for (let cords of points) {
        if (cords[0] > lastShot) {
            count++;
            lastShot = cords[1];
        }
    }

    return count;
}
