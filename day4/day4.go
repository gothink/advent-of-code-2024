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

// part 1
func parseGrid(lg LetterGrid) int {
	checkDirs := [8][2]int{{-1, -1}, {0, -1}, {1, -1}, {-1, 0}, {1, 0}, {-1, 1}, {0, 1}, {1, 1}}

	maxY := len(lg)
	maxX := len(lg[0])
	nextInLine := func(i int, j int, dir [2]int) (string, [2]int) {
		x, y := j+dir[0], i+dir[1]
		if x > -1 && x < maxX && y > -1 && y < maxY {
			return lg[y][x], [2]int{x, y}
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
						nextLetter, coords = nextInLine(coords[1], coords[0], dir)
						if nextLetter == "A" {
							nextLetter, coords = nextInLine(coords[1], coords[0], dir)
							if nextLetter == "S" {
								numMatches++
							}
						}
					}
				}
			}
		}
	}
	return numMatches
}

// part 2
func findMAS(lg LetterGrid) int {
	numMatches := 0

	for i, row := range lg[:len(lg)-2] {
		for j, letter := range row[:len(row)-2] {
			// part 2
			if (letter == "M" || letter == "S") && lg[i+1][j+1] == "A" {
				opp := "S"
				if letter == "S" {
					opp = "M"
				}
				if (lg[i][j+2] == letter && lg[i+2][j] == opp && lg[i+2][j+2] == opp) ||
					(lg[i+2][j] == letter && lg[i][j+2] == opp && lg[i+2][j+2] == opp) {
					numMatches++
					// break match detection on other 3 corners by removing middle 'A'
					lg[i+1][j+1] = "."
				}
			}
		}
	}

	return numMatches
}

func Day4() {
	lg := letterGridFromFile("input/4.txt")
	fmt.Println("Part 1 matches: ", parseGrid(lg))
	fmt.Println("Part 2 matches: ", findMAS(lg))
}
