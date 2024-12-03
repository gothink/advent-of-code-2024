package day1

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func listsFromFile(filePath string) ([]int, []int) {
	file, _ := os.Open(filePath)
	defer file.Close()

	list1, list2 := []int{}, []int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		nums := strings.Fields(scanner.Text())
		n1, _ := strconv.Atoi(nums[0])
		list1 = append(list1, n1)
		n2, _ := strconv.Atoi(nums[1])
		list2 = append(list2, n2)
	}

	return list1, list2
}

func totalDistance(list1 []int, list2 []int) int {
	sort.Ints(list1)
	sort.Ints(list2)
	distance := 0
	for i, n := range list1 {
		diff := n - list2[i]
		if diff < 0 {
			diff *= -1
		}
		distance += diff
	}
	return distance
}

func similarityScore(list1 []int, list2 []int) int {
	frequencyMap := map[int]int{}
	for _, n := range list2 {
		frequencyMap[n]++
	}

	frequencyScore := 0
	for _, n := range list1 {
		// if frequencyMap[n] > 0 {
		frequencyScore += n * frequencyMap[n]
		// }
	}
	return frequencyScore
}

func Part1() {
	list1, list2 := listsFromFile("input/1.txt")
	distance := totalDistance(list1, list2)
	fmt.Println("Total distance: ", distance)
}

func Part2() {
	list1, list2 := listsFromFile("input/1.txt")
	similarity := similarityScore(list1, list2)
	fmt.Println("Similarity score: ", similarity)
}
