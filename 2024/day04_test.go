package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHasNeighbor(t *testing.T) {
	//  0 1      8 9
	//   MMMSXXMASM
	//10 MSAMXMSMSA 19
	//   AMXSXMAAMM
	//   MSAMASMSMX
	//   XMASAMXAMM
	//   XXAMMXXAMA
	//   SMSMSASXSS
	//   SAXAMASAAA
	//80 MAMMMXMMMM 89
	//   MXMXAXMASX
	// 90 91    98 99

	b := board{
		width:  10,
		height: 10,
	}

	var tests = []struct {
		index int
		d     direction
		want  bool
	}{
		{0, left, false},
		{1, left, true},
		{8, left, true},
		{9, left, true},
		{10, left, false},
		{19, left, true},
		{80, left, false},
		{89, left, true},
		{90, left, false},
		{91, left, true},
		{98, left, true},
		{99, left, true},

		{0, right, true},
		{1, right, true},
		{8, right, true},
		{9, right, false},
		{10, right, true},
		{19, right, false},
		{80, right, true},
		{89, right, false},
		{90, right, true},
		{91, right, true},
		{98, right, true},
		{99, right, false},

		{0, top, false},
		{1, top, false},
		{8, top, false},
		{9, top, false},
		{10, top, true},
		{19, top, true},
		{80, top, true},
		{89, top, true},
		{90, top, true},
		{91, top, true},
		{98, top, true},
		{99, top, true},

		{0, bottom, true},
		{1, bottom, true},
		{8, bottom, true},
		{9, bottom, true},
		{10, bottom, true},
		{19, bottom, true},
		{80, bottom, true},
		{89, bottom, true},
		{90, bottom, false},
		{91, bottom, false},
		{98, bottom, false},
		{99, bottom, false},
	}

	for _, tc := range tests {
		got := b.hasNeighbor(tc.index, tc.d)
		require.Equal(t, tc.want, got)
	}
}

func TestHasNeighbor2(t *testing.T) {
	b := board{
		width:  4,
		height: 4,
	}

	var tests = []struct {
		index int
		d     direction
		want  bool
	}{
		{3, right, false},
	}

	for _, tc := range tests {
		got := b.hasNeighbor(tc.index, tc.d)
		require.Equal(t, tc.want, got)
	}
}

func TestGetNeighbor(t *testing.T) {
	//  0 12 3
	//   0123
	// 4 ABCD 7
	// 8 4567 11
	//   EFGH
	// 12131415

	b := board{
		width:   4,
		height:  4,
		letters: "0123ABCD4567EFGH",
	}

	var tests = []struct {
		index      int
		d          direction
		wantLetter rune
		wantIdx    int
	}{
		{0, left, '-', -1},
		{1, left, '0', 0},
		{2, left, '1', 1},
		{3, left, '2', 2},
		{4, left, '-', -1},
		{7, left, 'C', 6},
		{8, left, '-', -1},
		{11, left, '6', 10},
		{12, left, '-', -1},
		{13, left, 'E', 12},
		{14, left, 'F', 13},
		{15, left, 'G', 14},

		{0, right, '1', 1},
		{1, right, '2', 2},
		{2, right, '3', 3},
		{3, right, '-', -1},
		{4, right, 'B', 5},
		{7, right, '-', -1},
		{8, right, '5', 9},
		{11, right, '-', -1},
		{12, right, 'F', 13},
		{13, right, 'G', 14},
		{14, right, 'H', 15},
		{15, right, '-', -1},
	}

	for _, tc := range tests {
		if !b.hasNeighbor(tc.index, tc.d) {
			continue
		}
		gotLetter, gotIdx := b.getNeighbor(tc.index, tc.d)
		require.Equal(t, tc.wantIdx, gotIdx)
		require.Equal(t, tc.wantLetter, gotLetter, "(index %v, dir %v): want %v, got %v", tc.index, tc.d, string(tc.wantLetter), string(gotLetter))
	}
}
