package main

import (
	"fmt"
	"strconv"
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
	totalSum := 0
	for line := range strings.SplitSeq(lines[0], ",") {
		totalSum = totalSum + DoubleNum(line)
	}
	fmt.Println("Part 1:", totalSum)
}

func DoubleNum(line string) int {
	numSum := 0

	ranges := strings.Split(line, "-")

	startInt, _ := strconv.Atoi(ranges[0])
	endInt, _ := strconv.Atoi(ranges[1])

	for startInt <= endInt {

		stringInt := strconv.Itoa(startInt)
		middleVal := len(stringInt) / 2

		if stringInt[:middleVal] == stringInt[middleVal:] {
			numSum += startInt
		}

		startInt++
	}

	return numSum
}

func partTwo(lines []string) {
	totalSum := 0
	for line := range strings.SplitSeq(lines[0], ",") {
		totalSum = totalSum + TripleNum(line)
	}
	fmt.Println("Part 2:", totalSum)
}

func TripleNum(data string) int {
	numSum := 0

	ranges := strings.Split(data, "-")

	startInt, _ := strconv.Atoi(ranges[0])
	endInt, _ := strconv.Atoi(ranges[1])

	for startInt <= endInt {
		currNum := strconv.Itoa(startInt)

		// break the number into sections,
		maxLength := (len(currNum) / 2)

		for i := 1; i <= maxLength; i++ {
			foundPattern := false
			// Check if length % i == 0
			//	if not then it cant seperate evenly and no reason to check sections
			if len(currNum)%i != 0 {
				continue
			}

			potentialSequence := currNum[:i]

			// j is the number of potential matching blocks
			for j := 1; j <= (len(currNum)-1)/i; j++ {
				if potentialSequence != currNum[i*j:i*(j+1)] {
					break
				}

				if j == ((len(currNum) - 1) / i) {
					numSum += startInt
					foundPattern = true
					break
				}
			}

			if foundPattern {
				break
			}

		}

		startInt++
	}

	return numSum
}
