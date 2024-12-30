package goac

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGridCopy(t *testing.T) {
	g := Grid[rune]{
		{'a', 'b', 'c'},
		{'d', 'e', 'f'},
	}

	copy := g.Copy()

	copy[0][0] = 'x'
	copy = append(copy, []rune{'g', 'h', 'i'})

	require.Equal(t, 'a', g[0][0])
	require.Equal(t, 'x', copy[0][0])

	require.Equal(t, 2, g.Height())
	require.Equal(t, 3, copy.Height())
}
