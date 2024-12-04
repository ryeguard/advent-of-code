package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDay03(t *testing.T) {
	input, err := readInput("day03-example.txt")
	if err != nil {
		panic(err)
	}

	part1, part2, err := day03(input)
	if err != nil {
		t.Fatal(err)
	}

	require.Equal(t, 161, part1, "part 1, example input")
	require.Equal(t, 48, part2, "part 2, example input")

	input, err = readInput("day03-input.txt")
	if err != nil {
		panic(err)
	}

	part1, part2, err = day03(input)
	if err != nil {
		t.Fatal(err)
	}

	require.Equal(t, 179834255, part1)
	require.Equal(t, 80570939, part2)
}
