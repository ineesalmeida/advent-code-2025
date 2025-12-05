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

func decreaseNeighbours(positions [][]int, i, j int, maxRows, maxCols int) {
	// Decrease positions of the 8 nearby by 1
	if i > 0 {
		if j > 0 {
			positions[i-1][j-1] -= 1
		}
		positions[i-1][j] -= 1
		if j < maxCols-1 {
			positions[i-1][j+1] -= 1
		}
	}

	if j > 0 {
		positions[i][j-1] -= 1
	}
	if j < maxCols-1 {
		positions[i][j+1] -= 1
	}

	if i < maxRows-1 {
		if j > 0 {
			positions[i+1][j-1] -= 1
		}
		positions[i+1][j] -= 1
		if j < maxCols-1 {
			positions[i+1][j+1] -= 1
		}
	}
}

func part2(lines []string) any {
	// Build position matrix once
	positions := startMatrix(lines)
	countNeighbourBoxes(lines, positions)

	maxRows := len(lines)
	maxCols := len(lines[0])
	result := 0

	// Use a queue to track boxes to check
	type coord struct {
		i, j int
	}
	toCheck := make([]coord, 0, maxRows*maxCols)

	// Initially add all boxes with < 4 neighbors
	for i := 0; i < maxRows; i++ {
		for j := 0; j < maxCols; j++ {
			if lines[i][j] == '@' && positions[i][j] < 4 {
				toCheck = append(toCheck, coord{i, j})
			}
		}
	}

	// Process queue
	for len(toCheck) > 0 {
		c := toCheck[0]
		toCheck = toCheck[1:]

		// Check if still valid (might have been removed already)
		if lines[c.i][c.j] != '@' {
			continue
		}

		// Check if still has < 4 neighbors
		if positions[c.i][c.j] >= 4 {
			continue
		}

		// Remove this box
		lines[c.i] = replaceAtIndex(lines[c.i], '.', c.j)
		result++

		// Decrease neighbor counts and add neighbors that might now qualify
		neighbors := []coord{}
		if c.i > 0 {
			if c.j > 0 {
				neighbors = append(neighbors, coord{c.i - 1, c.j - 1})
			}
			neighbors = append(neighbors, coord{c.i - 1, c.j})
			if c.j < maxCols-1 {
				neighbors = append(neighbors, coord{c.i - 1, c.j + 1})
			}
		}
		if c.j > 0 {
			neighbors = append(neighbors, coord{c.i, c.j - 1})
		}
		if c.j < maxCols-1 {
			neighbors = append(neighbors, coord{c.i, c.j + 1})
		}
		if c.i < maxRows-1 {
			if c.j > 0 {
				neighbors = append(neighbors, coord{c.i + 1, c.j - 1})
			}
			neighbors = append(neighbors, coord{c.i + 1, c.j})
			if c.j < maxCols-1 {
				neighbors = append(neighbors, coord{c.i + 1, c.j + 1})
			}
		}

		for _, n := range neighbors {
			positions[n.i][n.j]--
			// If this neighbor is a box and now has < 4 neighbors, add to queue
			if lines[n.i][n.j] == '@' && positions[n.i][n.j] < 4 {
				toCheck = append(toCheck, n)
			}
		}
	}

	return result
}
