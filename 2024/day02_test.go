package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCheckReportSafeWithDampening(t *testing.T) {
	var tests = []struct {
		in   []int
		safe bool
	}{
		{
			in:   []int{1, 2, 3},
			safe: true,
		},
		{
			in:   []int{3, 2, 1},
			safe: true,
		},
		{
			in:   []int{9, 8, 7, 4, 7},
			safe: true,
		},
		{
			in:   []int{9, 5, 7, 4, 7},
			safe: false,
		},
		{
			in:   []int{19, 18, 15, 12},
			safe: true,
		},
		{
			in:   []int{19, 20, 18, 15, 12},
			safe: true,
		},
		{
			in:   []int{19, 20, 18, 14, 12},
			safe: false,
		},
		{
			in:   []int{14, 17, 20, 20, 23, 26, 29, 32},
			safe: true,
		},
		{
			in:   []int{14, 17, 20, 20, 23, 23, 29, 32},
			safe: false,
		},
		{
			in:   []int{14, 18, 21, 21, 23, 26, 29, 32},
			safe: false,
		},
		{
			in:   []int{28, 31, 33, 32, 33, 34},
			safe: true,
		},
	}

	for _, tc := range tests {
		got := checkReportSafeWithDampening(tc.in)
		require.Equal(t, tc.safe, got, tc.in)
	}
}

func TestPop(t *testing.T) {
	var tests = []struct {
		name     string
		in       []int
		index    int
		wantInt  int
		wantList []int
	}{
		{
			name:     "empty list",
			in:       nil,
			wantInt:  0,
			wantList: nil,
		},
		{
			name:     "empty list",
			in:       []int{},
			wantInt:  0,
			wantList: nil,
		},
		{
			name:     "single element list",
			in:       []int{1},
			index:    0,
			wantInt:  1,
			wantList: []int{},
		},
		{
			name:     "pop first",
			in:       []int{1, 2, 3},
			index:    0,
			wantInt:  1,
			wantList: []int{2, 3},
		},
		{
			name:     "pop second",
			in:       []int{1, 2, 3},
			index:    1,
			wantInt:  2,
			wantList: []int{1, 3},
		},
	}

	for _, tc := range tests {

		gotInt, gotList := pop(tc.index, tc.in)

		require.Equal(t, tc.wantInt, gotInt, tc.name)
		require.Len(t, gotList, len(tc.wantList), tc.name)
		require.Equal(t, tc.wantList, gotList, tc.name)
	}
}
