package goac

import (
	"bufio"
	"fmt"
	"os"
)

func ReadInput(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("open: %w", err)
	}
	defer file.Close()

	var ret []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ret = append(ret, scanner.Text())
	}
	return ret, nil
}

func ToGrid(input []string) [][]rune {
	ret := [][]rune{}
	for y, r := range input {
		ret = append(ret, []rune{})
		for _, c := range r {
			ret[y] = append(ret[y], c)
		}
	}
	return ret
}
