package internal

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

const blinks = 45

func Day11(input_file_path string) {
	input := ReadFile(input_file_path)
	stones := buildStones(input)

	fmt.Printf("Part1: %d\n", day11part2(stones, blinks))
}

func day11part1(stones []int) int {
	blinks := 0
	numStones := len(stones)
	splits := make(map[int][2]int)
	for blinks < 45 {
		newStones := make([]int, 0, 2*numStones)
		numStones = 0
		for _, stone := range stones {
			if stone == 0 {
				newStones = append(newStones, 1)
			} else if numDigs(stone)%2 == 0 {
				if split, ok := splits[stone]; ok {
					newStones = append(newStones, split[0], split[1])
				} else {
					split := splitStone(stone)
					newStones = append(newStones, split[0], split[1])
					splits[stone] = split
				}
			} else {
				newStones = append(newStones, stone*2024)
			}
		}
		stones = newStones
		blinks++
	}

	return len(stones)
}

func day11part2(stones []int, numBlinks int) int {
	blinks := 0
	numStones := len(stones)
	splits := make(map[int][2]int)
	for blinks < numBlinks {
		newStones := make([]int, 0, 2*numStones)
		numStones = 0
		for _, stone := range stones {
			if stone == 0 {
				newStones = append(newStones, 1)
			} else if numDigs(stone)%2 == 0 {
				if split, ok := splits[stone]; ok {
					newStones = append(newStones, split[0], split[1])
				} else {
					split := splitStone(stone)
					newStones = append(newStones, split[0], split[1])
					splits[stone] = split
				}
			} else {
				newStones = append(newStones, stone*2024)
			}
		}
		stones = newStones
		blinks++
	}

	return len(stones)
}

func buildStones(stones string) []int {
	stoneInts := make([]int, 0)
	for _, s := range strings.Fields(stones) {
		stone, _ := strconv.Atoi(s)
		stoneInts = append(stoneInts, stone)
	}

	return stoneInts
}

func splitStone(stone int) [2]int {
	digs := numDigs(stone)
	left := stone / int(math.Pow(10, float64(digs/2)))
	right := stone % int(math.Pow(10, float64(digs/2)))

	return [2]int{left, right}
}
