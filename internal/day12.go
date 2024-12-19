package internal

import "fmt"

func Day12(input_file_path string) {
	input := ReadCharGrid(input_file_path)

	fmt.Printf("Part1: %d\n", day12part1(input))
	fmt.Printf("Part2: %d\n", day12part2())
}

func day12part1(input [][]rune) int {
	visited := make(map[[2]int]struct{})
	total := 0
	xLim := len(input[0])
	yLim := len(input)

	for y, row := range input {
		for x := range row {
			if _, ok := visited[[2]int{x, y}]; !ok {
				size, perimeter := regionDimensions(x, y, xLim, yLim, input, visited)
				total += size * perimeter
			}
		}
	}
	return total
}

func day12part2() int {
	return 0
}

func getNeighbours(x, y, xLim, yLim int, grid [][]rune) [][2]int {
	plant := grid[y][x]
	neighbours := make([][2]int, 0, 4)
	if x-1 >= 0 && grid[y][x-1] == plant {
		neighbours = append(neighbours, [2]int{x - 1, y})
	}
	if y-1 >= 0 && grid[y-1][x] == plant {
		neighbours = append(neighbours, [2]int{x, y - 1})
	}
	if x+1 < xLim && grid[y][x+1] == plant {
		neighbours = append(neighbours, [2]int{x + 1, y})
	}
	if y+1 < yLim && grid[y+1][x] == plant {
		neighbours = append(neighbours, [2]int{x, y + 1})
	}

	return neighbours
}

func regionDimensions(x, y, xLim, yLim int, grid [][]rune, visited map[[2]int]struct{}) (int, int) {
	toVisit := make([][2]int, 0)
	toVisit = append(toVisit, [2]int{x, y})
	region := make(map[[2]int]struct{})
	region[[2]int{x, y}] = struct{}{}
	size := 0
	perimeter := 0

	for i := 0; i < len(toVisit); i++ {
		size++
		next := toVisit[i]
		visited[next] = struct{}{}

		neighbours := getNeighbours(next[0], next[1], xLim, yLim, grid)

		perimeter += 4 - len(neighbours)

		for _, n := range neighbours {
			if _, ok := region[n]; ok {
				continue
			}
			toVisit = append(toVisit, n)
			region[n] = struct{}{}
		}
	}

	return size, perimeter
}
