package day09

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNext(t *testing.T) {
	// 0..111....22222
	d := disk{diskMap: []int{1, 2, 3, 4, 5}}

	require.Equal(t, &block{pos: 0, id: ptr(0)}, d.next())
	require.Equal(t, &block{pos: 1}, d.next())
	require.Equal(t, &block{pos: 2}, d.next())
	require.Equal(t, &block{pos: 3, id: ptr(1)}, d.next())
	require.Equal(t, &block{pos: 4, id: ptr(1)}, d.next())
	require.Equal(t, &block{pos: 5, id: ptr(1)}, d.next())
	require.Equal(t, &block{pos: 6}, d.next())
	require.Equal(t, &block{pos: 7}, d.next())
	require.Equal(t, &block{pos: 8}, d.next())
	require.Equal(t, &block{pos: 9}, d.next())
	require.Equal(t, &block{pos: 10, id: ptr(2)}, d.next())
	require.Equal(t, &block{pos: 11, id: ptr(2)}, d.next())
	require.Equal(t, &block{pos: 12, id: ptr(2)}, d.next())
	require.Equal(t, &block{pos: 13, id: ptr(2)}, d.next())
	require.Equal(t, &block{pos: 14, id: ptr(2)}, d.next())

	var nilBlock *block
	require.Equal(t, nilBlock, d.next())
}

func TestLastFile(t *testing.T) {
	// 0..111....22222
	d := disk{diskMap: []int{1, 2, 3, 4, 5}, lastIdx: 4}

	require.Equal(t, 2, *d.lastFile())
	require.Equal(t, 2, *d.lastFile())
	require.Equal(t, 2, *d.lastFile())
	require.Equal(t, 2, *d.lastFile())
	require.Equal(t, 2, *d.lastFile())
	require.Equal(t, 1, *d.lastFile())
	require.Equal(t, 1, *d.lastFile())
	require.Equal(t, 1, *d.lastFile())
	require.Equal(t, 0, *d.lastFile())

	var nilInt *int
	require.Equal(t, nilInt, d.lastFile())
}

func TestSimple(t *testing.T) {
	// 0..111....22222
	// 022111222......
	// 0123456789
	d := disk{diskMap: []int{1, 2, 3, 4, 5}, lastIdx: 4}

	got, err := part1(d)
	require.NoError(t, err)
	require.Equal(t, (0*0 + 1*2 + 2*2 + 3*1 + 4*1 + 5*1 + 6*2 + 7*2 + 8*2), got)
}

func TestSimple2(t *testing.T) {
	// 54321
	// 00000....111..2
	// 000002111
	// 012345678
	d := disk{diskMap: []int{5, 4, 3, 2, 1}, lastIdx: 4}

	got, err := part1(d)
	require.NoError(t, err)
	require.Equal(t, (0*0 + 1*0 + 2*0 + 3*0 + 4*0 + 5*2 + 6*1 + 7*1 + 8*1), got)
}
