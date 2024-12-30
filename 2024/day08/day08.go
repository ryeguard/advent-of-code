package day08

import (
	"github.com/ryeguard/advent-of-code/goac"
)

type coordinate struct {
	x, y int
}

func Solution(input []string) (int, int, error) {
	grid := goac.RuneGrid(input)
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

	part1 := part1(antennas, grid.Width(), grid.Height())
	part2 := part2(antennas, grid.Width(), grid.Height())

	return countUnique(part1, grid.Width(), grid.Height()), countUnique(part2, grid.Width(), grid.Height()), nil
}

func countUnique(antinodes []coordinate, width, height int) int {
	uniqueAntinodes := map[coordinate]bool{}
	for _, an := range antinodes {
		// bound check
		if an.x < 0 || an.y < 0 || an.x >= width || an.y >= height {
			continue
		}

		if _, ok := uniqueAntinodes[an]; ok {
			continue
		} else {
			uniqueAntinodes[an] = true
		}
	}
	return len(uniqueAntinodes)
}

func part1(antennas map[rune][]coordinate, width, height int) []coordinate {
	var antinodes []coordinate
	for _, coordinates := range antennas {
		for _, coord1 := range coordinates {
			for _, coord2 := range coordinates {
				if coord1 == coord2 {
					continue
				}
				antinodes = append(antinodes, createAntiNode(coord1, coord2, width, height)...)
			}
		}
	}
	return antinodes
}

func part2(antennas map[rune][]coordinate, width, height int) []coordinate {
	var antinodes []coordinate
	for _, coordinates := range antennas {
		for _, coord1 := range coordinates {
			for _, coord2 := range coordinates {
				if coord1 == coord2 {
					continue
				}
				antinodes = append(antinodes, createAntiNodes(coord1, coord2, width, height)...)
			}
		}
	}

	return antinodes
}

// createAntiNode creates anti nodes for part 1
func createAntiNode(primary, secondary coordinate, width, height int) []coordinate {
	x := primary.x - 2*(primary.x-secondary.x)
	y := primary.y - 2*(primary.y-secondary.y)
	if x >= 0 && y >= 0 && x < width && y < height {
		return []coordinate{{
			x: x,
			y: y,
		}}
	}
	return nil
}

// createAntiNodes creates anti nodes for part 2
func createAntiNodes(primary, secondary coordinate, width, height int) []coordinate {
	var x, y int = primary.x, primary.y
	var coords []coordinate
	iter := 0
	for x >= 0 && y >= 0 && x < width && y < height {
		x = primary.x - iter*(primary.x-secondary.x)
		y = primary.y - iter*(primary.y-secondary.y)
		coords = append(coords, coordinate{
			x: x,
			y: y,
		})
		iter++
	}
	return coords
}
