import { readFileSync } from "fs";
import { part1, part2 } from "./day01.logic";

const input = readFileSync("input_data/day01.txt", "utf-8").trim().split("\n")

console.log(part1(input))
console.log(part2(input))