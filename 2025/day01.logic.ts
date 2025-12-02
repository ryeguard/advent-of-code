export function part1(lines: string[]): number {
    const startValue = 50;
    var value = startValue
    var zeroed = 0;
    for (const line of lines) {
        switch (line.at(0)) {
            case "L":
                value = value - +line.substring(1, line.length);
                break;
            case "R":
                value = value + +line.substring(1, line.length);
                break;
        }

        value = mod(value, 100)
        if (!value) {
            zeroed += 1;
        }
    }
    return zeroed
}

export function part2(lines: string[]): number {
    const startValue = 50;
    var value = startValue
    var zeroed = 0;
    for (const line of lines) {
        var diff = 0
        switch (line.at(0)) {
            case "L":
                diff = -+line.substring(1, line.length);
                break;
            case "R":
                diff = +line.substring(1, line.length);
                break;
        }

        const before = value
        const after = value + diff

        value = mod(value+diff, 100)

        var passed = 0
        if (after < 0) {
            if (before == 0) {
                passed -= 1
            }
            passed += 1 + div(after+1, 100)
        } else if (after > 0) {
            passed += div(after-1, 100)
        }

        if (value == 0) {
            zeroed += 1;
        }
        zeroed += passed
    }
    return zeroed
}

function mod(n: number, m: number) {
    return ((n % m) + m) % m;
}

export function div(n: number, m: number) {
    return Math.floor(Math.abs(n) / m)
}
