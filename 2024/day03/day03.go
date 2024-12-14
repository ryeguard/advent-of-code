package day03

import (
	"errors"
	"fmt"
	"regexp"
)

func Solution(input []string) (int, int, error) {
	part1, err := day03Part2(input, `mul\([0-9]*,[0-9]*\)`)
	if err != nil {
		return 0, 0, fmt.Errorf("part 1: %w", err)
	}

	part2, err := day03Part2(input, `mul\([0-9]*,[0-9]*\)|(do\(\))|don't\(\)`)
	if err != nil {
		return 0, 0, fmt.Errorf("part 2, %w", err)
	}

	return part1, part2, nil
}

func day03Part2(input []string, pattern string) (int, error) {
	var sum int
	instructionEnabled := true

	for _, line := range input {
		r, err := regexp.Compile(pattern)
		if err != nil {
			return 0, fmt.Errorf("compile: %w", err)
		}
		matches := r.FindAllString(line, -1)
		if matches == nil {
			return 0, errors.New("no matches")
		}

		for _, m := range matches {
			switch {
			case m == "do()":
				instructionEnabled = true
			case m == "don't()":
				instructionEnabled = false
			case instructionEnabled:
				var v1, v2 int
				_, err := fmt.Sscanf(m, "mul(%v,%v)", &v1, &v2)
				if err != nil {
					return 0, fmt.Errorf("sscanf: %w", err)
				}
				sum += v1 * v2
			}
		}
	}
	return sum, nil
}
