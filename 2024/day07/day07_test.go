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
		got := recurse(tc.eq.testValue, tc.eq.numbers[0], tc.eq.numbers[1:])
		require.Equal(t, tc.want, got, tc.name)
	}

}
