package goac

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGridCopy(t *testing.T) {
	g := Grid[rune]{
		Data: [][]rune{
			{'a', 'b', 'c'},
			{'d', 'e', 'f'},
		},
		Width:  3,
		Height: 2,
	}

	copy := g.Copy()

	copy.Data[0][0] = 'x'
	copy.Data = append(copy.Data, []rune{'g', 'h', 'i'})
	copy.Height = 3

	require.Equal(t, 'a', g.Data[0][0])
	require.Equal(t, 'x', copy.Data[0][0])

	require.Equal(t, 2, g.Height)
	require.Equal(t, 3, copy.Height)
}
