package main

import (
	"fmt"
	"strings"

	"github.com/matthewtrent/aoc/utils"
)

type LightStatus int

const (
	Off LightStatus = iota
	On
	Toggle
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
	totalLight := 0
	// Create 1000x1000 boolean table
	// roof[row][col]
	table := make([][]bool, 1000)
	for i := range table {
		table[i] = make([]bool, 1000)
	}

	// Parse input
	for _, line := range lines {
		// Check first word
		words := strings.Split(line, " ")
		switch words[0] {
		case "turn":
			pointOne := utils.ParsePointFromString(words[2])
			pointTwo := utils.ParsePointFromString(words[4])
			switch words[1] {
			case "on":
				totalLight += updateLights(table, On, pointOne, pointTwo)
			case "off":
				totalLight += updateLights(table, Off, pointOne, pointTwo)
			}
		case "toggle":
			pointOne := utils.ParsePointFromString(words[1])
			pointTwo := utils.ParsePointFromString(words[3])
			totalLight += updateLights(table, Toggle, pointOne, pointTwo)
		}
	}

	fmt.Println("Part 1: ", totalLight)
}

func updateLights(lights [][]bool, status LightStatus, pointOne utils.Point, pointTwo utils.Point) int {
	lightUpdate := 0

	for i := pointOne.X; i <= pointTwo.X; i++ {
		for j := pointOne.Y; j <= pointTwo.Y; j++ {
			deltaLight := 0
			switch status {
			case On:
				lights[i][j], deltaLight = switchOn(lights[i][j])
			case Off:
				lights[i][j], deltaLight = switchOff(lights[i][j])

			case Toggle:
				lights[i][j], deltaLight = ToggleLight(lights[i][j])

			}
			lightUpdate += deltaLight
		}
	}
	return lightUpdate
}

func ToggleLight(b bool) (bool, int) {
	if b {
		return false, -1
	} else {
		return true, 1
	}
}

func switchOff(b bool) (bool, int) {
	delta := 0
	if b {
		delta = -1
	}
	return false, delta
}

func switchOn(b bool) (bool, int) {
	delta := 0
	if !b {
		delta = 1
	}
	return true, delta
}

func partTwo(lines []string) {
	totalLight := 0
	// Create 1000x1000 boolean table
	// roof[row][col]
	table := make([][]int, 1000)
	for i := range table {
		table[i] = make([]int, 1000)
	}

	// Parse input
	for _, line := range lines {
		// Check first word
		words := strings.Split(line, " ")
		switch words[0] {
		case "turn":
			pointOne := utils.ParsePointFromString(words[2])
			pointTwo := utils.ParsePointFromString(words[4])
			switch words[1] {
			case "on":
				totalLight += updateLightsNew(table, On, pointOne, pointTwo)
			case "off":
				totalLight += updateLightsNew(table, Off, pointOne, pointTwo)
			}
		case "toggle":
			pointOne := utils.ParsePointFromString(words[1])
			pointTwo := utils.ParsePointFromString(words[3])
			totalLight += updateLightsNew(table, Toggle, pointOne, pointTwo)
		}
	}

	fmt.Println("Part 2: ", totalLight)
}

func updateLightsNew(lights [][]int, status LightStatus, pointOne, pointTwo utils.Point) int {
	lightUpdate := 0

	for i := pointOne.X; i <= pointTwo.X; i++ {
		for j := pointOne.Y; j <= pointTwo.Y; j++ {
			switch status {
			case On:
				lights[i][j]++
				lightUpdate++
			case Off:
				if lights[i][j] > 0 {
					lights[i][j]--
					lightUpdate--
				}

			case Toggle:
				lights[i][j] += 2
				lightUpdate += 2

			}
		}
	}
	return lightUpdate
}
