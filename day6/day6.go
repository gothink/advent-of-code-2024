package day6

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// GridMap stores state and vector bitmask of each cell with
// the following flags:
//
// 0 -> unvisited
// 1 -> obstacle
// 2..5 -> visited, vec: u=2,r=3,d=4,l=5
//
//	N.B.: visited bits are stored in same order as `vecMarkers` and `unitVecs`
type GridMap []int

type MapPuzzle struct {
	gm       GridMap
	height   int
	width    int
	vector   int
	position [2]int
}

var vecMarkers string = "^>v<"
var unitVecs = [4][2]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}

func solve(mp *MapPuzzle, g *GridMap, first bool) (visited [][2]int, loop bool) {
	pos, vec := mp.position, mp.vector
	gm := *g

	for {
		if pos[0] >= mp.width || pos[0] < 0 || pos[1] >= mp.height || pos[1] < 0 {
			// out of bounds
			break
		}

		idx := (pos[1] * mp.width) + pos[0]

		if gm[idx] == 0 {
			if first {
				visited = append(visited, pos)
			}
			// unvisited, assign vector bitmask
			gm[idx] = 2 << vec
		} else if gm[idx] == 1 {
			// obstacle, go back one cell
			pos[0] -= unitVecs[vec][0]
			pos[1] -= unitVecs[vec][1]

			// update unit vector
			vec++
			if vec >= len(unitVecs) {
				vec = 0
			}
		} else if gm[idx]&(2<<vec) > 0 {
			// visited cell on same vector, in a loop
			loop = true
			break
		} else {
			// visited cell, add vector bitmask
			gm[idx] += 2 << vec
		}

		// advance one cell
		pos[0] += unitVecs[vec][0]
		pos[1] += unitVecs[vec][1]
	}

	return
}

func newMap(fp string) *MapPuzzle {
	file, _ := os.Open(fp)
	defer file.Close()

	mp := new(MapPuzzle)
	i := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		for j, r := range scanner.Text() {
			if r == '.' {
				mp.gm = append(mp.gm, 0)
			} else if r == '#' {
				mp.gm = append(mp.gm, 1)
			} else if idx := strings.IndexRune(vecMarkers, r); idx > -1 {
				mp.gm = append(mp.gm, 2<<idx)
				mp.vector = idx
				mp.position[0] = j + unitVecs[idx][0]
				mp.position[1] = i + unitVecs[idx][1]
			}
		}

		i++
	}

	mp.height = i
	mp.width = len(mp.gm) / i

	return mp
}

func Day6() {
	mp := newMap("input/6.txt")
	gm := make(GridMap, len(mp.gm))
	copy(gm, mp.gm)

	visited, loop := solve(mp, &gm, true)
	if loop {
		fmt.Println("Guard is already stuck in a loop!")
		os.Exit(1)
	} else {
		fmt.Println("Part 1, cells visited: ", len(visited)+1) // skips first square, add 1 to result
	}

	numLoops := 0
	for _, coords := range visited {
		idx := (coords[1] * mp.width) + coords[0]
		gm[idx] = 1
		if _, loop := solve(mp, &gm, false); loop {
			numLoops++
		}
		copy(gm, mp.gm)
	}

	fmt.Println("Part 2, possible loops: ", numLoops)
}
