package day3

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func Part1() {
	filePath := "input.txt"
	file, _ := os.Open(filePath)
	defer file.Close()

	pattern := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	productSum := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		matches := pattern.FindAllStringSubmatch(line, -1)

		for _, match := range matches {
			n, _ := strconv.Atoi(match[1])
			m, _ := strconv.Atoi(match[2])
			productSum += n * m
		}
	}

	fmt.Println("Total: ", productSum)
}

func Part2() {
	filePath := "input.txt"
	file, _ := os.Open(filePath)
	defer file.Close()

	pattern := regexp.MustCompile(`do\(\)|don't\(\)|mul\((\d{1,3}),(\d{1,3})\)`)
	productSum := 0
	isEnabled := true

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		matches := pattern.FindAllStringSubmatch(line, -1)

		for _, match := range matches {
			if match[0] == "do()" {
				isEnabled = true
			} else if match[0] == "don't()" {
				isEnabled = false
			} else if isEnabled && len(match) == 3 {
				n, _ := strconv.Atoi(match[1])
				m, _ := strconv.Atoi(match[2])
				productSum += n * m
			}
		}
	}

	fmt.Println("Total: ", productSum)
}
