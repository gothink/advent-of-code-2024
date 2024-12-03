package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func strToInts(s string) []int {
	strs := strings.Fields(s)
	nums := []int{}
	for _, ss := range strs {
		n, _ := strconv.Atoi(ss)
		nums = append(nums, n)
	}
	return nums
}

func findUnsafeIdx(nums []int) int {
	dir := 1
	if nums[0] < nums[1] {
		dir = -1
	}
	unsafeIdx := -1

	for i, n := range nums[:len(nums)-1] {
		diff := n - nums[i+1]
		if diff == 0 || diff > 3 || diff < -3 || diff*dir < 0 {
			unsafeIdx = i
			break
		}
	}

	return unsafeIdx
}

func safetyDampener(nums []int, idx int) bool {
	idxDiffs := [3]int{idx, idx + 1, idx - 1}
	isSafe := false
	for _, i := range idxDiffs {
		dampenedNums := []int{}
		if i > 0 {
			dampenedNums = append(dampenedNums, nums[:i]...)
		}
		if i < len(nums)-1 {
			dampenedNums = append(dampenedNums, nums[i+1:]...)
		}
		unsafeIdx := findUnsafeIdx(dampenedNums)
		if unsafeIdx == -1 {
			isSafe = true
			break
		}
	}
	return isSafe
}

func getSafeReports(filePath string, dampen bool) int {
	file, _ := os.Open(filePath)
	defer file.Close()

	safeReports := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		nums := strToInts(scanner.Text())
		unsafeIdx := findUnsafeIdx(nums)

		if unsafeIdx == -1 {
			safeReports++
		} else if dampen {
			isDampened := safetyDampener(nums, unsafeIdx)
			if isDampened {
				safeReports++
			}
		}
	}

	return safeReports
}

func Day2() {
	safeReports := getSafeReports("input/2.txt", true)
	fmt.Println("Safe reports: ", safeReports)
}
