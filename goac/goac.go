package goac

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Gridable interface {
	rune | int
}

type Direction int

const (
	Right Direction = iota
	Down
	Left
	Up
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

type Grid[T Gridable] [][]T

type Position struct {
	Row, Col int
}

func (g Grid[T]) Width() int {
	return len(g[0])
}

func (g Grid[T]) Height() int {
	return len(g)
}

func RuneGrid(input []string) Grid[rune] {
	if len(input) == 0 {
		return Grid[rune]{}
	}
	g := Grid[rune]{}
	for y, r := range input {
		g = append(g, []rune{})
		for _, c := range r {
			g[y] = append(g[y], c)
		}
	}
	return g
}

func IntGrid(input []string) (Grid[int], error) {
	if len(input) == 0 {
		return Grid[int]{}, nil
	}

	g := Grid[int]{}
	for y, r := range input {
		g = append(g, []int{})
		for _, c := range r {
			n, err := strconv.Atoi(string(c))
			if err != nil {
				return Grid[int]{}, fmt.Errorf("atoi: %w", err)
			}
			g[y] = append(g[y], n)
		}
	}
	return g, nil
}

func (g Grid[T]) Get(pos Position, dir Direction) (T, Position, bool) {
	var zero T

	switch dir {
	case Right:
		if pos.Col+1 >= g.Width() {
			return zero, pos, false
		}
		pos.Col++
	case Down:
		if pos.Row+1 >= g.Height() {
			return zero, pos, false
		}
		pos.Row++
	case Left:
		if pos.Col-1 < 0 {
			return zero, pos, false
		}
		pos.Col--
	case Up:
		if pos.Row-1 < 0 {
			return zero, pos, false
		}
		pos.Row--
	}

	return g[pos.Row][pos.Col], pos, true
}

func (g Grid[T]) Here(pos Position) T {
	return g[pos.Row][pos.Col]
}

func (g Grid[Gridable]) Copy() Grid[Gridable] {
	copy := Grid[Gridable]{}

	for r, row := range g {
		copy = append(copy, []Gridable{})
		copy[r] = append(copy[r], row...)
	}
	return copy
}
