package internal

import (
	"fmt"
	"strconv"
	"strings"
)

func Day7(input_file_path string) {
	lines := ReadFileLines(input_file_path)

	targets, numLines := buildDay6Inputs(lines)

	fmt.Printf("Part1: %d\n", getTotal(targets, numLines, false))
	fmt.Printf("Part2: %d\n", getTotal(targets, numLines, true))
}

func getTotal(targets []int, numLines [][]int, part2 bool) int {
	total := 0

	for idx := range targets {
		if isValidLine(targets[idx], numLines[idx], part2) {
			total += targets[idx]
		}
	}

	return total
}

func isValidLine(target int, nums []int, part2 bool) bool {
	tots := []int{nums[0]}

	for _, num := range nums[1:] {
		news := make([]int, 0, 2*len(tots))
		for _, tot := range tots {
			add := tot + num
			mult := tot * num
			conc := 0
			if part2 {
				conc = concat(tot, num)
			}
			if add == target || mult == target || (part2 && conc == target) {
				return true
			} else {
				if add < target {
					news = append(news, add)
				}
				if mult < target {
					news = append(news, mult)
				}
				if part2 && conc < target {
					news = append(news, conc)
				}
			}
		}
		tots = news
	}

	return false
}

func buildDay6Inputs(input []string) ([]int, [][]int) {
	targets := make([]int, len(input))
	numLines := make([][]int, 0, len(input))
	for idx, line := range input {
		parts := strings.Split(line, ": ")
		target, _ := strconv.Atoi(parts[0])
		targets[idx] = target

		nums := strings.Split(parts[1], " ")
		numLine := []int{}
		for _, num := range nums {
			numInt, _ := strconv.Atoi(num)
			numLine = append(numLine, numInt)
		}
		numLines = append(numLines, numLine)
	}

	return targets, numLines
}
