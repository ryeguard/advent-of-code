package day08

import (
	"fmt"

	"github.com/ryeguard/advent-of-code/goac"
)

type coordinate struct {
	x, y int
}

func Solution(input []string) (int, int, error) {

	grid := goac.ToGrid(input)
	antennas := map[rune][]coordinate{}
	for y, row := range grid {
		for x, r := range row {
			if r == '.' {
				continue
			}
			if _, ok := antennas[r]; !ok {
				antennas[r] = []coordinate{}
			}
			antennas[r] = append(antennas[r], coordinate{x: x, y: y})
		}
	}

	part1, err := part1(grid, antennas)
	if err != nil {
		return 0, 0, fmt.Errorf("part 1: %w", err)
	}

	part2, err := part2(input)
	if err != nil {
		return 0, 0, fmt.Errorf("part 2: %w", err)
	}
	return part1, part2, nil
}

func part1(grid [][]rune, antennas map[rune][]coordinate) (int, error) {
	var antinodes []coordinate
	for frequency, coordinates := range antennas {
		fmt.Println("f:", string(frequency))
		for _, coord1 := range coordinates {
			for _, coord2 := range coordinates {
				if coord1 == coord2 {
					continue
				}
				fmt.Println("hey")
				antinodes = append(antinodes, createAntinodes(coord1, coord2))
			}
		}
	}

	printGrid(grid)

	fmt.Println(antennas)

	uniqueAntinodes := map[coordinate]bool{}
	for _, an := range antinodes {
		// bound check
		if an.x < 0 || an.y < 0 || an.x >= len(grid[0]) || an.y >= len(grid) {
			continue
		}

		if _, ok := uniqueAntinodes[an]; ok {
			continue
		} else {
			uniqueAntinodes[an] = true
			grid[an.y][an.x] = '#'
		}
	}
	fmt.Println()
	printGrid(grid)
	return len(uniqueAntinodes), nil
}

func createAntinodes(primary, secondary coordinate) coordinate {
	return coordinate{
		x: primary.x - 2*(primary.x-secondary.x),
		y: primary.y - 2*(primary.y-secondary.y),
	}
}

func printGrid(grid [][]rune) {
	for _, row := range grid {
		for _, r := range row {
			fmt.Printf("%v", string(r))
		}
		fmt.Println()
	}
}

func part2(input []string) (int, error) {
	return len(input) + 2, nil
}
