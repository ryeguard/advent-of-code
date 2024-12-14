package day02

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func part1(reports []string) (int, error) {
	var sumSafeReports int

	for _, report := range reports {

		split := strings.Split(report, " ")

		var levels []int
		for _, s := range split {
			l, err := strconv.Atoi(s)
			if err != nil {
				return 0, fmt.Errorf("atoi: %w", err)
			}
			levels = append(levels, l)
		}

		isSafe := checkReportSafe(levels)
		if isSafe {
			sumSafeReports++
		}

	}
	return sumSafeReports, nil
}

func checkReportSafe(levels []int) bool {
	shouldIncrease := levels[len(levels)-1] > levels[0]

	for i := 1; i < len(levels); i++ {
		diff := levels[i] - levels[i-1]
		absDiff := int(math.Abs(float64(diff)))

		increased := diff > 0
		okDiff := absDiff > 0 && absDiff <= 3

		if shouldIncrease != increased || !okDiff {
			return false
		}
	}
	return true
}

func checkReportSafeWithDampening(levels []int) bool {
	badLevels := map[int]bool{}

	shouldIncrease := levels[len(levels)-1] > levels[0]

	for i := 0; i < len(levels); i++ {
		if i != 0 {
			diff := levels[i] - levels[i-1]
			abs := int(math.Abs(float64(diff)))

			increased := diff > 0
			okDiff := abs > 0 && abs <= 3

			if shouldIncrease != increased || !okDiff {
				badLevels[i] = true
			}
		}
		if i < len(levels)-1 {
			diff := levels[i+1] - levels[i]
			abs := int(math.Abs(float64(diff)))

			increased := diff > 0
			okDiff := abs > 0 && abs <= 3

			if shouldIncrease != increased || !okDiff {
				badLevels[i] = true
			}
		}
	}

	if len(badLevels) == 0 {
		return true
	}

	for k := range badLevels {
		_, newLevels := pop(k, levels)
		if checkReportSafe(newLevels) {
			return true
		}
	}
	return false
}

func part2(reports []string) (int, error) {
	var sumSafeReports int

	for _, report := range reports {

		split := strings.Split(report, " ")

		var levels []int
		for _, s := range split {
			l, err := strconv.Atoi(s)
			if err != nil {
				return 0, fmt.Errorf("atoi: %w", err)
			}
			levels = append(levels, l)
		}

		isSafe := checkReportSafeWithDampening(levels)
		if isSafe {
			sumSafeReports++
		}
	}
	return sumSafeReports, nil
}

func Solution(reports []string) (int, int, error) {
	part1, err := part1(reports)
	if err != nil {
		return 0, 0, fmt.Errorf("part 1: %w", err)
	}

	part2, err := part2(reports)
	if err != nil {
		return 0, 0, fmt.Errorf("part 2: %w", err)
	}

	return part1, part2, nil
}

func pop(index int, list []int) (int, []int) {
	if len(list) == 0 {
		return 0, nil
	}
	listCopy := append([]int(nil), list...)
	v := listCopy[index]
	left := listCopy[:index]
	right := listCopy[index+1:]
	l := append(left, right...)
	return v, l
}
