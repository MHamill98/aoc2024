package internal

import (
	"fmt"
	"maps"
)

func Day6(input_file_path string) {
	grid := ReadCharGrid(input_file_path)

	part1, visited := day6part1(grid)

	fmt.Printf("Part1: %d\n", part1)
	fmt.Printf("Part2: %d\n", day6part2(grid, visited))

}

func day6part1(grid [][]rune) (int, map[[2]int]struct{}) {
	x, y := findStart(grid)
	total := 1

	directions := [][2]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
	visited := make(map[[2]int]struct{})
	visited[[2]int{x, y}] = struct{}{}

	d := 0

	for 0 <= x && x < len(grid[0]) && 0 <= y && y < len(grid) && 0 <= x+directions[d][0] && x+directions[d][0] < len(grid[0]) && 0 <= y+directions[d][1] && y+directions[d][1] < len(grid) {
		if grid[y+directions[d][1]][x+directions[d][0]] == '#' {
			d = (d + 1) % 4
		} else {
			x += directions[d][0]
			y += directions[d][1]
			if _, ok := visited[[2]int{x, y}]; !ok {
				visited[[2]int{x, y}] = struct{}{}
				total++
			}
		}
	}
	return total, visited
}

func day6part2(grid [][]rune, visited map[[2]int]struct{}) int {
	startX, startY := findStart(grid)
	total := 0

	for loc := range maps.Keys(visited) {
		x, y := loc[0], loc[1]
		if grid[y][x] == '.' && !(x == startX && y == startY) {
			grid[y][x] = '#'
			if isLoop(startX, startY, grid) {
				total++
			}
			grid[y][x] = '.'
		}
	}

	return total
}

func findStart(grid [][]rune) (int, int) {
	for y, row := range grid {
		for x, cell := range row {
			if cell == '^' {
				return x, y
			}
		}
	}
	return -1, -1
}

func isLoop(x, y int, grid [][]rune) bool {
	directions := [][2]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
	dir := 0
	visited := make(map[[3]int]struct{})
	visited[[3]int{x, y, dir}] = struct{}{}

	for 0 <= x && x < len(grid[0]) && 0 <= y && y < len(grid) && 0 <= x+directions[dir][0] && x+directions[dir][0] < len(grid[0]) && 0 <= y+directions[dir][1] && y+directions[dir][1] < len(grid) {
		if grid[y+directions[dir][1]][x+directions[dir][0]] == '#' {
			dir = (dir + 1) % 4
		} else {
			x += directions[dir][0]
			y += directions[dir][1]
			if _, ok := visited[[3]int{x, y, dir}]; ok {
				return true
			} else {
				visited[[3]int{x, y, dir}] = struct{}{}
			}
		}
	}
	return false
}
