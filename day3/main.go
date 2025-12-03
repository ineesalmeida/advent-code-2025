package main

import (
	"advent-code/aoc2025/utils"
	"fmt"
	"log"
	"math"
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

func strToInts(s string) []int {
	ints := make([]int, len(s))
	for i, c := range s {
		b, err := strconv.Atoi(string(c))
		if err != nil {
			fmt.Errorf("ERROR")
		}
		ints[i] = b
	}
	return ints
}

func firstMax(l []int) (int, int) {
	max := -1
	index := -1
	for i, n := range l {
		if n > max {
			max = n
			index = i
		}
	}
	return max, index
}

func findMaxBattery(batteries_str string) int {
	batteries := strToInts(batteries_str)
	l_max, l_i := firstMax(batteries[:len(batteries)-1])
	r_max, _ := firstMax(batteries[l_i+1:])

	return l_max*10 + r_max
}

func part1(lines []string) any {
	result := 0
	for _, line := range lines {
		result += findMaxBattery(line)
	}
	return result
}

func findMaxBattery2(batteries_str string) int {
	batteries := strToInts(batteries_str)
	max := 0

	l_i := 0
	next_l_i := -1
	l_max := -1
	for i := range 12 {
		l_max, next_l_i = firstMax(batteries[l_i : len(batteries)-12+i+1])
		l_i += next_l_i + 1
		max += l_max * int(math.Pow(float64(10), float64(12-i-1)))
	}
	// fmt.Println(max)
	return max
}

func part2(lines []string) any {
	result := 0
	for _, line := range lines {
		result += findMaxBattery2(line)
	}
	return result
}
