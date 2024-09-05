export function myAtoi(s: string): number {
    const maxInt32 = 2**31 - 1;
    const minInt32 = 2**31 * -1;

    let res = 0;
    let sign = 1;
    let index = 0;

    while (index < s.length && s[index] === ' ') {
        index++;
    }

    if (index < s.length && s[index] === '-') {
        sign = -1;
        index++;
    } else if (index < s.length && s[index] === '+') {
        index++;
    }

    for (let i = index; i < s.length; i++) {
        const char = s[i];

        if (char < '0' || char > '9') {
            break;
        }

        res = res * 10 + (char.charCodeAt(0) - '0'.charCodeAt(0));
        if (sign * res > maxInt32) {
            return maxInt32;
        }
        if (sign * res < minInt32) {
            return minInt32;
        }
    }

    if (res === 0) {
        // Обрабатываем случай, когда res = 0 и sign = -1. В таком случае JS вернет в качестве ответа «-0»
        return res
    }

    return sign * res;
}