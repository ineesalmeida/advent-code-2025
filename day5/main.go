package main

import (
	"advent-code/aoc2025/utils"
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"
	"time"
)

func main() {
	lines, err := utils.FileToLines("input.txt")
	if err != nil {
		log.Fatalf("[!] %s\n", err)
	}

	p1start := time.Now()
	part1Answer := part1(lines)
	p1duration := time.Since(p1start)

	p2start := time.Now()
	part2Answer := part2(lines)
	p2duration := time.Since(p2start)

	fmt.Printf("[+] day %s\n> part 1: %d (%s)\n> part 2: %d (%s)\n",
		"05",
		part1Answer, p1duration,
		part2Answer, p2duration,
	)
}

func getFreshRanges(lines []string) (int, [][2]int) {
	freshRanges := [][2]int{}
	for i, line := range lines {
		if line == "" {
			return i, freshRanges
		}
		r := strings.Split(line, "-")
		r0, _ := strconv.Atoi(r[0])
		r1, _ := strconv.Atoi(r[1])
		freshRanges = append(freshRanges, [2]int{r0, r1})
	}

	fmt.Errorf("Didn't find new line???")
	return 0, freshRanges
}

func part1(lines []string) any {
	l, freshRanges := getFreshRanges(lines)

	// Get fresh veggies
	result := 0
	for _, line := range lines[l:] {
		n, _ := strconv.Atoi(line)
		for _, r := range freshRanges {
			if n >= r[0] && n <= r[1] {
				result += 1
				break
			}
		}
	}

	return result
}

func part2(lines []string) any {

	// TODO optimize getting the freshRanges after sorting
	_, freshRanges := getFreshRanges(lines)

	r0 := make([]int, len(freshRanges))
	r1 := make([]int, len(freshRanges))
	for i, r := range freshRanges {
		r0[i] = r[0]
		r1[i] = r[1]
	}

	slices.Sort(r0)
	slices.Sort(r1)

	for i := range len(freshRanges) {
		freshRanges[i] = [2]int{r0[i], r1[i]}
	}

	result := 0
	for i, r := range freshRanges {

		result += r[1] - r[0] + 1
		if i > 0 {
			result -= utils.Max(0, freshRanges[i-1][1]-r[0]+1)
		}
	}

	return result
}

// 3, 5 -

// 1-   10
// 50  0-25  --> 10 - overlap
// 15 - 25  --> 10 - overlap2
// 2 5 -7  0  // 500-2  -->
// 550
//7  7  - 70
