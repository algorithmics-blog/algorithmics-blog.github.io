export function longestValidParenthesesStack(s: string): number {
    let stack: number[] = [-1];
    let res = 0;

    for (let i = 0; i < s.length; i++) {
        if (s[i] === '(') {
            stack.push(i);
            continue;
        }

        stack.pop();
        if (stack.length === 0) {
            stack.push(i);
            continue;
        }

        res = Math.max(res, i - stack[stack.length - 1]);
    }

    return res;
}
