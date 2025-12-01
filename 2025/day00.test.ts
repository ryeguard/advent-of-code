import { describe, it, expect } from "vitest";
import { part1, part2 } from './day00.logic'

describe("day 0 part 1", () => {
    it("just returns 1", () => {
        const exampleInput = [''];
        const result = part1(exampleInput);
        expect(result).toBe(1)
    })
})

describe("day 0 part 2", () => {
    it("just returns 2", () => {
        const exampleInput = ['']
        const result = part2(exampleInput)
        expect(result).toBe(2)
    })
})