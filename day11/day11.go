package day11

import (
	"aoc24/utils"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

func parseLine(line string) (stones []int) {
	for _, s := range strings.Fields(line) {
		n, _ := strconv.Atoi(s)
		stones = append(stones, n)
	}
	return
}

var nextMap = map[int][][]int{}

// var nMap = map[int][]int{}

// func bl(stone int) (nextStones []int) {
// 	if stone == 0 {
// 		return append(nextStones, 1)
// 	}

// 	stoneMap, ok := nMap[stone]
// 	if ok {
// 		return stoneMap
// 	}

// 	str := strconv.Itoa(stone)
// 	digits := len(str)
// 	if digits%2 == 0 {
// 		exp := int(math.Pow10(digits / 2))
// 		nextStones = append(nextStones, stone/exp, stone%exp)
// 	} else {
// 		nextStones = append(nextStones, stone*2024)
// 	}

// 	nMap[stone] = nextStones
// 	return
// }

// func bloop(stone, iters int) {
// 	stones := []int{stone}
// 	for range iters {
// 		i := 0
// 		for {
// 			if i >= len(stones) {
// 				break
// 			}

// 			nextStones := bl(stones[i])
// 			stones[i] = nextStones[0]
// 			if len(nextStones) > 1 {
// 				stones = slices.Insert(stones, i+1, nextStones[1:]...)
// 				i++
// 			}
// 			i++
// 		}
// 	}
// }

func blink(stone, iters int /*, ch chan int*/) int {
	stones := []int{stone}
	// nextMap[stone] = [][]int{}
	stoneMap, ok := nextMap[stone]
	// start := iters
	if ok {
		if len(stoneMap) > iters {
			return len(stoneMap[iters])
			// ch <- len(stoneMap[iters])
		} else if len(stoneMap) > 0 {
			iters -= len(stoneMap) - 1
			stones = append([]int{}, stoneMap[len(stoneMap)-1]...)
		}
	} else {
		nextMap[stone] = [][]int{}
	}
	for range iters {
		i := 0
		for {
			if i >= len(stones) {
				break
			}
			if stones[i] == 0 {
				stones[i] = 1
			} else {
				str := strconv.Itoa(stones[i])
				digits := len(str)
				// digits := int(math.Log10(float64(stones[i])) + 1)
				if digits%2 == 0 {
					exp := int(math.Pow10(digits / 2))
					// left, _ := strconv.Atoi(str[:digits/2])
					left := stones[i] / exp
					// right, _ := strconv.Atoi(str[digits/2:])
					right := stones[i] % exp
					stones[i] = left
					stones = slices.Insert(stones, i+1, right)
					i++
				} else {
					stones[i] *= 2024
				}
			}
			// next := append([]int{}, stones...)
			i++
		}
		nextMap[stone] = append(nextMap[stone], append([]int{}, stones...))
	}
	fmt.Println(len(stones))
	if stone == 773 {
		fmt.Println(nextMap[stone])
	}
	return len(stones)
	// ch <- len(stones)
}

func Day11() {
	fp, err := utils.Fetch("11")
	if err != nil {
		return
	}
	// fp := "input/test/11.txt"

	lines := make(chan string)
	go utils.Scan(fp, lines)

	// stones := []int{}
	// ch := make(chan int)
	count := 0
	total := 0

	for line := range lines {
		// stones = parseLine(line)
		for _, n := range parseLine(line) {
			count++
			total += blink(n, 25)
			// go blink(n, 30)
		}
	}

	// for length := range ch {
	// 	count--
	// 	total += length
	// 	if count == 0 {
	// 		close(ch)
	// 	}
	// }

	fmt.Println(total)
}
