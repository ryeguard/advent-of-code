package day06

import "fmt"

type direction int

const (
	up direction = iota
	right
	down
	left
)

type position struct {
	x, y int
	dir  direction
}

type board struct {
	grid  [][]rune
	guard position
}

func Solution(input []string) (int, int, error) {
	b := board{
		grid: [][]rune{},
	}
	for y, in := range input {
		b.grid = append(b.grid, []rune{})
		for x, c := range in {
			b.grid[y] = append(b.grid[y], c)
			if c == '^' {
				b.guard = position{
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

	part1, err := day06Part1(b)
	if err != nil {
		return 0, 0, fmt.Errorf("part 1: %w", err)
	}

	part2, err := day06Part2(input)
	if err != nil {
		return 0, 0, fmt.Errorf("part 1: %w", err)
	}
	return part1, part2, nil
}

func (b *board) move() (done, visited bool) {
	// mark current as visited
	b.grid[b.guard.y][b.guard.x] = 'X'

	switch b.guard.dir {
	case up:
		// check nextY position
		nextY := b.guard.y - 1
		if nextY < 0 {
			return true, false
		}

		if b.grid[nextY][b.guard.x] == '#' {
			b.guard.dir = (b.guard.dir + 1) % 4
			return b.move()
		}
		if b.grid[nextY][b.guard.x] == 'X' {
			b.guard.y--
			return false, true
		}
		b.guard.y--
		return false, false
	case down:
		nextY := b.guard.y + 1
		if nextY >= len(b.grid) {
			return true, false
		}

		if b.grid[nextY][b.guard.x] == '#' {
			b.guard.dir = (b.guard.dir + 1) % 4
			return b.move()
		}

		if b.grid[nextY][b.guard.x] == 'X' {
			b.guard.y++
			return false, true
		}

		b.guard.y++
		return false, false
	case right:
		nextX := b.guard.x + 1
		if nextX >= len(b.grid[b.guard.y]) {
			return true, false
		}

		if b.grid[b.guard.y][nextX] == '#' {
			b.guard.dir = (b.guard.dir + 1) % 4
			return b.move()
		}

		if b.grid[b.guard.y][nextX] == 'X' {
			b.guard.x++
			return false, true
		}
		b.guard.x++
		return false, false
	case left:
		nextX := b.guard.x - 1
		if nextX < 0 {
			return true, false
		}

		if b.grid[b.guard.y][nextX] == '#' {
			b.guard.dir = (b.guard.dir + 1) % 4
			return b.move()
		}
		if b.grid[b.guard.y][nextX] == 'X' {
			b.guard.x--
			return false, true
		}
		b.guard.x--
		return false, false
	default:
		panic("")
	}
}

func (b board) String() string {
	var out string
	for _, row := range b.grid {
		out = fmt.Sprintf("%v\n%v", out, string(row))
	}
	return out
}

func day06Part1(b board) (int, error) {
	moves := 1
	for {
		done, visited := b.move()
		if done {
			return moves, nil
		}
		if !visited {
			moves++
		}
	}
}

func day06Part2(input []string) (int, error) {
	return 2, nil
}
