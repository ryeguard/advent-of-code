package day07

import (
	"fmt"
	"strconv"
	"strings"
)

type equation struct {
	testValue int
	numbers   []int
}

func Solution(input []string) (int, int, error) {
	var equations []equation
	for _, in := range input {
		split := strings.Split(in, ": ")
		tv, err := strconv.Atoi(split[0])
		if err != nil {
			return 0, 0, fmt.Errorf("atoi: %w", err)
		}

		eq := equation{testValue: tv}
		split = strings.Split(split[1], " ")
		for _, s := range split {
			n, err := strconv.Atoi(s)
			if err != nil {
				return 0, 0, fmt.Errorf("atoi: %w", err)
			}
			eq.numbers = append(eq.numbers, n)
		}
		equations = append(equations, eq)
	}

	part1, err := part1(equations)
	if err != nil {
		return 0, 0, fmt.Errorf("part 1: %w", err)
	}

	part2, err := part2(input)
	if err != nil {
		return 0, 0, fmt.Errorf("part 2: %w", err)
	}
	return part1, part2, nil
}

func part1(eqs []equation) (int, error) {
	var sum int
	for _, eq := range eqs {
		done := recurse(eq.testValue, eq.numbers[0], eq.numbers[1:])
		if done {
			sum += eq.testValue
		}
	}
	return sum, nil
}

func recurse(target int, cumSum int, numbers []int) bool {
	if len(numbers) < 1 {
		return false
	}

	sum := cumSum + numbers[0]
	product := cumSum * numbers[0]

	if product > target && sum > target {
		return false
	}
	if len(numbers) == 1 && (sum == target || product == target) {
		return true
	}

	return recurse(target, product, numbers[1:]) || recurse(target, sum, numbers[1:])
}

func part2(input []string) (int, error) {
	return len(input) + 2, nil
}
