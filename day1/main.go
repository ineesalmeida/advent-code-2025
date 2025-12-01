package main

import (
	"advent-code/aoc2025/utils"
	"fmt"
	"log"
	"strconv"
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
		"02",
		part1Answer, p1duration,
		part2Answer, p2duration,
	)
}

func parseCode(code string) (int, int) {
	dir := 1
	if code[0] == 'L' {
		dir = -1
	}

	distance, _ := strconv.Atoi(code[1:])

	return dir, distance
}

func part1(lines []string) any {
	current := 50 // starting position
	countZeroes := 0

	for _, line := range lines {
		dir, distance := parseCode(line)
		current += dir * distance

		for {
			if current < 0 {
				current += 100
			} else if current > 99 {
				current -= 100
			} else {
				break
			}
		}

		if current == 0 {
			countZeroes += 1
		}
	}

	return countZeroes
}

func part2(lines []string) any {
	current := 50 // starting position
	countZeroes := 0

	for _, line := range lines {
		dir, distance := parseCode(line)

		if distance > 99 {
			countZeroes += distance / 100
			distance = distance % 100
		}

		_current := current + dir*distance

		if _current%100 == 0 {
			countZeroes += 1
			_current = 0
		} else if _current < 0 {
			// Edge case where we start at 0, we don't count again
			if current != 0 {
				countZeroes += 1
			}
			_current += 100
		} else if _current > 99 {
			countZeroes += 1
			_current -= 100
		} else {
			fmt.Errorf("ERROR")
		}

		current = _current
	}

	return countZeroes
}
