package internal

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func getPageAndRule(line string) (int, int) {
	var page, rule int
	fmt.Sscanf(line, "%d|%d", &page, &rule)
	return page, rule
}

func getPages(line string) []int {
	var pages []int
	for _, page := range strings.Split(line, ",") {
		p, _ := strconv.Atoi(page)
		pages = append(pages, p)
	}

	return pages
}

func buildInputs(input []string) (map[int][]int, [][]int) {
	rules := make(map[int][]int)
	lines := [][]int{}

	idx := 0

	for input[idx] != "" {
		page, rule := getPageAndRule(input[idx])
		rules[page] = append(rules[page], rule)
		idx++
	}

	idx++

	for idx < len(input) {
		lines = append(lines, getPages(input[idx]))
		idx++
	}

	return rules, lines
}

func getMiddleValue(pages []int) int {
	return pages[len(pages)/2]
}

func isValid(pages []int, rules map[int][]int) bool {
	seen := make(map[int]struct{})
	for _, page := range pages {
		for _, p := range rules[page] {
			if _, ok := seen[p]; ok {
				return false
			}
		}
		seen[page] = struct{}{}
	}
	return true
}

func day5part1(rules map[int][]int, lines [][]int) int {
	total := 0

	for _, pages := range lines {
		if isValid(pages, rules) {
			total += getMiddleValue(pages)
		}
	}

	return total
}

func day5part2(rules map[int][]int, lines [][]int) int {
	total := 0

	for _, pages := range lines {
		if !isValid(pages, rules) {
			slices.SortFunc(pages, func(a, b int) int {
				if slices.Contains(rules[a], b) {
					return -1
				} else if slices.Contains(rules[b], a) {
					return 1
				}
				return 0
			})

			total += getMiddleValue(pages)
		}
	}

	return total
}

func Day5(input_file_path string) {
	input := ReadFileLines(input_file_path)
	rules, pages := buildInputs(input)

	fmt.Printf("Part1: %d\n", day5part1(rules, pages))
	fmt.Printf("Part2: %d\n", day5part2(rules, pages))
}
