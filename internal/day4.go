package internal

import (
	"fmt"
)

func Day4(input_file_path string) {
	input := ReadCharGrid(input_file_path)

	fmt.Printf("Part1: %d\n", day4part1(input))
	fmt.Printf("Part2: %d\n", day4part2(input))
}

func day4part1(input [][]rune) int {
	total := 0
	xLim := len(input[0])
	yLim := len(input)
	for idx, row := range input {
		for jdx, char := range row {
			if char == 'X' {
				total += checkX(input, jdx, idx, xLim, yLim)
			} else if char == 'S' {
				total += checkS(input, jdx, idx, xLim, yLim)
			}
		}
	}

	return total
}

func day4part2(input [][]rune) int {
	total := 0
	xLim := len(input[0])
	yLim := len(input)
	for idx, row := range input {
		for jdx := range row {
			if checkMas(input, jdx, idx, xLim, yLim) {
				total++
			}
		}
	}

	return total
}

func checkX(input [][]rune, x, y, xLim, yLim int) int {
	matches := 0

	if xLim > x+3 && input[y][x+1] == 'M' && input[y][x+2] == 'A' && input[y][x+3] == 'S' {
		matches++
	}
	if yLim > y+3 && input[y+1][x] == 'M' && input[y+2][x] == 'A' && input[y+3][x] == 'S' {
		matches++
	}
	if xLim > x+3 && yLim > y+3 && input[y+1][x+1] == 'M' && input[y+2][x+2] == 'A' && input[y+3][x+3] == 'S' {
		matches++
	}
	if x-3 >= 0 && yLim > y+3 && input[y+1][x-1] == 'M' && input[y+2][x-2] == 'A' && input[y+3][x-3] == 'S' {
		matches++
	}

	return matches
}

func checkS(input [][]rune, x, y, xLim, yLim int) int {
	matches := 0

	if xLim > x+3 && input[y][x+1] == 'A' && input[y][x+2] == 'M' && input[y][x+3] == 'X' {
		matches++
	}
	if yLim > y+3 && input[y+1][x] == 'A' && input[y+2][x] == 'M' && input[y+3][x] == 'X' {
		matches++
	}
	if xLim > x+3 && yLim > y+3 && input[y+1][x+1] == 'A' && input[y+2][x+2] == 'M' && input[y+3][x+3] == 'X' {
		matches++
	}
	if x-3 >= 0 && yLim > y+3 && input[y+1][x-1] == 'A' && input[y+2][x-2] == 'M' && input[y+3][x-3] == 'X' {
		matches++
	}

	return matches
}

func checkMas(input [][]rune, x, y, xLim, yLim int) bool {
	if !(xLim > x+2 && yLim > y+2) || (input[y+1][x+1] != 'A') {
		return false
	}

	corners := []rune{input[y][x], input[y][x+2], input[y+2][x], input[y+2][x+2]}

	if corners[0] == corners[3] || corners[1] == corners[2] {
		return false
	}

	ms := 0
	ss := 0

	for _, corner := range corners {
		if corner == 'M' {
			ms++
		} else if corner == 'S' {
			ss++
		}
	}

	if ms == 2 && ss == 2 {
		return true
	} else {
		return false
	}
}
