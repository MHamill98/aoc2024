package aoc2024lib

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func Day1(input_file_path string) {
	list1, list2 := buildLists(input_file_path)
	fmt.Println("Part 1: ", part1(list1, list2))
	fmt.Println("Part 2: ", part2(list1, list2))
}

func buildLists(input_file_path string) ([]int, []int) {
	lines := ReadFileLines(input_file_path)

	length := len(lines)

	list1 := make([]int, length)
	list2 := make([]int, length)

	for idx, val := range lines {
		nums := strings.Fields(val)
		num1, _ := strconv.Atoi(nums[0])
		num2, _ := strconv.Atoi(nums[1])
		list1[idx] = num1
		list2[idx] = num2
	}

	return list1, list2
}

func part1(list1, list2 []int) int {
	total := 0

	sort.Ints(list1)
	sort.Ints(list2)

	for idx, val := range list1 {
		diff := val - list2[idx]
		if diff > 0 {
			total += diff
		} else {
			total += -diff
		}
	}

	return total
}

func part2(list1, list2 []int) int {
	counts := make(map[int]int)
	total := 0

	for _, val := range list2 {
		counts[val] += 1
	}

	for _, val := range list1 {
		total += val * counts[val]
	}

	return total
}
