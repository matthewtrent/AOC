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
		fmt.Println(err)
		return
	}

	// for _, line := range lines {
	// 	fmt.Println(line)
	// }

	partOne(lines)
	partTwo(lines)
}

func partOne(lines []string) {
	totalSum := 0

	data := make([][]string, len(lines))

	for i, line := range lines {
		for temp := range strings.SplitSeq(line, " ") {
			if temp != "" {
				data[i] = append(data[i], temp)
			}
		}
	}
	for i := range data[0] {
		op := data[len(data)-1][i]
		partialTotal := 0
		for j := 0; j < len(data)-1; j++ {
			// fmt.Println(data[j][i])
			num, _ := strconv.Atoi(data[j][i])
			if op == "+" {
				partialTotal += num
			}

			if op == "*" {
				if partialTotal == 0 {
					partialTotal = 1
				}
				partialTotal *= num
			}
		}
		totalSum += partialTotal
	}

	fmt.Println("Part 1: ", totalSum)
}

func partTwo(lines []string) {
	var ops []string
	opNum := 0
	totalSum := 0
	partialTotal := 0

	// Creating a simple array for all the operators
	for temp := range strings.SplitSeq(lines[len(lines)-1], " ") {
		if temp != "" {
			ops = append(ops, temp)
		}
	}

	currOp := ops[opNum]

	for i := range lines[0] {
		currNum := ""

		for j := 0; j < len(lines)-1; j++ {
			currNum += string(lines[j][i])
		}

		trimmedNum := strings.TrimSpace(currNum)

		// Finished current set of numbers, switch to next op and continue
		if trimmedNum == "" {
			totalSum += partialTotal
			// fmt.Println(" = ", partialTotal)
			partialTotal = 0
			opNum++
			currOp = ops[opNum]
			continue
		}

		num, _ := strconv.Atoi(trimmedNum)
		if currOp == "+" {
			// fmt.Print(num, " + ")
			partialTotal += num
		}

		if currOp == "*" {
			if partialTotal == 0 {
				partialTotal = 1
			}
			partialTotal *= num
			// fmt.Print(num, " * ")
		}
	}
	totalSum += partialTotal
	fmt.Println("Part 2: ", totalSum)
}
