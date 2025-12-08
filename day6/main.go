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
		"06",
		part1Answer, p1duration,
		part2Answer, p2duration,
	)
}

func parseProblems(lines []string) ([][]int, []string) {
	problems := make([][]int, 0)
	ops := make([]string, 0)

	for _, line := range lines {
		p := []int{}
		for _, v := range strings.Split(line, " ") {
			if v != "" {
				n, _ := strconv.Atoi(v)
				p = append(p, n)
			}
			if v == "*" || v == "+" {
				ops = append(ops, v)
			}
		}
		if len(ops) == 0 {
			problems = append(problems, p)
		}
	}
	return problems, ops
}

func part1(lines []string) any {
	problems, ops := parseProblems(lines)

	result := 0
	for i, op := range ops {
		r := 0
		if op == "*" {
			r = 1
			for j := range problems {
				// fmt.Println(" - ", problems[j][i])
				r *= problems[j][i]
			}
		} else {
			for j := range problems {
				// fmt.Println(" - ", problems[j][i])
				r += problems[j][i]
			}
		}
		result += r
	}

	return result
}

func parseProblems2(lines []string) ([][]int, []string) {
	opLine := lines[len(lines)-1]
	ops := make([]string, 0)
	for _, ch := range opLine {
		if ch == '*' || ch == '+' {
			ops = append(ops, string(ch))
		}
	}

	cols := make([]string, len(lines[0]))
	for i := 0; i < len(lines)-1; i++ {
		for j := 0; j < len(lines[i]); j++ {
			cols[j] += string(lines[i][j])
		}
	}

	problems := make([][]int, 0)
	curr := make([]int, 0)

	for col := len(cols) - 1; col >= 0; col-- {
		colStr := strings.TrimSpace(cols[col])

		if colStr == "" {
			if len(curr) > 0 {
				problems = append(problems, curr)
				curr = make([]int, 0)
			}
		} else {
			num, _ := strconv.Atoi(colStr)
			curr = append(curr, num)
		}
	}

	if len(curr) > 0 {
		problems = append(problems, curr)
	}

	for i := 0; i < len(ops)/2; i++ {
		ops[i], ops[len(ops)-1-i] = ops[len(ops)-1-i], ops[i]
	}

	return problems, ops
}

func part2(lines []string) any {
	problems, ops := parseProblems2(lines)

	result := 0
	for i, p := range problems {
		var r int
		if ops[i] == "*" {
			r = 1
			for _, n := range p {
				r *= n
			}
		} else {
			r = 0
			for _, n := range p {
				r += n
			}
		}
		result += r
	}

	return result
}
