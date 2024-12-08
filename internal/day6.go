package internal

import "fmt"

func Day6(input_file_path string) {
	grid := ReadCharGrid(input_file_path)

	fmt.Printf("Part1: %d\n", day6part1(grid))
	fmt.Printf("Part2: %d\n", day6part2())

}

func day6part1(grid [][]rune) int {
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
	return total
}

func day6part2() int {
	return 0
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
