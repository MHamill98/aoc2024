package internal

import (
	"fmt"
	"strconv"
)

func Day22(input_file_path string) {
	input := ReadFileLines(input_file_path)

	fmt.Println("Part1:", day22part1(input))
	fmt.Println("Part2:", day22part2(input))
}

func day22part1(input []string) int {
	tot := 0

	for _, line := range input {
		lineint, _ := strconv.Atoi(line)
		for i := 0; i < 2000; i++ {
			lineint = processLine(lineint)
		}
		tot += lineint
	}

	return tot
}

func day22part2(input []string) int {
	return 0
}

func processLine(line int) int {
	mixin := line * 64
	line ^= mixin
	line %= 16777216

	mixin = line / 32
	line ^= mixin
	line %= 16777216

	mixin = line * 2048
	line ^= mixin
	line %= 16777216

	return line
}
