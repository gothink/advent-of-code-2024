package day9

import (
	"aoc24/utils"
	"fmt"
	"strconv"
)

var fileSizes = []int{}
var freeSizes = []int{}
var blockOffsets = []int{}
var blockStarts = []int{}
var diskSize = 0
var checksum = 0
var currentBlock = 0

func calcChecksum(fileID, numBlocks int) (done bool) {
	for range numBlocks {
		checksum += currentBlock * fileID
		currentBlock++
		if currentBlock >= diskSize {
			return true
		}
	}
	return
}

func solvePart1() {
	currentID := len(fileSizes) - 1
	fileSize := fileSizes[currentID]

	nextFile := func() {
		currentID--
		fileSize = fileSizes[currentID]
	}

outer:
	for i, freeSize := range freeSizes {
		if done := calcChecksum(i, fileSizes[i]); done {
			break
		}

		if freeSize == 0 {
			continue
		}

		if fileSize == 0 {
			nextFile()
		}

		for {
			if fileSize >= freeSize {
				if done := calcChecksum(currentID, freeSize); done {
					break outer
				}
				fileSize -= freeSize
				break
			} else {
				if done := calcChecksum(currentID, fileSize); done {
					break outer
				}
				freeSize -= fileSize
				nextFile()
			}
		}
	}
}

func calcChecksum2(fileID, numBlocks, startBlock int) {
	for i := startBlock; i < startBlock+numBlocks; i++ {
		checksum += i * fileID
	}
}

func solvePart2() {
	for i := len(fileSizes) - 1; i >= 0; i-- {
		hasMoved := false
		for j := 0; j < i; j++ {
			if freeSizes[j] >= fileSizes[i] {
				freeSizes[j] -= fileSizes[i]
				calcChecksum2(i, fileSizes[i], blockOffsets[j])
				blockOffsets[j] += fileSizes[i]
				blockOffsets[i] -= fileSizes[i]
				hasMoved = true
				break
			}
		}

		if !hasMoved {
			calcChecksum2(i, fileSizes[i], blockStarts[i])
		}
	}
}

func parseLine(line string) {
	blockOffset, blockStart := 0, 0

	for i := 0; i < len(line); i += 2 {
		fileSize, err := strconv.Atoi(string(line[i]))
		if err != nil {
			continue
		}

		diskSize += fileSize
		fileSizes = append(fileSizes, fileSize)

		freeSize := 0
		if i+1 < len(line) {
			freeSize, err = strconv.Atoi(string(line[i+1]))
			if err != nil {
				continue
			}
		}

		freeSizes = append(freeSizes, freeSize)

		// part 2
		blockOffset += fileSize
		blockOffsets = append(blockOffsets, blockOffset)
		blockOffset += freeSize

		blockStarts = append(blockStarts, blockStart)
		blockStart += fileSize + freeSize
	}
}

func Day9() {
	fp, err := utils.Fetch("9")
	if err != nil {
		return
	}

	lines := make(chan string)
	go utils.Scan(fp, lines)

	for line := range lines {
		parseLine(line)
	}

	part := 2
	switch part {
	case 1:
		solvePart1()
		fmt.Println("Checksum part 1: ", checksum)
	case 2:
		solvePart2()
		fmt.Println("Checksum part 2: ", checksum)
	}
}
