package day10

import (
	"fmt"

	"github.com/ryeguard/advent-of-code/goac"
)

var directions = []goac.Direction{
	goac.Right, goac.Down, goac.Left, goac.Up,
}

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

func parse(input []string) (goac.Grid[int], error) {
	return goac.IntGrid(input)
}

func part1(g goac.Grid[int]) (int, error) {

	var trailheads []goac.Position
	for r, row := range g.Data {
		for c, e := range row {
			if e != 0 {
				continue
			}

			trailheads = append(trailheads, goac.Position{Row: r, Col: c})
		}
	}

	var success int

	for _, th := range trailheads {
		trails := findPathsRec(g, th)
		found := map[goac.Position]bool{}
		for _, t := range trails {
			if g.Here(t) == 9 {
				if _, ok := found[t]; ok {
					continue
				}
				found[t] = true
				success++
			}
		}
	}

	return success, nil
}

func findPathsRec(grid goac.Grid[int], pos goac.Position) []goac.Position {
	var out []goac.Position
	for _, dir := range directions {
		curr := grid.Here(pos)
		if curr == 9 {
			return []goac.Position{pos}
		}

		nextValue, nextPos, ok := grid.Get(pos, dir)
		if !ok {
			continue
		}

		if nextValue-curr != 1 {
			continue
		}

		out = append(out, findPathsRec(grid, nextPos)...)
	}
	return out
}

func findPaths(grid goac.Grid[int], pos goac.Position) []goac.Position {
	var out []goac.Position
	for _, dir := range directions {
		curr := grid.Here(pos)
		if curr == 9 {
			return []goac.Position{pos}
		}

		nextValue, nextPos, ok := grid.Get(pos, dir)
		if !ok {
			continue
		}

		if nextValue-curr != 1 {
			continue
		}

		out = append(out, nextPos)
	}
	return out
}

func part2(_ any) (int, error) {
	return 2, nil
}
