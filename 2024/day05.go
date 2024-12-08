package main

import (
	"fmt"
	"strconv"
	"strings"
)

type rule struct {
	first, second int
}

func day05(input []string) (int, int, error) {
	var rules []rule
	var pages []map[int][]int
	var middlePages []int
	for _, r := range input {
		switch {
		case strings.Contains(r, "|"):
			var v1, v2 int
			if _, err := fmt.Sscanf(r, "%d|%d", &v1, &v2); err != nil {
				return 0, 0, fmt.Errorf("sscanf: %w", err)
			}
			rules = append(rules, rule{first: v1, second: v2})
		case strings.Contains(r, ","):
			pages = append(pages, map[int][]int{})
			split := strings.Split(r, ",")
			for i, s := range split {
				n, err := strconv.Atoi(s)
				if err != nil {
					return 0, 0, fmt.Errorf("atoi: %w", err)
				}

				if i == len(split)/2 {
					middlePages = append(middlePages, n)
				}

				if len(pages[len(pages)-1][n]) == 0 {
					pages[len(pages)-1][n] = []int{}
				}

				pages[len(pages)-1][n] = append(pages[len(pages)-1][n], i)
			}
		}
	}

	fmt.Println(middlePages)

	part1, err := day05Part1(rules, pages, middlePages)
	if err != nil {
		return 0, 0, fmt.Errorf("part 1: %w", err)
	}

	part2, err := day05Part2(input)
	if err != nil {
		return 0, 0, fmt.Errorf("part 1: %w", err)
	}
	return part1, part2, nil
}

func day05Part1(rules []rule, pages []map[int][]int, middlePages []int) (int, error) {
	var sum int
	for i, page := range pages {
		if checkPage(page, rules) {
			sum += middlePages[i]
		}
	}
	return sum, nil
}

func checkPageOrder(beforeIndexes, afterIndexes []int) bool {
	for _, bIdx := range beforeIndexes {
		for _, aIDx := range afterIndexes {
			if bIdx >= aIDx {
				return false
			}
		}
	}
	return true
}

func checkPage(page map[int][]int, rules []rule) bool {
	for _, rule := range rules {
		beforeIndexes, ok := page[rule.first]
		if !ok {
			continue
		}
		afterIndexes, ok := page[rule.second]
		if !ok {
			continue
		}

		if !checkPageOrder(beforeIndexes, afterIndexes) {
			return false
		}
	}
	return true
}

func day05Part2(input []string) (int, error) {
	return 2, nil
}
