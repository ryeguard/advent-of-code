package main

//go:generate go run gen.go

import (
	"fmt"

	"github.com/ryeguard/advent-of-code/goac"
	_ "golang.org/x/tools/go/packages" // import here, used in go:build ignore'd gen code
)

func main() {
	daysToRun := []int{
		6,
	}

	for _, d := range daysToRun {
		input, err := goac.ReadInput(fmt.Sprintf("input_data/day%02d.txt", d))
		if err != nil {
			panic(err)
		}

		part1, part2, err := solutionFuncs[d](input)
		if err != nil {
			panic(fmt.Sprintf("error day %02d: %v", d, err))
		}
		fmt.Printf("Day %01d\n\tPart 1: %v\n\tPart 2: %v\n", d, part1, part2)
	}
}
