package main

import "fmt"

func day00(input []string) (int, int, error) {
	part1, err := day00Part1(input)
	if err != nil {
		return 0, 0, fmt.Errorf("part 1: %w", err)
	}

	part2, err := day00Part2(input)
	if err != nil {
		return 0, 0, fmt.Errorf("part 1: %w", err)
	}
	return part1, part2, nil
}

func day00Part1(input []string) (int, error) {
	return 1, nil
}

func day00Part2(input []string) (int, error) {
	return 2, nil
}
