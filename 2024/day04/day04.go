package day04

type board struct {
	letters string
	width   int
	height  int
}

const searchWord string = "XMAS"

func Solution(input []string) (int, int, error) {
	b := board{}
	for _, in := range input {
		b.letters += in
		b.width = len(in)
		b.height++
	}

	return part1(b), part2(b), nil
}

func part1(b board) int {
	sums := map[direction]int{}
	lineSums := map[int]int{}
	for i, l := range b.letters {
		if l != 'X' {
			continue
		}

		for _, d := range directions {
			idx := i
			letterIdx := 1
			for b.hasNeighbor(idx, d) {
				nextLetter, nextIndex := b.getNeighbor(idx, d)
				if rune(searchWord[letterIdx]) != nextLetter {
					break
				}
				idx = nextIndex
				letterIdx++
				if letterIdx == 4 {
					sums[d]++
					lineSums[i/b.width]++
					break
				}
			}
		}
	}
	var sum int
	for _, v := range sums {
		sum += v
	}
	return sum
}

func part2(b board) int {
	var sum int
	for i, l := range b.letters {
		if l != 'A' {
			continue
		}

		if !b.hasNeighbor(i, topLeft) || !b.hasNeighbor(i, topRight) || !b.hasNeighbor(i, bottomLeft) || !b.hasNeighbor(i, bottomRight) {
			continue
		}

		tl, _ := b.getNeighbor(i, topLeft)
		br, _ := b.getNeighbor(i, bottomRight)
		if !((tl == 'M' && br == 'S') || (tl == 'S' && br == 'M')) {
			continue
		}

		tr, _ := b.getNeighbor(i, topRight)
		bl, _ := b.getNeighbor(i, bottomLeft)
		if !((tr == 'M' && bl == 'S') || (tr == 'S' && bl == 'M')) {
			continue
		}
		sum++
	}
	return sum
}

//            L  R  T  B TL TR BL BR
// ....XXMAS.    1                 1 = 2
// .SAMXMS... 1                      = 1
// ...S..A...                        = 0
// ..A.A.MS.X          1        1    = 2
// XMASAMX.MM 1  1  1                = 3
// X.....XA.A             1  1       = 2
// S.S.S.S.SS                        = 0
// .A.A.A.A.A                        = 0
// ..M.M.M.MM                        = 0
// .X.X.XMASX    1  1     3  3       = 8
//            2 +3 +2 +1 +4 +4 +1 +1 =18

type direction int

const (
	left direction = iota
	right
	top
	bottom

	topLeft // this
	topRight
	bottomLeft
	bottomRight
)

var directions = []direction{
	left, right, top, bottom,
	topLeft, topRight, bottomLeft, bottomRight,
}

func (b *board) hasNeighbor(index int, d direction) bool {
	switch d {
	case left:
		return index%b.width > 0
	case right:
		return index%b.width < b.width-1
	case top:
		return index >= b.width
	case bottom:
		return index < (b.width)*(b.height-1)
	case topLeft:
		return b.hasNeighbor(index, top) && b.hasNeighbor(index, left)
	case topRight:
		return b.hasNeighbor(index, top) && b.hasNeighbor(index, right)
	case bottomLeft:
		return b.hasNeighbor(index, bottom) && b.hasNeighbor(index, left)
	case bottomRight:
		return b.hasNeighbor(index, bottom) && b.hasNeighbor(index, right)
	default:
		return false
	}
}

func (b *board) getNeighbor(index int, d direction) (rune, int) {
	switch d {
	case left:
		return rune(b.letters[index-1]), index - 1
	case right:
		return rune(b.letters[index+1]), index + 1
	case top:
		return rune(b.letters[index-b.width]), index - b.width
	case bottom:
		return rune(b.letters[index+b.width]), index + b.width
	case topLeft:
		_, index := b.getNeighbor(index, top)
		return b.getNeighbor(index, left)
	case topRight:
		_, index := b.getNeighbor(index, top)
		return b.getNeighbor(index, right)
	case bottomLeft:
		_, index := b.getNeighbor(index, bottom)
		return b.getNeighbor(index, left)
	case bottomRight:
		_, index := b.getNeighbor(index, bottom)
		return b.getNeighbor(index, right)
	default:
		panic("unexpected get direction")
	}
}
