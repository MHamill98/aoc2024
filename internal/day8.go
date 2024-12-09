package internal

import "fmt"

func Day8(input_file_path string) {
	// Read the input file
	grid := ReadCharGrid(input_file_path)

	fmt.Printf("Part1: %d\n", solve(grid, true))
	fmt.Printf("Part2: %d\n", solve(grid, false))
}

func solve(grid [][]rune, part1 bool) int {
	towerMap := buildTowerMap(grid)
	antinodes := make(map[[2]int]struct{})
	xMax := len(grid[0])
	yMax := len(grid)

	for _, nodes := range towerMap {
		for i, node1 := range nodes {
			for j, node2 := range nodes {
				if i != j {
					if part1 {
						day8part1(xMax, yMax, node1, node2, antinodes)
					} else {
						day8part2(xMax, yMax, node1, node2, antinodes)
					}
				}
			}
		}
	}

	return len(antinodes)
}

func day8part1(xMax, yMax int, node1, node2 [2]int, antinodes map[[2]int]struct{}) {
	dist := distance(node1, node2)
	if canFit(xMax, yMax, addPair(node1, dist)) {
		antinodes[addPair(node1, dist)] = struct{}{}
	}
	if canFit(xMax, yMax, subPair(node2, dist)) {
		antinodes[subPair(node2, dist)] = struct{}{}
	}
}

func day8part2(xMax, yMax int, node1, node2 [2]int, antinodes map[[2]int]struct{}) {
	antinodes[node1] = struct{}{}
	antinodes[node2] = struct{}{}
	dist := reduceDist(distance(node1, node2))
	start := addPair(node1, dist)
	for canFit(xMax, yMax, start) {
		antinodes[start] = struct{}{}
		start = addPair(start, dist)
	}
	start = subPair(node1, dist)
	for canFit(xMax, yMax, start) {
		antinodes[start] = struct{}{}
		start = subPair(start, dist)
	}
}

func buildTowerMap(grid [][]rune) map[rune][][2]int {
	towerMap := make(map[rune][][2]int)

	for y, row := range grid {
		for x, cell := range row {
			if cell != '.' {
				towerMap[cell] = append(towerMap[cell], [2]int{x, y})
			}
		}
	}

	return towerMap
}

func distance(node1, node2 [2]int) [2]int {
	xDist := node1[0] - node2[0]
	yDist := node1[1] - node2[1]
	return [2]int{xDist, yDist}
}

func canFit(xMax, yMax int, node [2]int) bool {
	return node[0] >= 0 && node[0] < xMax && node[1] >= 0 && node[1] < yMax
}

func reduceDist(dist [2]int) [2]int {
	gcd := GCD(dist[0], dist[1])
	for gcd != 1 {
		dist[0] /= gcd
		dist[1] /= gcd
		gcd = GCD(dist[0], dist[1])
	}
	return dist
}
