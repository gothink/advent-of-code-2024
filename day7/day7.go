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

func findResult(nums []uint64, oppMask int) (result uint64) {
	result += nums[0]

	for i, n := range nums[1:] {
		if oppMask&(1<<i) == 0 {
			result += n
		} else {
			result *= n
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

func Day7() {
	file, _ := os.Open("input/7.txt")
	defer file.Close()

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
			result := solve(goal, append([]uint64{}, nums...))
			results <- result
		}()
	}

	var total uint64 = 0
	for c := range results {
		total += c
	}

	fmt.Println("Total: ", total)
}
