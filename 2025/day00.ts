import { readFileSync } from "fs";
import { part1, part2 } from "./day00.logic";

const input = readFileSync("example_data/day00.txt", "utf-8").trim().split("\n")

console.log("part1:", part1(input))
console.log("part2:", part2(input))