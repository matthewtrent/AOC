package main

import (
	"fmt"
	"strings"

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
	wall := createWall(lines)

	foundIter := 0
	foundIter, _ = openRolls(wall, 4)

	total += foundIter

	fmt.Println("Part 1: ", total)
}

func partTwo(lines []string) {
	total := 0
	wall := createWall(lines)

	for {
		foundIter := 0
		foundIter, wall = openRolls(wall, 4)
		if foundIter == 0 {
			break
		}

		total += foundIter
	}

	fmt.Println("Part 2: ", total)
}

func createWall(lines []string) []string {
	var wall []string

	wall = append(wall, strings.Repeat(".", len(lines[0])+2))
	for _, line := range lines {

		data := "." + line + "."
		wall = append(wall, data)
	}

	wall = append(wall, strings.Repeat(".", len(lines[0])+2))
	return wall
}

func openRolls(wall []string, closedRolls int) (int, []string) {
	total := 0
	var takenRoll []utils.Point
	// Go through each row of the wall
	for i := 1; i < len(wall)-1; i++ {
		// fmt.Println(wall[i-1])
		// fmt.Println(wall[i])
		// fmt.Println(wall[i+1])
		for j := 1; j < len(wall[i])-1; j++ {

			//fmt.Printf("currRune: %c\n", wall[i][j])

			subTotal := 0
			if wall[i][j] == '@' {
				if wall[i-1][j-1] == '@' {
					subTotal += 1
				}
				if wall[i][j-1] == '@' {
					subTotal += 1
				}
				if wall[i+1][j-1] == '@' {
					subTotal += 1
				}
				if wall[i-1][j] == '@' {
					subTotal += 1
				}
				if wall[i+1][j] == '@' {
					subTotal += 1
				}
				if wall[i-1][j+1] == '@' {
					subTotal += 1
				}
				if wall[i][j+1] == '@' {
					subTotal += 1
				}
				if wall[i+1][j+1] == '@' {
					subTotal += 1
				}
				if subTotal < closedRolls {
					//fmt.Println("Less than 4 @ found")
					total += 1
					takenRoll = append(takenRoll, utils.Point{
						X: i,
						Y: j,
					})
				}
			}

		}
	}

	for _, index := range takenRoll {
		line := []rune(wall[index.X])
		line[index.Y] = '.'
		wall[index.X] = string(line)
	}

	return total, wall
}
