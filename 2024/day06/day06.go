package day06

import (
	"fmt"

	"github.com/ryeguard/advent-of-code/goac"
)

type direction int

const (
	up direction = iota
	right
	down
	left
)

type guard struct {
	x, y int
	dir  direction
}

type board struct {
	grid          [][]rune
	guard         guard
	seenObstacles map[guard]bool
}

func Solution(input []string) (int, int, error) {
	b := board{grid: goac.ToGrid(input), seenObstacles: map[guard]bool{}}

	for y, in := range input {
		for x, c := range in {
			if c == '^' {
				b.guard = guard{
					x:   x,
					y:   y,
					dir: up,
				}
			}
		}
	}

	if b.guard.x == 0 && b.guard.y == 0 {
		return 0, 0, fmt.Errorf("start position not found")
	}

	part1, err := part1(b)
	if err != nil {
		return 0, 0, fmt.Errorf("part 1: %w", err)
	}

	part2, err := part2(b)
	if err != nil {
		return 0, 0, fmt.Errorf("part 2: %w", err)
	}
	return part1, part2, nil
}

func (b *board) move() (done, loop bool) {
	// mark current as visited
	b.grid[b.guard.y][b.guard.x] = 'X'

	nextX, nextY := b.guard.next()

	// bounds check
	if nextX < 0 || nextX >= len(b.grid) || nextY < 0 || nextY >= len(b.grid[b.guard.y]) {
		return true, false
	}

	if b.grid[nextY][nextX] == '#' {
		// Part 2: Check/save seen obstacles
		if _, ok := b.seenObstacles[b.guard]; ok {
			return true, true
		} else {
			b.seenObstacles[b.guard] = true
		}

		b.guard.turn()
		return b.move()
	}
	b.guard.x = nextX
	b.guard.y = nextY
	return false, false
}

func (g *guard) turn() {
	g.dir = (g.dir + 1) % 4
}

func (g *guard) next() (x, y int) {
	switch g.dir {
	case up:
		return g.x, g.y - 1
	case down:
		return g.x, g.y + 1
	case left:
		return g.x - 1, g.y
	case right:
		return g.x + 1, g.y
	default:
		panic("unexpected direction")
	}
}

func (b board) String() string {
	var out string
	for _, row := range b.grid {
		out = fmt.Sprintf("%v\n%v", out, string(row))
	}
	return out
}

func part1(b board) (int, error) {
	for {
		if done, _ := b.move(); done {
			break
		}
	}

	moves := 0
	for _, row := range b.grid {
		for _, r := range row {
			if r == 'X' {
				moves++
			}
		}
	}
	return moves, nil
}

func part2(b board) (int, error) {
	boardPart1 := b.copy()

	// Solve like part 1 to get possible obstacle placements
	for {
		if done, _ := boardPart1.move(); done {
			break
		}
	}

	var infLoops int
	for y, row := range b.grid {
		for x, r := range row {
			if r != 'X' {
				continue
			}
			boardWithObstacle := b.copy()
			boardWithObstacle.grid[y][x] = '#'

			for {
				done, loop := boardWithObstacle.move()
				if loop {
					infLoops++
				}
				if done {
					break
				}
			}
		}
	}
	return infLoops, nil
}

func (b board) copy() board {
	board := board{grid: [][]rune{},
		guard: guard{
			x:   b.guard.x,
			y:   b.guard.y,
			dir: b.guard.dir,
		},
		seenObstacles: map[guard]bool{},
	}
	for y, row := range b.grid {
		board.grid = append(board.grid, []rune{})
		board.grid[y] = append(board.grid[y], row...)
	}
	return board
}
