package main

import (
	"bufio"
	"fmt"
	"os"
)

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

	input, err = readInput("day02-input.txt")
	if err != nil {
		panic(err)
	}
	part1, part2, err = day02(input)
	if err != nil {
		panic(err)
	}
	fmt.Println(part1, part2)
}
