package main

import (
	"advent-code/aoc2025/utils"
	"fmt"
	"log"
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
		"02",
		part1Answer, p1duration,
		part2Answer, p2duration,
	)
}

func getInvalidIDs(r0, r1 int) int {
	result := 0
	for i := range r1 - r0 {
		s := strconv.Itoa(r0 + i)
		if len(s)%2 == 0 {
			if s[:len(s)/2] == s[len(s)/2:] {
				result += r0 + i
			}
		}
	}

	return result
}

func getInvalidIDs2(r0, r1 int) int {
	result := 0
	for i := range r1 - r0 {
		s := strconv.Itoa(r0 + i)

		// we would only need to check prime numbers
		// a number that repeats 4 times, also repeats 2 times, etc

		primes := [4]int{2, 3, 5, 7}
		for _, p := range primes {
			if len(s)%p == 0 {
				// all numbers divisible by p
				if s == strings.Repeat(s[:len(s)/p], p) {
					result += r0 + i
					break
				}
			}
		}
	}

	// 2, 3, (4), 5, (6), 7, (8), (9)
	// 2, 3, 5, 7

	return result
}
func part1(lines []string) any {
	ranges := strings.Split(lines[0], ",")
	result := 0
	for _, r := range ranges {
		_r := strings.Split(r, "-")
		r0, _ := strconv.Atoi(_r[0])
		r1, _ := strconv.Atoi(_r[1])
		result += getInvalidIDs(r0, r1)
	}
	return result
}

func part2(lines []string) any {
	ranges := strings.Split(lines[0], ",")
	result := 0
	for _, r := range ranges {
		_r := strings.Split(r, "-")
		r0, _ := strconv.Atoi(_r[0])
		r1, _ := strconv.Atoi(_r[1])
		result += getInvalidIDs2(r0, r1)
	}
	return result
}
