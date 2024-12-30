package day10

import (
	"testing"

	"github.com/ryeguard/advent-of-code/goac"
	"github.com/stretchr/testify/require"
)

func TestFindCompleteTrails(t *testing.T) {
	grid := goac.Grid[int]{
		{0, 1, 2, 3},
		{8, 8, 8, 4},
		{0, 0, 0, 5},
		{9, 8, 7, 6},
	}

	var tests = []struct {
		pos  goac.Position
		want []goac.Position
	}{
		{
			pos:  goac.Position{Row: 0, Col: 0},
			want: []goac.Position{{Row: 3, Col: 0}},
		},
	}

	for _, tc := range tests {
		got := findCompleteTrails(grid, tc.pos)
		require.Equal(t, tc.want, got)
	}
}
