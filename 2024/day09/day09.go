package day09

import (
	"fmt"
	"strconv"
)

type disk struct {
	diskMap []int
	pos     int

	nextIdx  int
	blockIdx int

	lastIdx      int
	lastBlockIdx int
}

type block struct {
	id  *int
	pos int
}

func (d *disk) next() *block {
	if d.nextIdx >= len(d.diskMap) {
		return nil
	}

	var b block = block{pos: d.pos}

	if d.nextIdx%2 == 0 {
		// at a file
		b.id = ptr(d.nextIdx / 2)
	}

	d.blockIdx++
	d.pos++

	// if next will be beyond block, increment
	if d.blockIdx >= d.diskMap[d.nextIdx] {
		d.blockIdx = 0
		d.nextIdx++
	}

	return &b
}

func (d *disk) lastFile() *int {
	if d.lastIdx < 0 {
		return nil
	}
	if d.lastIdx+(d.diskMap[d.lastIdx]-d.lastBlockIdx) < d.pos/2 {
		return nil
	}

	ret := d.lastIdx / 2
	d.lastBlockIdx++
	if d.lastBlockIdx >= d.diskMap[d.lastIdx] {
		d.lastBlockIdx = 0
		d.lastIdx -= 2
	}
	return &ret
}

func Solution(input []string) (int, int, error) {
	d := disk{}

	for _, r := range input[0] {
		n, err := strconv.Atoi(string(r))
		if err != nil {
			return 0, 0, fmt.Errorf("atoi: %w", err)
		}
		d.diskMap = append(d.diskMap, n)
	}
	d.lastIdx = len(d.diskMap) - 1

	part1, err := part1(d)
	if err != nil {
		return 0, 0, fmt.Errorf("part 1: %w", err)
	}

	part2, err := part2(input)
	if err != nil {
		return 0, 0, fmt.Errorf("part 2: %w", err)
	}
	return part1, part2, nil
}

func part1(d disk) (int, error) {
	checksum := 0
	next := d.next()
	for next != nil {
		if next.id != nil {
			// if file, add it to checksum
			fmt.Println("adding 1:", d.pos-1, "*", *next.id)
			checksum += (d.pos - 1) * *next.id
		} else {
			// if empty space,
			// get next file from end and add that to the checksum
			last := d.lastFile()
			if last == nil {
				return checksum, nil
			} else {
				// checksum += *last * d.diskMap[*nex]
				fmt.Println("adding 2:", d.pos-1, "*", *last)
				checksum += (d.pos - 1) * *last
			}
		}
		next = d.next()
	}
	return checksum, nil
}

func part2(input []string) (int, error) {
	return len(input) + 2, nil
}

func ptr(v int) *int {
	return &v
}
