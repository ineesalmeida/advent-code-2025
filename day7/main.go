package main

import (
	"advent-code/aoc2025/utils"
	"fmt"
	"log"
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
		"07",
		part1Answer, p1duration,
		part2Answer, p2duration,
	)
}

func splitBeam(index int, lines []string, seen map[[2]int]bool) int {
	// reached the end
	if len(lines) == 0 {
		return 0
	}

	result := 0
	for i, line := range lines {
		if line[index] == '^' {
			// has it been seen before?

			// a bit of a workaround because I was too lazy to refactor to add row indexes,
			// but seen stores index (column) + distance from the bottom (row)
			_, ok := seen[[2]int{index, len(lines) - i}]
			if ok {
				return 0
			}

			result += 1
			seen[[2]int{index, len(lines) - i}] = true

			// check left
			if index > 0 {
				result += splitBeam(index-1, lines[i+1:], seen)
			}

			// check right
			if index < len(lines[0])-1 {
				result += splitBeam(index+1, lines[i+1:], seen)
			}
			break
		}
	}

	return result
}

func part1(lines []string) any {

	// get starting point
	start := 0
	for i, c := range lines[0] {
		if c == 'S' {
			start = i
			break
		}
	}

	seen := make(map[[2]int]bool)
	return splitBeam(start, lines[1:], seen)
}

func alternateTimelines(index int, lines []string, cache map[[2]int]int) int {
	// reached the end
	if len(lines) == 0 {
		return 1
	}

	value, ok := cache[[2]int{index, len(lines)}]
	if ok {
		return value
	}

	result := 0
	for i, line := range lines {
		if line[index] == '^' {

			// check left
			if index > 0 {
				result += alternateTimelines(index-1, lines[i+1:], cache)
			}

			// check right
			if index < len(lines[0])-1 {
				result += alternateTimelines(index+1, lines[i+1:], cache)
			}

			cache[[2]int{index, len(lines)}] = result
			return result
		}
	}

	// if we reach here, we went to the end so +1
	return result + 1
}

func part2(lines []string) any {

	// get starting point
	start := 0
	for i, c := range lines[0] {
		if c == 'S' {
			start = i
			break
		}
	}

	cache := make(map[[2]int]int)
	return alternateTimelines(start, lines[1:], cache)
}
