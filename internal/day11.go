package internal

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Day11(input_file_path string) {
	input := ReadFile(input_file_path)
	stones := buildStones(input)

	fmt.Printf("Part1: %d\n", day11part1(stones))
}

func day11part1(stones []int) int {
	blinks := 0
	numStones := len(stones)
	for blinks < 25 {
		newStones := make([]int, 0, 2*numStones)
		numStones = 0
		for _, stone := range stones {
			if stone == 0 {
				newStones = append(newStones, 1)
			} else if numDigs(stone)%2 == 0 {
				newStones = append(newStones, splitStone(stone)...)
			} else {
				newStones = append(newStones, stone*2024)
			}
		}
		stones = newStones
		blinks++
	}

	return len(stones)
}

func day11part2(stones []int) int {
	return 0
}

func buildStones(stones string) []int {
	stoneInts := make([]int, 0)
	for _, s := range strings.Fields(stones) {
		stone, _ := strconv.Atoi(s)
		stoneInts = append(stoneInts, stone)
	}

	return stoneInts
}

func splitStone(stone int) []int {
	digs := numDigs(stone)
	left := stone / int(math.Pow(10, float64(digs/2)))
	right := stone % int(math.Pow(10, float64(digs/2)))

	return []int{left, right}
}
