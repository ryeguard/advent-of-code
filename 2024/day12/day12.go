package day12

import (
	"fmt"
	"slices"

	"github.com/ryeguard/advent-of-code/goac"
)

// Iterate over grid
// Check neighbors, if neighbor is part of area, also add current plot to that area ID
// At the same time, calculate area/perimeter that current plot contributes with (area always 1, perimeter depending on neighbors)

func Solution(input []string) (int, int, error) {
	grid := goac.RuneGrid(input)

	part1 := part1(grid)

	part2, err := part2(grid)
	if err != nil {
		return 0, 0, fmt.Errorf("part 2: %w", err)
	}
	return part1, part2, nil
}

var directions = []goac.Direction{
	goac.Right, goac.Down, goac.Left, goac.Up,
}

type region struct {
	letter    rune
	id        int
	area      int
	perimeter int
}

func (r region) String() string {
	return fmt.Sprintf("%v: %v * %v = %v", string(r.letter), r.area, r.perimeter, r.area*r.perimeter)
}

func part1(grid goac.Grid[rune]) int {

	regions := findRegions(grid)
	slices.SortFunc(regions, func(a, b *region) int {
		return a.id - b.id
	})

	price := 0
	for _, r := range regions {
		fmt.Println(r)
		price += r.area * r.perimeter
	}

	return price
}

func findRegions(grid goac.Grid[rune]) []*region {
	regions := map[goac.Position]*region{}
	nbrRegions := 0

	for r, row := range grid {
		for c, elem := range row {
			currPos := goac.Position{Row: r, Col: c}

			nbrNeighbors := 0
			var neighborRegion *region
			for _, dir := range directions {

				val, neighbor, ok := grid.Get(currPos, dir)
				if !ok {
					continue
				}

				if val != elem {
					continue
				}

				nbrNeighbors++

				if neighborRegion != nil {
					regions[neighbor] = neighborRegion
				}

				// if neighbor is part of a region, assign it to current plot
				if reg, ok := regions[neighbor]; ok {
					if neighborRegion == nil {
						neighborRegion = reg
					}
					regions[currPos] = reg
				}

			}
			if _, ok := regions[currPos]; !ok {
				regions[currPos] = &region{id: nbrRegions, letter: elem}
				nbrRegions++
			}
			regions[currPos].area++
			regions[currPos].perimeter += 4 - nbrNeighbors
		}
	}

	var uniqueRegions []*region
	seen := map[int]bool{}
	for _, reg := range regions {
		if _, ok := seen[reg.id]; ok {
			continue
		}
		uniqueRegions = append(uniqueRegions, reg)
		seen[reg.id] = true
	}

	return uniqueRegions
}

func part2(_ any) (int, error) {
	return 2, nil
}
