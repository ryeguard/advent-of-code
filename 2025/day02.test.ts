import { describe, it, expect } from "vitest";
import { part1, part2, isRepeating } from './day02.logic'

describe("day 0 part 1", () => {
    it("just returns 1", () => {
        const exampleInput = ['11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124'];
        const result = part1(exampleInput);
        expect(result).toBe(1227775554)
    })
})

describe("day 0 part 2", () => {
    it("just returns 2", () => {
        const exampleInput = ['11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124'];
        const result = part2(exampleInput)
        expect(result).toBe(4174379265)
    })
})

describe("isRepeating", () => {
    it("11", () => {
        expect(isRepeating("11", 1)).toBe(true)
    })

    it("12", () => {
        expect(isRepeating("1212", 2)).toBe(true)
    })
    it("123", () => {
        expect(isRepeating("123123", 2)).toBe(false)
    })
    it("123", () => {
        expect(isRepeating("123123", 3)).toBe(true)
    })
})