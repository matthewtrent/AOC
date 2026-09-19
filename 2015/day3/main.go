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

	partOne(lines)
	partTwo(lines)
}

func partTwo(lines []string) {
	for _, line := range lines {
		set := utils.NewSet[utils.Point]()
		santaLocation := utils.Point{X: 0, Y: 0}
		robotLocation := utils.Point{X: 0, Y: 0}
		for i, char := range line {
			set.Add(santaLocation)
			set.Add(robotLocation)

			movementPoint := utils.Point{}

			switch char {
			case '^':
				movementPoint = utils.Up
			case 'v':
				movementPoint = utils.Down
			case '>':
				movementPoint = utils.Right
			case '<':
				movementPoint = utils.Left
			default:
			}

			if i%2 == 0 {
				santaLocation = santaLocation.Add(movementPoint)
			} else {
				robotLocation = robotLocation.Add(movementPoint)
			}
		}

		fmt.Println("Part 2: ", len(set))
	}
}

func partOne(lines []string) {
	for _, line := range lines {
		set := utils.NewSet[utils.Point]()
		currLocation := utils.Point{X: 0, Y: 0}
		for _, char := range line {
			set.Add(currLocation)
			switch char {
			case '^':
				currLocation = currLocation.Add(utils.Up)
			case 'v':
				currLocation = currLocation.Add(utils.Down)
			case '>':
				currLocation = currLocation.Add(utils.Right)
			case '<':
				currLocation = currLocation.Add(utils.Left)
			default:
			}
		}

		fmt.Println("Part 1: ", len(set))
	}
}
