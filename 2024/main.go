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
	var leftLocations, rightLocations []int
	for _, row := range input {
		var v1, v2 int
		_, err := fmt.Sscanf(row, "%d %d", &v1, &v2)
		if err != nil {
			return 0, 0, fmt.Errorf("scan: %w", err)
		}
		leftLocations = append(leftLocations, v1)
		rightLocations = append(rightLocations, v2)
	}

	totalDistance, err := day01Part1(leftLocations, rightLocations)
	if err != nil {
		return 0, 0, fmt.Errorf("part 1: %w", err)
	}

	similarityScore := day01Part2(leftLocations, rightLocations)

	return totalDistance, similarityScore, nil

}

func day01Part1(leftLocations, rightLocations []int) (totalDistance int, err error) {
	if len(leftLocations) != len(rightLocations) {
		return 0, errors.New("left and right lists have different length")
	}

	slices.Sort(leftLocations)
	slices.Sort(rightLocations)

	for i := 0; i < len(leftLocations); i++ {
		distance := int(math.Abs(float64(leftLocations[i] - rightLocations[i])))
		totalDistance += distance
	}
	return
}

func day01Part2(left, right []int) (similarityScore int) {
	var leftOccurrences, rightOccurrences = map[int]int{}, map[int]int{}
	for _, l := range left {
		leftOccurrences[l]++
	}
	for _, r := range right {
		rightOccurrences[r]++
	}

	for k, v := range leftOccurrences {
		similarityScore += k * v * rightOccurrences[k]
	}
	return
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
