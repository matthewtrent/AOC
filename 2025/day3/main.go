package main

import (
	"fmt"
	"strconv"

	"github.com/matthewtrent/aoc/utils"
)

func main() {
	lines, err := utils.ReadFile()
	if err != nil {
		print(err)
		return
	}

	partOne(lines)
	partTwo(lines)
}

func partOne(lines []string) {
	total := 0
	for _, line := range lines {
		total += largestBattery(line, 2)
	}
	fmt.Println("Part 1: ", total)
}

func partTwo(lines []string) {
	total := 0
	for _, line := range lines {
		total += largestBattery(line, 12)
	}
	fmt.Println("Part 2: ", total)
}

func largestBattery(data string, length int) int {
	largest := make([]rune, length)

	for i, curRune := range data {
		runesLeft := len(data) - i
		indexToCheck := 0

		if runesLeft < length {
			indexToCheck = i - (len(data) - length)
		}

		replaced := false
		for j := indexToCheck; j < length; j++ {
			if replaced {
				largest[j] = '0'
			} else if curRune > largest[j] {
				largest[j] = curRune
				replaced = true
			}
		}
	}

	val, _ := strconv.Atoi(string(largest))

	return val
}
