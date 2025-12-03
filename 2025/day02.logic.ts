export function part1(lines: string[]): number {
    var inputs = lines[0] ? lines[0]?.split(",") : [];

    var invalidSum = 0

    for (const input of inputs) {
        const bounds = input.split("-")
        const lower = bounds[0] ? +bounds[0] : 0
        const upper = bounds[1] ? +bounds[1] : 0
        for (let i = lower; i <= upper; i++) {
            const numStr = "" + i
            const numStrLen = numStr.length
            if (numStrLen % 2 == 0) {
                const left = numStr.substring(0, numStrLen / 2)
                const right = numStr.substring(numStrLen / 2, numStr.length)
                if (left == right) {
                    invalidSum += + i
                }
            }
        }
    }
    return invalidSum
}

export function part2(lines: string[]): number {
    var inputs = lines[0] ? lines[0]?.split(",") : [];

    var invalidSum = 0

    for (const input of inputs) {
        const bounds = input.split("-")
        const lower = bounds[0] ? +bounds[0] : 0
        const upper = bounds[1] ? +bounds[1] : 0
        outerLoop: for (let i = lower; i <= upper; i++) {
            const numStr = "" + i
            for (let j = 1; j <= numStr.length / 2; j++) {
                if (isRepeating(numStr, j)) {
                    invalidSum += i
                    continue outerLoop
                }
            }
        }
    }
    return invalidSum
}

export function isRepeating(input: string, digits: number): boolean {
    const arrays = input.split(input.substring(0, digits));
    return arrays.every(({ length }) => length === 0);
}   