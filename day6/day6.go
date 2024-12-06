package day6

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type GuardMap [][]string

type MapPuzzle struct {
	grid     GuardMap
	vector   int
	position [2]int
}

var vecMarkers string = "^>v<"
var unitVecs = [4][2]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}

func solve(mp *MapPuzzle) (visited [][2]int, loop bool) {
	pos, vec := mp.position, mp.vector
	g := make(GuardMap, len(mp.grid))
	for i := range mp.grid {
		g[i] = make([]string, len(mp.grid[i]))
		copy(g[i], mp.grid[i])
	}

	for {
		if pos[0] > len(g[0])-1 || pos[0] < 0 || pos[1] > len(g)-1 || pos[1] < 0 {
			// out of bounds
			break
		}

		if g[pos[1]][pos[0]] == "." {
			visited = append(visited, pos)
			// add vector marker to cell
			g[pos[1]][pos[0]] = vecMarkers[vec : vec+1]
		} else if g[pos[1]][pos[0]] == "#" {
			// go back one cell
			pos[0] -= unitVecs[vec][0]
			pos[1] -= unitVecs[vec][1]

			// update unit vector
			vec++
			if vec >= len(unitVecs) {
				vec = 0
			}
		} else if strings.Contains(g[pos[1]][pos[0]], vecMarkers[vec:vec+1]) {
			// returning on a previous path, in a loop
			loop = true
			break
		} else {
			// append vector marker to cell
			g[pos[1]][pos[0]] += vecMarkers[vec : vec+1]
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
		line := []string{}
		for j, r := range scanner.Text() {
			if idx := strings.IndexRune(vecMarkers, r); idx > -1 {
				mp.vector = idx
				mp.position[0] = j
				mp.position[1] = i
				r = '.'
			}

			line = append(line, string(r))
		}

		mp.grid = append(mp.grid, line)
		i++
	}

	return mp
}

func Day6() {
	mp := newMap("input/6.txt")
	visited, loop := solve(mp)
	if loop {
		fmt.Println("Guard is stuck in a loop!")
	} else {
		fmt.Println("Visited: ", len(visited))
	}

	numLoops := 0
	for _, coords := range visited {
		mp.grid[coords[1]][coords[0]] = "#"
		if _, loop := solve(mp); loop {
			numLoops++
		}
		mp.grid[coords[1]][coords[0]] = "."
	}

	fmt.Println("Possible loops: ", numLoops)
}
