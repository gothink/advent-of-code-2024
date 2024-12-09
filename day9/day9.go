package day9

import (
	"aoc24/utils"
	"fmt"
	"strconv"
)

var fileSizes = []int{}
var freeSizes = []int{}
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

func parseLine(line string) {
	for i := 0; i < len(line); i += 2 {
		if fileSize, err := strconv.Atoi(string(line[i])); err == nil {
			diskSize += fileSize
			fileSizes = append(fileSizes, fileSize)
		}

		space := 0
		if i+1 < len(line) {
			freeSize, err := strconv.Atoi(string(line[i+1]))
			if err != nil {
				continue
			}
			space = freeSize
		}
		freeSizes = append(freeSizes, space)
	}

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

func Day9() {
	fp, err := utils.Fetch("9")
	if err != nil {
		return
	}
	lines := make(chan string)
	// fp := "input/test/9.txt"
	go utils.Scan(fp, lines)

	for line := range lines {
		parseLine(line)
	}
	fmt.Println("Checksum: ", checksum)
}
