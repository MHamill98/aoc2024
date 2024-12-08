package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/MHamill98/aoc2024/internal"
)

func main() {
	day := flag.Int("day", 1, "Day to run")
	example := flag.Bool("example", false, "Use example data.")
	flag.Parse()

	input_file_path := fmt.Sprintf("input/day%d.txt", *day)
	if *example {
		input_file_path = fmt.Sprintf("input/day%d-example.txt", *day)
	}

	switch *day {
	case 1:
		internal.Day1(input_file_path)
	case 2:
		internal.Day2(input_file_path)
	case 3:
		internal.Day3(input_file_path)
	case 4:
		internal.Day4(input_file_path)
	case 5:
		internal.Day5(input_file_path)
	case 6:
		internal.Day6(input_file_path)
	case 7:
		internal.Day7(input_file_path)
	default:
		log.Fatalf("Day %d not implemented", *day)
	}
}
