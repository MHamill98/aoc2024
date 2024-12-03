package internal

import (
	"fmt"
	"regexp"
	"strconv"
)

func Day3(input_file_path string) {
	input := ReadFile(input_file_path)

	fmt.Printf("Part1: %d\n", day3part1(input))
	fmt.Printf("Part2: %d\n", day3part2(input))
}

func day3part1(input string) int {
	r, _ := regexp.Compile(`mul\(([0-9]{1,3}),([0-9]{1,3})\)`)

	vals := r.FindAllStringSubmatch(input, -1)
	total := 0

	for _, matches := range vals {
		num1, _ := strconv.Atoi(matches[1])
		num2, _ := strconv.Atoi(matches[2])
		total += num1 * num2
	}

	return total
}

func day3part2(input string) int {
	r, _ := regexp.Compile(`mul\(([0-9]{1,3}),([0-9]{1,3})\)|(do\(\))|(don't\(\))`)

	vals := r.FindAllStringSubmatch(input, -1)
	total := 0
	do := true

	for _, matches := range vals {
		if matches[0] == "do()" {
			do = true
			continue
		} else if matches[0] == "don't()" {
			do = false
			continue
		}

		if do {
			num1, _ := strconv.Atoi(matches[1])
			num2, _ := strconv.Atoi(matches[2])
			total += num1 * num2
		}
	}

	return total
}
