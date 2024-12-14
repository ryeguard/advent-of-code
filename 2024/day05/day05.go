package day05

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func Solution(input []string) (int, int, error) {
	var rules [][2]int
	var pages [][]int

	for _, r := range input {
		switch {
		case strings.Contains(r, "|"):
			var v1, v2 int
			if _, err := fmt.Sscanf(r, "%d|%d", &v1, &v2); err != nil {
				return 0, 0, fmt.Errorf("sscanf: %w", err)
			}
			rules = append(rules, [2]int{v1, v2})
		case strings.Contains(r, ","):
			split := strings.Split(r, ",")
			pages = append(pages, []int{})
			for _, s := range split {
				n, err := strconv.Atoi(s)
				if err != nil {
					return 0, 0, fmt.Errorf("atoi: %w", err)
				}
				pages[len(pages)-1] = append(pages[len(pages)-1], n)
			}
		}
	}

	sortFunc := createSortingFunc(rules)

	var part1, part2 int
	for _, page := range pages {
		if slices.IsSortedFunc(page, sortFunc) {
			part1 += page[len(page)/2]
		} else {
			slices.SortFunc(page, sortFunc)
			part2 += page[len(page)/2]
		}
	}

	return part1, part2, nil
}

func createSortingFunc(rules [][2]int) func(a, b int) int {
	// "cmp(a, b) should return a negative number when a < b, a positive number when a > b and zero when a == b."
	// In this context, it means returning a negative number when the rule says that number should be first.
	return func(a, b int) int {
		for _, r := range rules {
			// order matches rule
			if r[0] == a && r[1] == b {
				return -1
			}
			// order does not match rule
			if r[0] == b && r[1] == a {
				return 1
			}
		}
		// no rule match, order does not matter
		return 0
	}
}
