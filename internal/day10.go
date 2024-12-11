package internal

import "fmt"

func Day10(input_file_path string) {
	cells := ReadNumGrid(input_file_path)

	fmt.Printf("Part1: %d\n", day10part1(cells))
	fmt.Printf("Part2: %d\n", day10part2(cells))
}

func day10part1(cells [][]int) int {
	total := 0

	for y, row := range cells {
		for x, cell := range row {
			if cell == 0 {
				val := 0
				validNeighbours := getValidNeighbours(cells, [2]int{x, y})
				reachable := make(map[[2]int]struct{})

				for val < 8 {
					newNeighbours := make([][2]int, 0, 4)
					for _, n := range validNeighbours {
						newNeighbours = append(newNeighbours, getValidNeighbours(cells, n)...)
					}

					validNeighbours = newNeighbours
					val++
				}
				for _, n := range validNeighbours {
					reachable[n] = struct{}{}
				}
				total += len(reachable)
			}
		}
	}

	return total
}

func day10part2(cells [][]int) int {
	total := 0

	for y, row := range cells {
		for x, cell := range row {
			if cell == 0 {
				val := 0
				validNeighbours := getValidNeighbours(cells, [2]int{x, y})

				for val < 8 {
					newNeighbours := make([][2]int, 0, 4)
					for _, n := range validNeighbours {
						newNeighbours = append(newNeighbours, getValidNeighbours(cells, n)...)
					}

					validNeighbours = newNeighbours
					val++
				}
				total += len(validNeighbours)
			}
		}
	}

	return total
}

func getValidNeighbours(cells [][]int, cell [2]int) [][2]int {
	x := cell[0]
	y := cell[1]
	val := cells[y][x]
	neighbours := make([][2]int, 0, 4)

	if x-1 >= 0 && cells[y][x-1] == val+1 {
		neighbours = append(neighbours, [2]int{x - 1, y})
	}
	if x+1 < len(cells[0]) && cells[y][x+1] == val+1 {
		neighbours = append(neighbours, [2]int{x + 1, y})
	}
	if y-1 >= 0 && cells[y-1][x] == val+1 {
		neighbours = append(neighbours, [2]int{x, y - 1})
	}
	if y+1 < len(cells) && cells[y+1][x] == val+1 {
		neighbours = append(neighbours, [2]int{x, y + 1})
	}

	return neighbours
}
