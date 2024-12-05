package day5

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func total(strs []string) (acc int) {
	for _, midVal := range strs {
		n, _ := strconv.Atoi(midVal)
		acc += n
	}

	return
}

func Day5() {
	file, _ := os.Open("input/5.txt")
	defer file.Close()

	ruleMap := map[string][]string{}
	inRules := true
	badMids := []string{}
	goodMids := []string{}

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
						badUpdate = true
						break
					}
				}
			}

			if beforeIdx > -1 {
				seen = slices.Insert(seen, beforeIdx, val)
			} else {
				seen = append(seen, val)
			}
		}

		if badUpdate {
			badMids = append(badMids, seen[len(seen)/2])
		} else {
			goodMids = append(goodMids, seen[len(seen)/2])
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

	badMidTotal := total(badMids)
	goodMidTotal := total(goodMids)

	fmt.Println("Part 1 total: ", goodMidTotal)
	fmt.Println("Part 2 total: ", badMidTotal)
}
