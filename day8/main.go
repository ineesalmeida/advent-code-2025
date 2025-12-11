package main

import (
	"advent-code/aoc2025/utils"
	"fmt"
	"log"
	"sort"
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
		"08",
		part1Answer, p1duration,
		part2Answer, p2duration,
	)
}

type Box struct {
	X       int
	Y       int
	Z       int
	cluster *Cluster
}

type Pair struct {
	A        *Box
	B        *Box
	distance int
}

type Cluster struct {
	size int
}

func parseInput(lines []string) []*Box {
	boxes := make([]*Box, len(lines))
	for i, line := range lines {
		coords := strings.Split(line, ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		z, _ := strconv.Atoi(coords[2])

		// it self is a connection
		boxes[i] = &Box{X: x, Y: y, Z: z}
	}
	return boxes
}

func distance(A *Box, B *Box) int {
	dx := A.X - B.X
	dy := A.Y - B.Y
	dz := A.Z - B.Z
	return dx*dx + dy*dy + dz*dz
}

func part1(lines []string) any {
	N_CONNECTIONS := 1000
	boxes := parseInput(lines)

	// get pairs and their distances and sort them by distance
	n_boxes := len(boxes)
	n_pairs := (n_boxes*n_boxes)/2 - len(boxes)/2
	pairs := make([]Pair, n_pairs)
	k := 0
	for i, ba := range boxes {
		for _, bb := range boxes[i+1:] {
			pairs[k] = Pair{A: ba, B: bb, distance: distance(ba, bb)}
			k += 1
		}
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].distance < pairs[j].distance
	})

	// do the connections and sort boxes list by number of connections
	clusters := make([]*Cluster, 0)
	for i := range N_CONNECTIONS {
		A, B := pairs[i].A, pairs[i].B
		if A.cluster == nil && B.cluster == nil {
			c := Cluster{size: 2}
			clusters = append(clusters, &c)
			A.cluster = &c
			B.cluster = &c
		} else if A.cluster == nil {
			A.cluster = B.cluster
			B.cluster.size += 1
		} else if B.cluster == nil {
			B.cluster = A.cluster
			A.cluster.size += 1
		} else {
			// merge clusters (keep A)
			if A.cluster == B.cluster {
				// already connected
				continue
			}

			clusterB := B.cluster
			A.cluster.size += clusterB.size
			// Update all boxes in cluster B to point to cluster A
			for _, box := range boxes {
				if box.cluster == clusterB {
					box.cluster = A.cluster
				}
			}
			clusterB.size = 0
		}
	}

	sort.Slice(clusters, func(i, j int) bool {
		return clusters[i].size > clusters[j].size
	})

	result := 1
	for i := range 3 {
		result *= clusters[i].size
	}

	return result
}

func part2(lines []string) any {
	N_CONNECTIONS := 1000
	boxes := parseInput(lines)

	// get pairs and their distances and sort them by distance
	n_boxes := len(boxes)
	n_pairs := (n_boxes*n_boxes)/2 - len(boxes)/2
	pairs := make([]Pair, n_pairs)
	k := 0
	for i, ba := range boxes {
		for _, bb := range boxes[i+1:] {
			pairs[k] = Pair{A: ba, B: bb, distance: distance(ba, bb)}
			k += 1
		}
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].distance < pairs[j].distance
	})

	// do the connections and sort boxes list by number of connections
	clusters := make([]*Cluster, 0)
	for i := range N_CONNECTIONS {
		A, B := pairs[i].A, pairs[i].B
		if A.cluster == nil && B.cluster == nil {
			c := Cluster{size: 2}
			clusters = append(clusters, &c)
			A.cluster = &c
			B.cluster = &c
		} else if A.cluster == nil {
			A.cluster = B.cluster
			B.cluster.size += 1
		} else if B.cluster == nil {
			B.cluster = A.cluster
			A.cluster.size += 1
		} else {
			// merge clusters (keep A)
			if A.cluster == B.cluster {
				// already connected
				continue
			}

			clusterB := B.cluster
			A.cluster.size += clusterB.size
			// Update all boxes in cluster B to point to cluster A
			for _, box := range boxes {
				if box.cluster == clusterB {
					box.cluster = A.cluster
				}
			}
			clusterB.size = 0
		}
	}

	// ===== actual part 2 =====
	// Above tihs is just a full copy of part 1
	// Here I am brute forcing by just continue going through the full pairs
	// list until cluster size == boxes size
	// TODO: refactor
	for _, pair := range pairs {
		A, B := pair.A, pair.B
		// fmt.Println(A.cluster, B.cluster)
		if A.cluster == nil && B.cluster == nil {
			c := Cluster{size: 2}
			clusters = append(clusters, &c)
			A.cluster = &c
			B.cluster = &c
		} else if A.cluster == nil {
			A.cluster = B.cluster
			B.cluster.size += 1
		} else if B.cluster == nil {
			B.cluster = A.cluster
			A.cluster.size += 1
		} else {
			// merge clusters (keep A)
			if A.cluster == B.cluster {
				// already connected
				continue
			}

			clusterB := B.cluster
			A.cluster.size += clusterB.size
			// Update all boxes in cluster B to point to cluster A
			for _, box := range boxes {
				if box.cluster == clusterB {
					box.cluster = A.cluster
				}
			}
			clusterB.size = 0

		}

		if A.cluster.size == n_boxes {
			return A.X * B.X
		}
	}

	return 0
}
