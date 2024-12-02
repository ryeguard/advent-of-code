package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func day02(reports []string) (int, int, error) {

	var sumSafeReports int

reportLoop:
	for r, report := range reports {

		levels := strings.Split(report, " ")

		var previousLevel int
		var isIncreasing *bool = nil
		for i, l := range levels {
			v, err := strconv.Atoi(l)
			if err != nil {
				return 0, 0, fmt.Errorf("atoi: %w", err)
			}

			if i == 0 {
				previousLevel = v
				continue
			}

			if isIncreasing == nil && v > previousLevel {
				isIncreasing = ptrTo(true)
				if v-previousLevel > 3 {
					continue reportLoop
				}
				previousLevel = v
				continue
			} else if isIncreasing == nil && v < previousLevel {
				isIncreasing = ptrTo(false)
				if v-previousLevel < -3 {
					continue reportLoop
				}
				previousLevel = v
				continue
			}

			diff := int(math.Abs(float64(v - previousLevel)))

			if diff == 0 {
				continue reportLoop
			}

			increased := v > previousLevel

			if *isIncreasing != increased {
				continue reportLoop
			}

			if diff > 3 {
				continue reportLoop
			}
			previousLevel = v
		}
		sumSafeReports++
		fmt.Printf("report %v is safe (%v) %v\n", r, sumSafeReports, levels)
	}
	return sumSafeReports, 0, nil
}

func ptrTo(b bool) *bool {
	return &b
}
