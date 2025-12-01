package day11

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	tests := []struct {
		stones map[int]int
		blinks int
		want   int
	}{
		{
			stones: map[int]int{125: 1, 17: 1},
			blinks: 1,
			want:   3,
		},
		{
			stones: map[int]int{125: 1, 17: 1},
			blinks: 2,
			want:   4,
		},
		{
			stones: map[int]int{125: 1, 17: 1},
			blinks: 3,
			want:   5,
		},
		{
			stones: map[int]int{125: 1, 17: 1},
			blinks: 4,
			want:   9,
		},
		{
			stones: map[int]int{125: 1, 17: 1},
			blinks: 5,
			want:   13,
		},
		{
			stones: map[int]int{125: 1, 17: 1},
			blinks: 6,
			want:   22,
		},
	}

	for _, tc := range tests {
		got, err := part1And2(tc.stones, tc.blinks)
		require.NoError(t, err)
		require.Equal(t, tc.want, got)
	}
}
