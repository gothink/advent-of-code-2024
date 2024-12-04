package day4

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type LetterGrid [][]string

func letterGridFromFile(filePath string) LetterGrid {
	file, _ := os.Open(filePath)
	defer file.Close()

	var lg LetterGrid

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		letters := strings.Split(scanner.Text(), "")
		lg = append(lg, letters)
	}
	return lg
}

func parseGrid(lg LetterGrid) int {
	checkDirs := [9][2]int{
		{-1, -1},
		{0, -1},
		{1, -1},
		{-1, 0},
		{1, 0},
		{-1, 1},
		{0, 1},
		{1, 1}}

	maxY := len(lg)
	maxX := len(lg[0])
	nextInLine := func(i int, j int, dir [2]int) (string, [2]int) {
		x, y := j+dir[0], i+dir[1]
		if x > -1 && x < maxX && y > -1 && y < maxY {
			fmt.Printf("(%d, %d) -> (%d, %d) | ", j, i, x, y)
			return lg[x][y], [2]int{x, y}
		}
		return "", [2]int{x, y}
	}

	numMatches := 0

	for i, row := range lg {
		for j, letter := range row {
			if letter == "X" {
				for _, dir := range checkDirs {
					nextLetter, coords := nextInLine(i, j, dir)
					if nextLetter == "M" {
						nextLetter, coords = nextInLine(coords[0], coords[1], dir)
						if nextLetter == "A" {
							nextLetter, coords = nextInLine(coords[0], coords[1], dir)
							if nextLetter == "S" {
								numMatches++
								fmt.Printf("Found match: (%d, %d)\n", j, i)
							}
						}
					}
				}
			}
		}
	}
	return numMatches
}

func Day4() {
	lg := letterGridFromFile("input/test/4.txt")
	fmt.Println(lg)
	fmt.Println("Num matches: ", parseGrid(lg))
}
