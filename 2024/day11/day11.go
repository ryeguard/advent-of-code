package day11

import (
	"fmt"
	"strconv"
	"strings"
)

func Solution(input []string) (int, int, error) {
	parsed, err := parse(input)
	if err != nil {
		return 0, 0, fmt.Errorf("parse: %w", err)
	}

	part1, err := part1And2(parsed, 25)
	if err != nil {
		return 0, 0, fmt.Errorf("part 1: %w", err)
	}

	part2, err := part1And2(parsed, 75)
	if err != nil {
		return 0, 0, fmt.Errorf("part 2: %w", err)
	}
	return part1, part2, nil
}

func parse(input []string) (map[int]int, error) {
	out := map[int]int{}
	split := strings.Split(input[0], " ")
	for _, s := range split {
		n, err := strconv.Atoi(s)
		if err != nil {
			return nil, fmt.Errorf("atoi: %w", err)
		}
		out[n]++
	}
	return out, nil
}

func part1And2(stones map[int]int, blinks int) (int, error) {
	for i := 0; i < blinks; i++ {
		afterBlink := map[int]int{}
		for k, v := range stones {
			switch {
			case k == 0:
				afterBlink[1] += v
			case len(fmt.Sprint(k))%2 == 0:
				numLength := len(fmt.Sprint(k))

				left, err := strconv.Atoi(fmt.Sprint(k)[numLength/2:])
				if err != nil {
					return 0, fmt.Errorf("atoi: %w", err)
				}
				afterBlink[left] += v

				right, err := strconv.Atoi(fmt.Sprint(k)[:numLength/2])
				if err != nil {
					return 0, fmt.Errorf("atoi: %w", err)
				}
				afterBlink[right] += v
			default:
				afterBlink[k*2024] += v
			}
		}
		stones = afterBlink
	}
	
	numStones := 0
	for _, s := range stones {
		numStones += s
	}
	return numStones, nil
}
