package day08

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateAntinode(t *testing.T) {
	var tests = []struct {
		primary, secondary coordinate
		want               coordinate
	}{
		{
			primary:   coordinate{2, 0},
			secondary: coordinate{4, 0},
			want:      coordinate{6, 0},
		},
		{
			primary:   coordinate{0, 2},
			secondary: coordinate{0, 4},
			want:      coordinate{0, 6},
		},
		{
			primary:   coordinate{3, 3},
			secondary: coordinate{6, 6},
			want:      coordinate{9, 9},
		},
		{
			primary:   coordinate{3, 6},
			secondary: coordinate{6, 3},
			want:      coordinate{9, 0},
		},
	}

	for _, tc := range tests {
		require.Equal(t, tc.want, createAntinodes(tc.primary, tc.secondary))
	}
}
