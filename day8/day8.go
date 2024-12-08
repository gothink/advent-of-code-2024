package day8

import (
	"aoc24/utils"
	"fmt"
	"math/bits"
)

type coords [2]int

var freqMap = map[rune][]coords{}

func parseLine(n int, line string) {
	for j, r := range line {
		if r == '.' {
			continue
		}
		if c, ok := freqMap[r]; ok {
			freqMap[r] = append(c, [2]int{j, n})
		} else {
			freqMap[r] = append([]coords{}, [2]int{j, n})
		}
	}
}

func Day8() {
	lines := make(chan string)
	go utils.Scan("input/8.txt", lines)
	rows := 0
	cols := 0
	for line := range lines {
		if rows == 0 {
			cols = len(line)
		}
		parseLine(rows, line)
		rows++
	}

	nodeMask := make([]uint64, rows)

	addNodes := func(c1, c2 coords) {
		// add tower locations (part 2)
		nodeMask[c1[1]] |= 1 << c1[0]
		nodeMask[c2[1]] |= 1 << c2[0]

		// calculate distance between towers
		diffX, diffY := c1[0]-c2[0], c1[1]-c2[1]

		// find nodes and set bit if on map
		nodeX, nodeY := c1[0]+diffX, c1[1]+diffY

		// part 2: find all nodes in line
		for {
			if nodeX >= 0 && nodeX < cols && nodeY >= 0 && nodeY < rows {
				nodeMask[nodeY] |= 1 << nodeX
				nodeX += diffX
				nodeY += diffY
			} else {
				break
			}
		}

		nodeX, nodeY = c2[0]-diffX, c2[1]-diffY

		for {
			if nodeX >= 0 && nodeX < cols && nodeY >= 0 && nodeY < rows {
				nodeMask[nodeY] |= 1 << nodeX
				nodeX -= diffX
				nodeY -= diffY
			} else {
				break
			}
		}
	}

	for _, c := range freqMap {
		if len(c) > 1 {
			// loop over tower pairs
			for i := 1; i < len(c); i++ {
				for j := range i {
					addNodes(c[i], c[j])
				}
			}
		}
	}

	total := 0
	for _, mask := range nodeMask {
		total += bits.OnesCount64(mask)
	}

	fmt.Println("Total: ", total)
}
