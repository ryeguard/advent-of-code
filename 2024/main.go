package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ryeguard/advent-of-code/2024/day00"
	"github.com/ryeguard/advent-of-code/2024/day06"
)

var dayFuncs = [](func([]string) (int, int, error)){
	day00.Solution, // template
	day01, day02, day03, day04, day05,
	day06.Solution,
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
	daysToRun := []int{
		6,
	}

	for _, d := range daysToRun {
		input, err := readInput(fmt.Sprintf("input_data/day%02d.txt", d))
		if err != nil {
			panic(err)
		}

		part1, part2, err := dayFuncs[d](input)
		if err != nil {
			panic(fmt.Sprintf("error day %02d: %v", d, err))
		}
		fmt.Printf("Day %01d\n\tPart 1: %v\n\tPart 2: %v\n", d, part1, part2)
	}
}
