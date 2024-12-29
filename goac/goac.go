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

type Grid[T Gridable] struct {
	Data   [][]T
	Width  int // The number of columns
	Height int // The number of rows
}

type Position struct {
	Row, Col int
}

func RuneGrid(input []string) Grid[rune] {
	if len(input) == 0 {
		return Grid[rune]{}
	}
	g := Grid[rune]{
		Height: len(input),
		Width:  len(input[0]),
	}
	for y, r := range input {
		g.Data = append(g.Data, []rune{})
		for _, c := range r {
			g.Data[y] = append(g.Data[y], c)
		}
	}
	return g
}

func IntGrid(input []string) (Grid[int], error) {
	if len(input) == 0 {
		return Grid[int]{}, nil
	}

	g := Grid[int]{
		Height: len(input),
		Width:  len(input[0]),
	}
	for y, r := range input {
		g.Data = append(g.Data, []int{})
		for _, c := range r {
			n, err := strconv.Atoi(string(c))
			if err != nil {
				return Grid[int]{}, fmt.Errorf("atoi: %w", err)
			}
			g.Data[y] = append(g.Data[y], n)
		}
	}
	return g, nil
}

func (g Grid[T]) Get(pos Position, dir Direction) (T, Position, bool) {
	var zero T

	switch dir {
	case Right:
		if pos.Col+1 >= g.Width {
			return zero, pos, false
		}
		pos.Col++
	case Down:
		if pos.Row+1 >= g.Height {
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

	return g.Data[pos.Row][pos.Col], pos, true
}

func (g Grid[T]) Here(pos Position) T {
	return g.Data[pos.Row][pos.Col]
}

func (g Grid[Gridable]) Copy() Grid[Gridable] {
	copy := Grid[Gridable]{
		Data:   [][]Gridable{},
		Width:  g.Width,
		Height: g.Height,
	}

	for r, row := range g.Data {
		copy.Data = append(copy.Data, []Gridable{})
		copy.Data[r] = append(copy.Data[r], row...)
	}
	return copy
}
