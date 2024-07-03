export function longestValidParenthesesWindowing(s: string): number {
    return recursiveLongestValidParenthesesWindowing(s, s.length)
}

function recursiveLongestValidParenthesesWindowing(s: string, substrLen: number): number {
    if (substrLen === 0) {
        return 0;
    }

    let left = 0
    let right = substrLen

    while (right <= s.length) {
        if (isValidParenthesesByCounter(s.substring(left, right))) {
            return right - left
        }
        left++
        right++
    }

    return recursiveLongestValidParenthesesWindowing(s, substrLen - 1);
}

function isValidParenthesesByCounter(s: string): boolean {
    let counter = 0

    for (const char of s) {
        if (char === '(') {
            counter++
            continue
        }

        if (counter === 0) {
            return false
        }

        counter--
    }

    return counter === 0
}