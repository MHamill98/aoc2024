package internal

import (
	"fmt"
	"maps"
	"strconv"
	"strings"
)

func Day9(input_file_path string) {
	nums := ReadFile(input_file_path)

	bitList := buildBitList(nums)

	properList, spaces, spaceLengths := buildProperList(bitList)

	fmt.Printf("Part1: %d\n", day9part1(properList))
	fmt.Printf("Part2: %d\n", day9part2(properList, spaces, spaceLengths))
}

func day9part1(input []int) int {
	left := 0
	right := len(input) - 1
	tot := 0

	for left <= right {
		if input[left] != -1 {
			tot += input[left] * left
		} else {
			for input[len(input)-1] == -1 {
				input = input[:len(input)-1]
				right--
			}
			tot += input[len(input)-1] * left
			input = input[:len(input)-1]
			right--
		}
		left++
	}

	return tot
}

func day9part2(input []int, spaces [][2]int, spaceLengths map[int]int) int {
	tot := 0
	n := len(input) - 1
	for input[n] == -1 {
		n--
	}
	idx := input[n] + 1

	for n > 0 {
		idx--
		for n >= 0 && (input[n] == -1 || input[n] != idx) {
			n--
		}
		sliceStart := n
		spaceSize := 0
		for n >= 0 && input[n] == idx {
			spaceSize++
			n--
		}
		if n < 0 || !hasFreeSpace(spaceLengths, spaceSize) {
			continue
		}
		var space [2]int = [2]int{-1, -1}
		var spaceIdx int
		for i, s := range spaces {
			if s[1]-s[0]+1 >= spaceSize {
				space = s
				spaceIdx = i
				break
			}
		}
		start := space[0]
		end := space[1]
		spaceLength := end - start + 1
		if start == -1 || start+spaceSize-1 > n {
			continue
		}
		toFill := spaceSize
		for toFill > 0 {
			input[start] = idx
			start++
			toFill--
		}
		spaces[spaceIdx] = [2]int{start, end}
		spaceLengths[end-start+1]++
		spaceLengths[spaceLength]--
		for sliceStart > n {
			input[sliceStart] = -1
			sliceStart--
		}
	}

	for i, val := range input {
		if val != -1 {
			tot += i * val
		}
	}

	return tot
}

func buildBitList(nums string) []int {
	bitList := make([]int, 0, len(nums))
	for _, num := range strings.Split(nums, "") {
		n, _ := strconv.Atoi(num)
		bitList = append(bitList, n)
	}

	return bitList
}

func buildProperList(input []int) ([]int, [][2]int, map[int]int) {
	n := 0
	idx := 0
	output := make([]int, 0, 9*len(input))
	spaces := make([][2]int, 0)
	spaceLengths := make(map[int]int)
	length := len(input)

	for n < length {
		if n%2 == 0 {
			i := 0
			for i < input[n] {
				output = append(output, idx)
				i++
			}
			idx++
		} else {
			spaceSize := input[n]
			spaceLengths[spaceSize]++
			if spaceSize > 0 {
				start := len(output)
				spaces = append(spaces, [2]int{start, start + input[n] - 1})
				i := 0
				for i < spaceSize {
					output = append(output, -1)
					i++
				}
			}
		}
		n++
	}

	return output, spaces, spaceLengths
}

func hasFreeSpace(spaceLengths map[int]int, needed int) bool {
	for i := range maps.Keys(spaceLengths) {
		if i >= needed && spaceLengths[i] > 0 {
			return true
		}
	}
	return false
}
