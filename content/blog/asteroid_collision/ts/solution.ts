export function asteroidCollision(asteroids: number[]): number[] {
    let stack: number[] = []

    asteroids.forEach((asteroid) => {
        stack = appendAsteroidToStack(stack, asteroid)
    })

    return stack
}

function appendAsteroidToStack(stack: number[], asteroid: number): number[] {
    if (stack.length == 0) {
        stack.push(asteroid)
        return stack
    }

    const current = stack[stack.length - 1]

    if (current * asteroid > 0 || current < 0) {
        stack.push(asteroid)
        return stack
    }

    if (current > -1 * asteroid) {
        return stack
    }

    if (current < -1 * asteroid) {
        stack.pop()
        return appendAsteroidToStack(stack, asteroid)
    }

    stack.pop()

    return stack
}
