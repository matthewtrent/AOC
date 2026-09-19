package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/matthewtrent/aoc/utils"
)

var seenSplit map[string]int

func main() {
	lines, err := utils.ReadFile()
	if err != nil {
		fmt.Println(err)
		return
	}

	// for _, line := range lines {
	// 	fmt.Println(line)
	// }

	seenSplit = make(map[string]int)

	partOne(lines)
	partTwo(lines)
}

func partOne(lines []string) {
	data := make([][]string, len(lines))
	totalSplits := 0

	// Populate Data
	for i, line := range lines {
		data[i] = strings.Split(line, "")
	}

	for i := 1; i < len(data); i++ {
		for j, char := range data[i] {
			// Propagate beam down if needed
			if data[i-1][j] == "l" || data[i-1][j] == "S" {

				// Checking if currently on a splitter
				if char == "^" {
					totalSplits++
					data[i][j-1] = "l"
					data[i][j+1] = "l"
					continue
				}
				data[i][j] = "l"
			}
		}
	}

	// for i := range data {
	// 	for _, char := range data[i] {
	// 		fmt.Print(char, " ")
	// 	}
	// 	fmt.Println()
	// }

	fmt.Println("Part One: ", totalSplits)
}

func partTwo(lines []string) {
	data := make([][]string, len(lines))
	totalTimelines := 0

	// Populate Data
	for i, line := range lines {
		data[i] = strings.Split(line, "")
	}

	// Finding where the S is
	startX := 0
	for i, char := range data[0] {
		if char == "S" {
			startX = i
		}
	}

	totalTimelines = crawlTimeLines(data, startX, 1)
	fmt.Println("Part Two: ", totalTimelines)
}

func crawlTimeLines(data [][]string, x int, y int) int {
	// Move down till encountering a "^" or the bottom
	for i := y; i < len(data); i++ {
		if data[i][x] == "^" {
			if seenSplit[strconv.Itoa(i)+strconv.Itoa(x)] == 0 {
				seenSplit[strconv.Itoa(i)+strconv.Itoa(x)] = crawlTimeLines(data, x-1, i) + crawlTimeLines(data, x+1, i)
			}
			return seenSplit[strconv.Itoa(i)+strconv.Itoa(x)]
		}
	}
	return 1
}
