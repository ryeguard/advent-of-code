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

func Multiply(x, y int) int {
	return x * y
}

func Add(x, y int) int {
	return x + y
}

func Concat(x, y int) int {
	concat, err := strconv.Atoi(fmt.Sprintf("%d%d", x, y))
	if err != nil {
		panic(err)
	}
	return concat
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

	part1, err := partX(equations, []func(int, int) int{Add, Multiply})
	if err != nil {
		return 0, 0, fmt.Errorf("part 1: %w", err)
	}

	part2, err := partX(equations, []func(int, int) int{Add, Multiply, Concat})
	if err != nil {
		return 0, 0, fmt.Errorf("part 2: %w", err)
	}
	return part1, part2, nil
}

func partX(eqs []equation, operators []func(int, int) int) (int, error) {
	var sum int
	for _, eq := range eqs {
		done := recurse(eq.testValue, eq.numbers[0], eq.numbers[1:], operators)
		if done {
			sum += eq.testValue
		}
	}
	return sum, nil
}

func recurse(target int, currentValue int, numbers []int, operators []func(int, int) int) bool {
	if len(numbers) < 1 {
		return false
	}

	var alternatives []int
	for _, op := range operators {
		alternatives = append(alternatives, op(currentValue, numbers[0]))
	}

	// If value surpasses the target value we can return,
	// since all operators only increase the value.
	if all(alternatives, func(n int) bool {
		return n > target
	}) {
		return false
	}

	if len(numbers) == 1 && any(alternatives, func(n int) bool {
		return n == target
	}) {
		return true
	}

	return any(alternatives, func(alt int) bool {
		return recurse(target, alt, numbers[1:], operators)
	})
}

func all[E comparable](list []E, predicate func(e E) bool) bool {
	for _, n := range list {
		if !predicate(n) {
			return false
		}
	}
	return true
}

func any[E comparable](list []E, predicate func(e E) bool) bool {
	for _, e := range list {
		if predicate(e) {
			return true
		}
	}
	return false
}
