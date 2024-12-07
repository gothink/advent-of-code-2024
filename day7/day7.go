package day7

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
)

func findResult(nums []uint64, opMask int) (result uint64) {
	result += nums[0]

	for i, n := range nums[1:] {
		if opMask&(1<<i) == 0 {
			result += n
		} else {
			result *= n
		}
	}

	return
}

func findResPart2(nums []uint64, opMap []int) (result uint64) {
	result += nums[0]

	for i, n := range nums[1:] {
		switch opMap[i] {
		case 0:
			result += n
		case 1:
			result *= n
		case 2:
			result, _ = strconv.ParseUint(strconv.FormatUint(result, 10)+strconv.FormatUint(n, 10), 10, 64)
		}
	}

	return
}

func parseLine(line string) (goal uint64, nums []uint64) {
	for i, v := range strings.Fields(line) {
		if i == 0 {
			if n, err := strconv.Atoi(v[:len(v)-1]); err == nil {
				goal = uint64(n)
			}
		} else if n, err := strconv.Atoi(v); err == nil {
			nums = append(nums, uint64(n))
		}
	}

	return
}

func solve(goal uint64, nums []uint64) (r uint64) {
	// bitmask for operations, 0 = add (+); 1 = mul (*)
	opMask := int(math.Exp2(float64(len(nums) - 1)))

	for i := range opMask {
		if result := findResult(nums, i); result == goal {
			r = result
			break
		}
	}

	return
}

func solvePart2(goal uint64, nums []uint64) (r uint64) {
	numOps := len(nums) - 1
	opMap := make([]int, numOps)
	for i := range int(math.Pow(3, float64(numOps))) {
		num := i
		for j := 0; j < numOps; j++ {
			opMap[j] = num % 3
			num /= 3
		}
		if result := findResPart2(nums, opMap); result == goal {
			r = result
			break
		}
	}

	return
}

func Day7() {
	file, _ := os.Open("input/7.txt")
	defer file.Close()

	part1 := false

	var wg sync.WaitGroup
	results := make(chan uint64)

	go func() {
		wg.Wait()
		close(results)
	}()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		goal, nums := parseLine(scanner.Text())
		wg.Add(1)
		go func() {
			defer wg.Done()
			var result uint64
			if part1 {
				result = solve(goal, append([]uint64{}, nums...))
			} else {
				result = solvePart2(goal, append([]uint64{}, nums...))
			}
			results <- result
		}()
	}

	var total uint64 = 0
	for c := range results {
		total += c
	}

	fmt.Println("Total: ", total)
}
