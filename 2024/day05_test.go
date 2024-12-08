package main

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDay5ByPage(t *testing.T) {

	rules := createSortingFunc(
		[][2]int{
			{47, 53},
			{97, 13},
			{97, 61},
			{97, 47},
			{75, 29},
			{61, 13},
			{75, 53},
			{29, 13},
			{97, 29},
			{53, 29},
			{61, 53},
			{97, 53},
			{61, 29},
			{47, 13},
			{75, 47},
			{97, 75},
			{47, 61},
			{75, 61},
			{47, 29},
			{75, 13},
			{53, 13},
		},
	)

	var tests = []struct {
		pages    []int
		isSorted bool
		sorted   []int
	}{
		{
			pages:    []int{75, 47, 61, 53, 29},
			isSorted: true,
			sorted:   []int{75, 47, 61, 53, 29},
		},
		{
			pages:    []int{97, 61, 53, 29, 13},
			isSorted: true,
			sorted:   []int{97, 61, 53, 29, 13},
		},
		{
			pages:    []int{75, 29, 13},
			isSorted: true,
			sorted:   []int{75, 29, 13},
		},
		{
			pages:    []int{75, 97, 47, 61, 53},
			isSorted: false,
			sorted:   []int{97, 75, 47, 61, 53},
		},
		{
			pages:    []int{61, 13, 29},
			isSorted: false,
			sorted:   []int{61, 29, 13},
		},
		{
			pages:    []int{97, 13, 75, 29, 47},
			isSorted: false,
			sorted:   []int{97, 75, 47, 29, 13},
		},
	}

	for _, tc := range tests {
		gotIsSorted := slices.IsSortedFunc(tc.pages, rules)
		require.Equal(t, tc.isSorted, gotIsSorted)

		slices.SortFunc(tc.pages, rules)
		require.Equal(t, tc.sorted, tc.pages)
	}
}
