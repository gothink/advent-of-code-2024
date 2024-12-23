package day20

import (
	"aoc24/utils"
	"fmt"
	"strings"
)

type cell struct {
	seen      bool
	val       string
	order     int
	shortcuts []coords
}

type cellGrid map[coords]cell

type coords [2]int

var rows = 0
var cols = 0

var startCoords coords
var endCoords coords

var grid = cellGrid{}
var pathCount = 0
var shortcutCount = 0

// up, down, left, right
var unitVecs = [4]coords{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}

func (g *cellGrid) at(c coords) (cell cell, ok bool) {
	if c[0] >= cols || c[0] < 0 || c[1] >= rows || c[1] < 0 {
		return
	}

	if cell, ok = (*g)[c]; ok {
		return cell, true
	}

	return
}

func (c *coords) move(dir int, num int) coords {
	return coords{
		c[0] + (unitVecs[dir][0] * num),
		c[1] + (unitVecs[dir][1] * num),
	}
}

func parseLine(line string, rows int) {
	for col, s := range strings.SplitAfter(line, "") {
		if rows == 0 {
			cols++
		}
		if s == "S" {
			startCoords = coords{col, rows}
			s = "."

		} else if s == "E" {
			endCoords = coords{col, rows}
			s = "."
		}

		grid[coords{col, rows}] = cell{val: s}
	}
}

func isShortcut(loc coords, dir int) bool {
	toCoords := loc.move(dir, 1)
	fromCoords := loc.move(dir, -1)
	if shortcut, ok := grid.at(toCoords); ok {
		if !shortcut.seen && shortcut.val == "." {
			if len(shortcut.shortcuts) > 0 {
				shortcut.shortcuts = append(shortcut.shortcuts, fromCoords)
			} else {
				shortcut.shortcuts = []coords{fromCoords}
			}
			grid[toCoords] = shortcut
			return true
		}
	}
	return false
}

func checkCell(loc coords, dir int) (nextCoords coords, nextDir int) {
	cell, ok := grid.at(loc)
	if !ok {
		return
	}

	cell.seen = true
	cell.order = pathCount
	pathCount++

	// process shortcuts to this cell
	if len(cell.shortcuts) > 0 {
		for _, shortCoords := range cell.shortcuts {
			shortcut, ok := grid.at(shortCoords)
			if !ok {
				continue
			}
			if cell.order-shortcut.order-2 >= 100 {
				shortcutCount++
			}
		}
	}

	// advance one cell
	nextDir = dir
	nextCoords = loc.move(dir, 1)

	// check adjacent cells
	oppDir := (dir & 2) ^ 2
	for i := oppDir; i < oppDir+2; i++ {
		checkLoc := loc.move(i, 1)
		if grid[checkLoc].val == "#" {
			isShortcut(checkLoc, i)
		} else {
			isShortcut(loc.move(dir, 1), dir)
			nextDir = i
			nextCoords = loc.move(i, 1)
		}
	}

	grid[loc] = cell

	return
}

func traversePath() {
	nextCoords, dir := checkCell(startCoords, 0)

	for {
		nextCoords, dir = checkCell(nextCoords, dir)
		if nextCoords == endCoords {
			checkCell(nextCoords, dir)
			break
		}

	}

}

func Day20() {
	fp, err := utils.Fetch("20")
	if err != nil {
		fmt.Println(err)
		return
	}

	lines := make(chan string)
	go utils.Scan(fp, lines)

	for line := range lines {
		parseLine(line, rows)
		rows++
	}

	fmt.Println(startCoords, endCoords, rows, cols)

	traversePath()
	fmt.Println(shortcutCount)

}
