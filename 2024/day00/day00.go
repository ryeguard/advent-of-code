package day00

import "fmt"

func Solution(input []string) (int, int, error) {
	parsed, err := parse(input)
	if err != nil {
		return 0, 0, fmt.Errorf("parse: %w", err)
	}

	part1, err := part1(parsed)
	if err != nil {
		return 0, 0, fmt.Errorf("part 1: %w", err)
	}

	part2, err := part2(parsed)
	if err != nil {
		return 0, 0, fmt.Errorf("part 2: %w", err)
	}
	return part1, part2, nil
}

func parse(input []string) (any, error) {
	return len(input), nil
}

func part1(_ any) (int, error) {
	return 1, nil
}

func part2(_ any) (int, error) {
	return 2, nil
}
