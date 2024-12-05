package internal

import (
	"log"
	"os"
	"strings"
)

func ReadFile(input_file_path string) string {
	content, err := os.ReadFile(input_file_path)
	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}

func ReadFileLines(input_file_path string) []string {
	content := ReadFile(input_file_path)
	return strings.Split(content, "\n")
}

func ReadCharGrid(input_file_path string) [][]rune {
	// Take the input file path and read the file into a grid of characters.
	lines := ReadFileLines(input_file_path)
	grid := make([][]rune, len(lines))

	for i, line := range lines {
		for _, char := range line {
			grid[i] = append(grid[i], char)
		}
	}

	return grid
}
