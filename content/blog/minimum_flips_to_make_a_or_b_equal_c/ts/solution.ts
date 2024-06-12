export const minFlips = (a: number, b: number, c: number): number => {
    let res = 0;

    while (a > 0 || b > 0 || c > 0) {
        let bitA = a % 2;
        let bitB = b % 2;
        let bitC = c % 2;

        a = Math.floor(a / 2);
        b = Math.floor(b / 2);
        c = Math.floor(c / 2);

        if ((bitA | bitB) == bitC) {
            // Условие задачи выполняется, значит биты на этой позиции менять не нужно
            continue;
        }

        if (bitC == 0) {
            res += bitA + bitB;
        } else {
            res += 1;
        }
    }


    return res;
};
