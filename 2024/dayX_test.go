package main

import (
	"fmt"
	"testing"

	"github.com/ryeguard/advent-of-code/goac"
	"github.com/stretchr/testify/require"
)

type DayTest struct {
	day   int
	part1 PartTest
	part2 PartTest
}

type PartTest struct {
	answer int // 0 if answer is unknown

	incorrectGuess int
	// E.g., require.Greater if the incorrect guess was met with "Your answer was too low".
	hintFunc func(t require.TestingT, e1 interface{}, e2 interface{}, msgAndArgs ...interface{})
}

func TestDaysWithExample(t *testing.T) {
	var tests = []DayTest{
		{
			day:   0,
			part1: PartTest{answer: 1},
			part2: PartTest{
				hintFunc:       require.Greater,
				incorrectGuess: 1,
			},
		},
		{
			day:   2,
			part1: PartTest{answer: 2},
			part2: PartTest{answer: 4},
		},
		{
			day:   3,
			part1: PartTest{answer: 161},
			part2: PartTest{answer: 48},
		},
		{
			day:   4,
			part1: PartTest{answer: 18},
			part2: PartTest{answer: 9},
		},
		{
			day:   5,
			part1: PartTest{answer: 143},
			part2: PartTest{answer: 123},
		},
		{
			day:   6,
			part1: PartTest{answer: 41},
			part2: PartTest{answer: 6},
		},
		{
			day:   7,
			part1: PartTest{answer: 3749},
			part2: PartTest{answer: 11387},
		},
		{
			day:   8,
			part1: PartTest{answer: 14},
			part2: PartTest{answer: 34},
		},
	}

	for _, tc := range tests {
		msg := fmt.Sprintf("day %v [example data]", tc.day)

		input, err := goac.ReadInput(fmt.Sprintf("example_data/day%02d.txt", tc.day))
		require.NoError(t, err, msg)

		part1, part2, err := solutionFuncs[tc.day](input)
		require.NoError(t, err, msg)

		if tc.part1.answer != 0 {
			require.Equal(t, tc.part1.answer, part1, "%v: part 1", msg)
		} else if tc.part1.hintFunc != nil {
			tc.part1.hintFunc(t, part1, tc.part1.incorrectGuess, "%v: part 1", msg)
		}

		if tc.part2.answer != 0 {
			require.Equal(t, tc.part2.answer, part2, "%v: part 2", msg)
		} else if tc.part2.hintFunc != nil {
			tc.part2.hintFunc(t, part2, tc.part2.incorrectGuess, "%v: part 2", msg)
		}
	}
}

func TestDays(t *testing.T) {
	var tests = []DayTest{
		{
			day:   0,
			part1: PartTest{answer: 1},
			part2: PartTest{answer: 2},
		},
		{
			day:   1,
			part1: PartTest{answer: 3246517},
			part2: PartTest{answer: 29379307},
		},
		{
			day:   2,
			part1: PartTest{answer: 411},
			part2: PartTest{answer: 465},
		},
		{
			day:   3,
			part1: PartTest{answer: 179834255},
			part2: PartTest{answer: 80570939},
		},
		{
			day:   4,
			part1: PartTest{answer: 2462},
			part2: PartTest{answer: 1877},
		},
		{
			day:   5,
			part1: PartTest{answer: 6498},
			part2: PartTest{answer: 5017},
		},
		{
			day:   6,
			part1: PartTest{answer: 5242},
			part2: PartTest{answer: 1424},
		},
		{
			day:   7,
			part1: PartTest{answer: 1399219271639},
			part2: PartTest{answer: 275791737999003},
		},
	}

	for _, tc := range tests {
		msg := fmt.Sprintf("day %v [input data]", tc.day)

		input, err := goac.ReadInput(fmt.Sprintf("input_data/day%02d.txt", tc.day))
		require.NoError(t, err, msg)

		part1, part2, err := solutionFuncs[tc.day](input)
		require.NoError(t, err, msg)

		if tc.part1.answer != 0 {
			require.Equal(t, tc.part1.answer, part1, "%v: part 1", msg)
		} else if tc.part1.hintFunc != nil {
			tc.part1.hintFunc(t, part1, tc.part1.incorrectGuess, "%v: part 1", msg)
		}

		if tc.part2.answer != 0 {
			require.Equal(t, tc.part2.answer, part2, "%v: part 1", msg)
		} else if tc.part2.hintFunc != nil {
			tc.part2.hintFunc(t, part2, tc.part2.incorrectGuess, "%v: part 2", msg)
		}
	}
}

func ptr(i int) *int {
	return &i
}
