package day07

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRecurse(t *testing.T) {
	var tests = []struct {
		name string
		eq   equation
		want bool
	}{
		{
			name: "2 numbers, addition",
			eq:   equation{numbers: []int{1, 2}, testValue: 3},
			want: true,
		},
		{
			name: "2 numbers, multiplication",
			eq:   equation{numbers: []int{1, 2}, testValue: 2},
			want: true,
		},
		{
			name: "3 numbers, addition",
			eq:   equation{numbers: []int{1, 2, 4}, testValue: 7},
			want: true,
		},
		{
			name: "3 numbers, multiplication",
			eq:   equation{numbers: []int{2, 3, 5}, testValue: 30},
			want: true,
		},
		{
			name: "3 numbers, addition then multiplication",
			eq:   equation{numbers: []int{2, 3, 5}, testValue: 25},
			want: true,
		},
		{
			name: "3 numbers, multiplication then addition",
			eq:   equation{numbers: []int{2, 3, 5}, testValue: 30},
			want: true,
		},
		{
			name: "3 numbers, multiplication then addition",
			eq:   equation{numbers: []int{2, 3, 5}, testValue: 30},
			want: true,
		},
		{
			name: "3 numbers, last number counts",
			eq:   equation{numbers: []int{2, 3, 5}, testValue: 6},
			want: false,
		},
		{
			name: "2 numbers, incorrect",
			eq:   equation{numbers: []int{1, 2, 3}, testValue: 4},
			want: false,
		},
	}

	for _, tc := range tests {
		got := recurse(tc.eq.testValue, tc.eq.numbers[0], tc.eq.numbers[1:], []func(int, int) int{Add, Multiply})
		require.Equal(t, tc.want, got, tc.name)
	}
}

func TestRecursePart2(t *testing.T) {
	var tests = []struct {
		name string
		eq   equation
		want bool
	}{
		{
			name: "2 numbers, concatenation",
			eq:   equation{numbers: []int{1, 2}, testValue: 12},
			want: true,
		},
		{
			name: "3 numbers, concatenation",
			eq:   equation{numbers: []int{1, 2, 3}, testValue: 123},
			want: true,
		},
	}

	for _, tc := range tests {
		got := recurse(tc.eq.testValue, tc.eq.numbers[0], tc.eq.numbers[1:], []func(int, int) int{Add, Multiply, Concat})
		require.Equal(t, tc.want, got, tc.name)
	}
}

func TestConcat(t *testing.T) {
	var tests = []struct {
		x, y int
		want int
	}{
		{x: 1, y: 2, want: 12},
		{x: 102, y: 20, want: 10220},
	}

	for _, tc := range tests {
		got := Concat(tc.x, tc.y)
		require.Equal(t, tc.want, got)
	}
}
