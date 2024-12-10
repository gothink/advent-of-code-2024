package day10

import (
	"aoc24/utils"
	"fmt"
	"slices"
)

type coords [2]int

var grid = [][]int{}
var rows = 0
var cols = 0
var trailheads = []coords{}

var trailends = map[coords][]coords{}

// up, down, left, right
var unitVecs = [4]coords{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
var part1Total = 0
var part2Total = 0

func parseLine(line string, y int) {
	ints := make([]int, len(line))
	for x, r := range line {
		n := int(r - '0')
		if n == 0 {
			trailheads = append(trailheads, coords{x, y})
		}
		ints[x] = int(r - '0')
	}
	grid = append(grid, ints)
}

func (c1 *coords) move(dir int) (c2 coords, oob bool) {
	c2[0], c2[1] = c1[0]+unitVecs[dir][0], c1[1]+unitVecs[dir][1]

	if c2[0] >= cols || c2[0] < 0 || c2[1] >= rows || c2[1] < 0 {
		return
	}

	return c2, true
}

func findNext(coord coords, from int, orig coords) {
	if grid[coord[1]][coord[0]] == 9 {
		part2Total++

		trailstarts, ok := trailends[coord]
		if !ok {
			trailends[coord] = []coords{orig}
		}
		if slices.Contains(trailstarts, orig) {
			return
		}
		trailends[coord] = append(trailends[coord], orig)

		part1Total++
		return
	}

	for i := range unitVecs {
		if from == i {
			continue
		}
		next, ok := coord.move(i)
		if !ok {
			continue
		}
		if grid[next[1]][next[0]] == grid[coord[1]][coord[0]]+1 {
			findNext(next, i^1, orig)
		}
	}
}

func Day10() {
	fp, err := utils.Fetch("10")
	if err != nil {
		return
	}

	lines := make(chan string)
	go utils.Scan(fp, lines)

	for line := range lines {
		parseLine(line, rows)
		rows++
	}

	cols = len(grid[0])

	for _, th := range trailheads {
		findNext(th, -1, th)
	}

	fmt.Println("Part 1 total:", part1Total)
	fmt.Println("Part 2 total:", part2Total)
}
