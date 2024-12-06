package day6

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	// "slices"
)

type GuardMap [][]string

type MapPuzzle struct {
	Grid     GuardMap
	Height   int
	Width    int
	Vector   int
	Position [2]int
	Visited  int
}

var startMarkers string = "^>v<"

var unitVecs = [4][2]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}

func (mp *MapPuzzle) OutOfBounds() bool {
	if mp.Position[0] > mp.Width-1 || mp.Position[0] < 0 || mp.Position[1] > mp.Height-1 || mp.Position[1] < 0 {
		return true
	} else {
		return false
	}
}

func (mp *MapPuzzle) SolveMap() {
	for {
		if out := mp.OutOfBounds(); out {
			break
		}
		if mp.Grid[mp.Position[1]][mp.Position[0]] == "." {
			mp.Visited++
			mp.Grid[mp.Position[1]][mp.Position[0]] = "-"
		} else if mp.Grid[mp.Position[1]][mp.Position[0]] == "#" {
			// go back one cell
			mp.Position[0] -= unitVecs[mp.Vector][0]
			mp.Position[1] -= unitVecs[mp.Vector][1]

			// update unit vector
			mp.Vector++
			if mp.Vector >= len(unitVecs) {
				mp.Vector = 0
			}
		}

		// advance one cell
		mp.Position[0] += unitVecs[mp.Vector][0]
		mp.Position[1] += unitVecs[mp.Vector][1]
	}
}

func NewMap(fp string) *MapPuzzle {
	file, _ := os.Open(fp)
	defer file.Close()

	mp := new(MapPuzzle)
	// mp.Grid = [][]string{}
	i := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := []string{}
		for j, r := range scanner.Text() {
			if idx := strings.IndexRune(startMarkers, r); idx > -1 {
				mp.Vector = idx
				mp.Position[0] = j
				mp.Position[1] = i
				r = '-'
				mp.Visited++
			}

			line = append(line, string(r))
		}

		mp.Grid = append(mp.Grid, line)
		i++
	}

	mp.Height = len(mp.Grid)
	mp.Width = len(mp.Grid[0])

	return mp
}

func Day6() {
	mp := NewMap("input/6.txt")
	mp.SolveMap()
	// fmt.Println(mp.Grid)
	for _, row := range mp.Grid {
		fmt.Println(row)
	}
	fmt.Println("Visited: ", mp.Visited)
}
