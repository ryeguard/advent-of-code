package day00

import "fmt"

func Solution(input []string) (int, int, error) {
	part1, err := part1(input)
	if err != nil {
		return 0, 0, fmt.Errorf("part 1: %w", err)
	}

	part2, err := part2(input)
	if err != nil {
		return 0, 0, fmt.Errorf("part 1: %w", err)
	}
	return part1, part2, nil
}

func part1(input []string) (int, error) {
	return len(input) + 1, nil
}

func part2(input []string) (int, error) {
	return len(input) + 2, nil
}
