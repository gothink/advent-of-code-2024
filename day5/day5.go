package day5

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Day5() {
	file, _ := os.Open("input/test/5.txt")
	defer file.Close()

	ruleMap := map[int][]int{}
	inRules := true
	badMidTotal := 0
	// badMids := []int{}

	parseRule := func(line string) {
		before, after, _ := strings.Cut(line, "|")
		b, _ := strconv.Atoi(before)
		a, _ := strconv.Atoi(after)
		ruleMap[b] = append(ruleMap[b], a)
	}

	parseUpdate := func(vals []string) {
		seen := []int{}
		beforeIdx := -1

		for _, val := range vals {
			n, _ := strconv.Atoi(val)
			if ruleMap[n] != nil {
				for i, prev := range seen {
					if slices.Contains(ruleMap[n], prev) {
						beforeIdx = i
						break
					}
				}
			}

			// Part 2
			if beforeIdx > -1 {
				seen = slices.Insert(seen, beforeIdx, n)
			} else {
				seen = append(seen, n)
			}

			// Part 1
			// seen = append(seen, n)
		}

		// Part 2
		if beforeIdx > -1 {
			badMidTotal += seen[len(seen)/2]
		}

		// Part 1
		// if badUpdate == -1 {
		// 	midIdx := len(vals) / 2
		// 	var midVal int
		// 	if midIdx < len(seen) {
		// 		midVal = seen[midIdx]
		// 	} else {
		// 		midVal, _ = strconv.Atoi(vals[midIdx])
		// 	}
		// 	badMidTotal += midVal
		// }
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			inRules = false
			continue
		}

		if inRules {
			parseRule(line)
		} else {
			vals := strings.Split(line, ",")
			parseUpdate(vals)
		}
	}

	fmt.Println("Total: ", badMidTotal)
	// fmt.Println("Mid vals: ", badMids)
}
