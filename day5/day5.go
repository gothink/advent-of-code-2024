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
	file, _ := os.Open("input/5.txt")
	defer file.Close()

	ruleMap := map[string][]string{}
	inRules := true
	badMidTotal := 0
	badMids := []string{}

	parseRule := func(line string) {
		before, after, _ := strings.Cut(line, "|")
		ruleMap[before] = append(ruleMap[before], after)
	}

	parseUpdate := func(vals []string) {
		seen := []string{}
		badUpdate := false

		for _, val := range vals {
			beforeIdx := -1
			if ruleMap[val] != nil {
				for i, prev := range seen {
					if slices.Contains(ruleMap[val], prev) {
						beforeIdx = i
						break
					}
				}
			}

			// Part 2
			if beforeIdx > -1 {
				seen = slices.Insert(seen, beforeIdx, val)
				badUpdate = true
			} else {
				seen = append(seen, val)
			}

			// Part 1
			// seen = append(seen, val)
		}

		// Part 1
		// if !badUpdate {

		// Part 2
		if badUpdate {
			badMids = append(badMids, seen[len(seen)/2])
		}
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

	for _, midVal := range badMids {
		n, _ := strconv.Atoi(midVal)
		badMidTotal += n
	}

	fmt.Println("Total: ", badMidTotal)
}
