package day12

import (
	"slices"
	"testing"

	"github.com/ryeguard/advent-of-code/goac"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	grid := goac.Grid[rune]{
		{'A', 'A', 'A', 'A'},
		{'B', 'B', 'C', 'D'},
		{'B', 'B', 'C', 'C'},
		{'E', 'E', 'E', 'C'},
	}

	got := findRegions(grid)
	slices.SortFunc(got, func(a, b *region) int {
		return a.id - b.id
	})

	require.Equal(t, []*region{
		{letter: 'A', id: 0, area: 4, perimeter: 10},
		{letter: 'B', id: 1, area: 4, perimeter: 8},
		{letter: 'C', id: 2, area: 4, perimeter: 10},
		{letter: 'D', id: 3, area: 1, perimeter: 4},
		{letter: 'E', id: 4, area: 3, perimeter: 8},
	}, got)
}
