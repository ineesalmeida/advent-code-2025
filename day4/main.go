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
		"02",
		part1Answer, p1duration,
		part2Answer, p2duration,
	)
}

func startMatrix(lines []string) [][]int {
	positions := make([][]int, len(lines))
	for i := range positions {
		positions[i] = make([]int, len(lines[0]))
	}
	return positions
}

func countNeighbourBoxes(lines []string, positions [][]int) {

	for i, line := range lines {
		for j, col := range line {
			if col == '@' {
				// Increase positions of the 8 nearby by 1
				if i > 0 {
					if j > 0 {
						positions[i-1][j-1] += 1
					}
					positions[i-1][j] += 1
					if j < len(line)-1 {
						positions[i-1][j+1] += 1
					}
				}

				if j > 0 {
					positions[i][j-1] += 1
				}
				if j < len(line)-1 {
					positions[i][j+1] += 1
				}

				if i < len(lines)-1 {
					if j > 0 {
						positions[i+1][j-1] += 1
					}
					positions[i+1][j] += 1
					if j < len(line)-1 {
						positions[i+1][j+1] += 1
					}
				}
			}
		}
	}
}

func part1(lines []string) any {
	positions := startMatrix(lines)
	countNeighbourBoxes(lines, positions)

	result := 0
	for i, line := range positions {
		for j, col := range line {
			if col < 4 && lines[i][j] == '@' {
				result += 1
			}
		}
	}

	return result
}

func replaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}

func part2(lines []string) any {

	result := 0

	for true {
		positions := startMatrix(lines)
		countNeighbourBoxes(lines, positions)

		// utils.PrintMatrixStr(lines)

		// Remove boxes
		count := 0
		for i, line := range positions {
			for j, col := range line {
				if col < 4 && lines[i][j] == '@' {
					lines[i] = replaceAtIndex(lines[i], '.', j)
					count += 1
				}
			}
		}
		if count == 0 {
			break
		}

		result += count
	}

	return result
}
