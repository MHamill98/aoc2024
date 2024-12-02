package internal

import (
	"fmt"
	"strconv"
	"strings"
)

func Day2(input_file_path string) {
	lines := ReadFileLines(input_file_path)
	reports := buildReports(lines)

	fmt.Printf("Part1: %d\n", day2part1(reports))
	fmt.Printf("Part2: %d\n", day2part2(reports))
}

func buildReports(lines []string) [][]int {
	length := len(lines)
	reports := make([][]int, length)

	for idx, val := range lines {
		nums := strings.Fields(val)
		reports[idx] = make([]int, len(nums))
		for idx2, val2 := range nums {
			num, _ := strconv.Atoi(val2)
			reports[idx][idx2] = num
		}
	}

	return reports
}

func day2part1(reports [][]int) int {
	count := 0

	for _, report := range reports {
		if isSafe(report) {
			count++
		}
	}

	return count
}

func day2part2(reports [][]int) int {
	count := 0

	for _, report := range reports {
		if safe, idx := isSafePart2(report, -1); safe {
			count++
		} else {
			for i := idx - 1; i <= idx+1; i++ {
				if safe, _ := isSafePart2(report, i); safe {
					count++
					break
				}
			}
		}
	}

	return count
}

func isSafe(report []int) bool {
	ascending := report[1]-report[0] > 0

	for idx := range report {
		if idx == 0 {
			continue
		}
		if !isSafePair(report[idx-1], report[idx], ascending) {
			return false
		}
	}

	return true
}

func isSafePart2(report []int, nodeToSkip int) (bool, int) {
	first, second := 0, 1

	if first == nodeToSkip {
		first++
		second++
	} else if second == nodeToSkip {
		second++
	}

	ascending := report[second]-report[first] > 0
	for second < len(report) {
		if !isSafePair(report[first], report[second], ascending) {
			return false, first
		}

		first = second
		second++
		if second == nodeToSkip {
			second++
		}
	}

	return true, first
}

func isSafePair(a, b int, ascending bool) bool {
	diff := b - a
	if (ascending && diff < 0) || (!ascending && diff > 0) || abs(diff) > 3 || abs(diff) < 1 {
		return false
	}
	return true
}
