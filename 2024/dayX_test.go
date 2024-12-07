package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

type DayTest struct {
	id    int
	part1 PartTest
	part2 PartTest
}

type PartTest struct {
	answer *int // Nil if answer is unknown

	incorrectGuess int
	// E.g., require.Greater if the incorrect guess was met with "Your answer was too low".
	hintFunc func(t require.TestingT, e1 interface{}, e2 interface{}, msgAndArgs ...interface{})
}

func TestDaysWithExample(t *testing.T) {
	var tests = []DayTest{
		{
			id: 0,
			part1: PartTest{
				answer: ptr(1),
			},
			part2: PartTest{
				hintFunc:       require.Greater,
				incorrectGuess: 1,
			},
		},
	}

	for _, tc := range tests {
		msg := fmt.Sprintf("day %v [example data]", tc.id)

		input, err := readInput(fmt.Sprintf("day%02d-example.txt", tc.id))
		require.NoError(t, err, msg)

		part1, part2, err := dayFuncs[tc.id](input)
		require.NoError(t, err, msg)

		if tc.part1.answer != nil {
			require.Equal(t, *tc.part1.answer, part1, "%v: part 1", msg)
		} else if tc.part1.hintFunc != nil {
			tc.part1.hintFunc(t, part1, tc.part1.incorrectGuess, "%v: part 1", msg)
		}

		if tc.part2.answer != nil {
			require.Equal(t, *tc.part2.answer, part2, "%v: part 1", msg)
		} else if tc.part2.hintFunc != nil {
			tc.part2.hintFunc(t, part2, tc.part2.incorrectGuess, "%v: part 2", msg)
		}
	}
}

func ptr(i int) *int {
	return &i
}
