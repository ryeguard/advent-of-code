package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"slices"
)

func day01(input []string) (int, int, error) {
	var left, right []int
	for _, row := range input {
		var v1, v2 int
		_, err := fmt.Sscanf(row, "%d %d", &v1, &v2)
		if err != nil {
			return 0, 0, fmt.Errorf("scan: %w", err)
		}
		left = append(left, v1)
		right = append(right, v2)
	}

	part1, err := day01Part1(left, right)
	if err != nil {
		return 0, 0, fmt.Errorf("part 1: %w", err)
	}

	part2, err := day01Part2(left, right)
	if err != nil {
		return 0, 0, fmt.Errorf("part 2: %w", err)
	}

	return part1, part2, nil

}

func day01Part1(left, right []int) (int, error) {
	if len(left) != len(right) {
		return 0, errors.New("left and right lists have different length")
	}

	slices.Sort(left)
	slices.Sort(right)

	var sum int
	for i := 0; i < len(left); i++ {
		distance := int(math.Abs(float64(left[i] - right[i])))
		sum += distance
	}

	return sum, nil
}

func day01Part2(left, right []int) (int, error) {
	var lOccurances, rOccurances = map[int]int{}, map[int]int{}
	for _, l := range left {
		lOccurances[l]++
	}
	for _, r := range right {
		rOccurances[r]++
	}

	var similarityScore int
	for k, v := range lOccurances {
		similarityScore += k * v * rOccurances[k]
	}

	return similarityScore, nil
}

func readInput(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("open: %w", err)
	}
	defer file.Close()

	var ret []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ret = append(ret, scanner.Text())
	}
	return ret, nil
}

func main() {

	input, err := readInput("day01-input.txt")
	if err != nil {
		panic(err)
	}

	part1, part2, err := day01(input)
	if err != nil {
		panic(err)
	}

	fmt.Println(part1, part2)
}
