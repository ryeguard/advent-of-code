package main

import "testing"

func TestDay01(t *testing.T) {
	input, err := readInput("day01-input.txt")
	if err != nil {
		panic(err)
	}

	part1, part2, err := day01(input)
	if err != nil {
		t.Fatal(err)
	}
	if part1 != 3246517 {
		t.Fatal("day 1 part 1")
	}

	if part2 != 29379307 {
		t.Fatal("day 1 part 2")
	}
}
