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
	currNum := 50
	numZero := 0

	for _, line := range lines {

		currNum, _ = rotateLock(line, currNum)
		if currNum == 0 {
			numZero += 1
		}
	}

	fmt.Println("Part 1: ", numZero)
}

func partTwo(lines []string) {
	currNum := 50
	numZero := 0

	for _, line := range lines {
		gotToZero := 0

		currNum, gotToZero = rotateLock(line, currNum)
		numZero += gotToZero
	}

	fmt.Println("Part 2: ", numZero)
}

/*
* simulates rotating a lock dial, numbered 0-99
*	@Args: data - the rotation direction and the amount rotated, example R55
*				currNum - the current dial number
*	@returns:
*				int - the new current dial number
*				int - the amount of times that the dial landed exactly on 0
* */
func rotateLock(data string, currNum int) (int, int) {
	total0 := 0

	rotateDir := 1 // Default rotate left

	if data[0] == 'L' {
		rotateDir = -1
	}

	// Extract the number of times to rotate
	num := data[1:]
	numSpin, _ := strconv.Atoi(num)

	for range numSpin{
		currNum += rotateDir
		if currNum == -1 {
			currNum = 99
		}

		currNum = currNum % 100

		if currNum == 0 {
			total0 += 1
		}
	}

	return currNum, total0
}
