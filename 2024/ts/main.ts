function day01(input: string): { part1: number, part2: number } {
  const split = input.split("\n")
  console.log(split)
  const leftRight = split.map((val: string) => ({ left: val.split("  ")[0], right: val.split("  ")[1] }))
  const left = leftRight.map((val, _) => (val.left)).sort()
  console.log(left)
  return { part1: input.length, part2: 1 }
}

if (import.meta.main) {
  const text = await Deno.readTextFile("./../example_data/day01.txt");
  const { part1, part2 } = day01(text)
  console.log(`part1: ${part1}, part2: ${part2}`)
}
