import { describe, it, expect } from "vitest";
import { part1, part2, div } from './day01.logic'
import { readFileSync } from "fs";

describe("day 1 part 1", () => {
    it("", () => {
        const exampleInput = readFileSync("example_data/day01.txt", "utf-8").trim().split("\n");
        const result = part1(exampleInput);
        expect(result).toBe(3)
    })
})



describe("integer division", () => {
    it("works with positive ints", () => {
        const result = div(10, 10)
        expect(result).toBe(1)
    })

    it("works with positive ints", () => {
        const result = div(5, 10)
        expect(result).toBe(0)
    })

    it("works with positive ints", () => {
        const result = div(10, 4)
        expect(result).toBe(2)
    })

    it("works with negative ints", () => {
        const result = div(-10, 10)
        expect(result).toBe(1)
    })

    it("works with negative ints", () => {
        const result = div(-10, 4)
        expect(result).toBe(2)
    })
})

describe("day 1 part 2", () => {
    it("", () => {
        const exampleInput = ["L50", "R50"]
        const result = part2(exampleInput)
        expect(result).toBe(1)
    })
    it("", () => {
        const exampleInput = ["L50", "L50"]
        const result = part2(exampleInput)
        expect(result).toBe(1)
    })
    it("", () => {
        const exampleInput = ["R50", "L50"]
        const result = part2(exampleInput)
        expect(result).toBe(1)
    })
    it("", () => {
        const exampleInput = ["R50", "R50"]
        const result = part2(exampleInput)
        expect(result).toBe(1)
    })

    it("", () => {
        const exampleInput = ["L150", "L50"]
        const result = part2(exampleInput)
        expect(result).toBe(2)
    })
    it("", () => {
        const exampleInput = ["L150", "R50"]
        const result = part2(exampleInput)
        expect(result).toBe(2)
    })
    it("", () => {
        const exampleInput = ["R150", "L50"]
        const result = part2(exampleInput)
        expect(result).toBe(2)
    })
    it("", () => {
        const exampleInput = ["R150", "R50"]
        const result = part2(exampleInput)
        expect(result).toBe(2)
    })

    it("", () => {
        const exampleInput = readFileSync("example_data/day01.txt", "utf-8").trim().split("\n");
        const result = part2(exampleInput)
        expect(result).toBe(6)
    })
})