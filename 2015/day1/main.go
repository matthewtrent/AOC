package main

import (
	"fmt"

	"github.com/matthewtrent/aoc/utils"
)

func main() {
	lines, err := utils.ReadFile()
	if err != nil {
		fmt.Println(err)
		return
	}

	sol(lines)
}

func sol(lines []string) {
	for i := range lines {
		floor := 0
		firstBasement := -1
		for j, char := range lines[i] {
			switch char {
			case '(':
				floor++
			case ')':
				floor--
			default:
			}

			if floor == -1 && firstBasement == -1 {
				firstBasement = j + 1
			}
		}
		fmt.Println("Part 1: ", floor)
		fmt.Println("Part 2: ", firstBasement)
	}
}
