package internal

import (
	"fmt"
	"maps"
)

const (
	xLim, yLim = 101, 103
)

type Robot struct {
	xPos, yPos, xVel, yVel int
}

func (r *Robot) move(times int) {
	newX := (r.xPos + times*r.xVel) % xLim
	newY := (r.yPos + times*r.yVel) % yLim

	if newX < 0 {
		newX += xLim
	}
	if newY < 0 {
		newY += yLim
	}

	r.xPos, r.yPos = newX, newY
}

func Day14(input_file_path string) {
	input := ReadFileLines(input_file_path)
	robots := parseInput(input)

	fmt.Printf("Part1: %d\n", day14part1(robots))
	// For part 2, ran 10000 iterations, output resulting grids to file and grepped the output for long strings of 1s.
}

func day14part1(robots []Robot) int {
	robLocs := make(map[[2]int]int)

	for _, r := range robots {
		r.move(100)
		robLocs[[2]int{r.xPos, r.yPos}]++
	}

	return multQuads(robLocs)
}

func parseInput(input []string) []Robot {
	robots := make([]Robot, 0)
	for _, line := range input {
		var robot Robot
		fmt.Sscanf(line, "p=%d,%d v=%d,%d", &robot.xPos, &robot.yPos, &robot.xVel, &robot.yVel)
		robots = append(robots, robot)
	}
	return robots
}

func multQuads(robLocs map[[2]int]int) int {
	quads := make(map[int]int)
	for r := range maps.Keys(robLocs) {
		quad := quadrant(r)
		if quad > 0 && robLocs[r] > 0 {
			quads[quadrant(r)] += robLocs[r]
		}
	}
	tot := 1
	for q := range maps.Keys(quads) {
		tot *= quads[q]
	}
	return tot
}

func quadrant(r [2]int) int {
	mid := [2]int{xLim / 2, yLim / 2}
	if r[0] < mid[0] && r[1] < mid[1] {
		return 1
	}
	if r[0] > mid[0] && r[1] < mid[1] {
		return 2
	}
	if r[0] < mid[0] && r[1] > mid[1] {
		return 3
	}
	if r[0] > mid[0] && r[1] > mid[1] {
		return 4
	}
	return -1
}
