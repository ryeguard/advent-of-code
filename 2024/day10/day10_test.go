package day10

import (
	"testing"

	"github.com/ryeguard/advent-of-code/goac"
	"github.com/stretchr/testify/require"
)

func TestFindPaths(t *testing.T) {
	grid := goac.Grid[int]{
		Data: [][]int{
			{0, 1, 2, 3},
			{1, 2, 3, 4},
			{8, 7, 6, 5},
			{9, 8, 7, 6},
		},
		Height: 4,
		Width:  4,
	}

	var tests = []struct {
		pos  goac.Position
		want []goac.Position
	}{
		{
			pos:  goac.Position{Row: 0, Col: 0},
			want: []goac.Position{{Row: 0, Col: 1}, {Row: 1, Col: 0}},
		},
		{
			pos:  goac.Position{Row: 1, Col: 0},
			want: []goac.Position{{Row: 1, Col: 1}},
		},
	}

	for _, tc := range tests {
		got := findPaths(grid, tc.pos)
		require.Equal(t, tc.want, got)
	}
}
