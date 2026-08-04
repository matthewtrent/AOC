package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines, err := readFile()
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

func readFile() ([]string, error) {
	fileName := "input.txt"

	if len(os.Args) > 1 {
		fileName = os.Args[1]
	}

	fmt.Println("File Name: ", fileName)

	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	reader := bufio.NewReader(file)

	var lines []string

	for {
		data, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		data = data[:len(data)-1]
		lines = append(lines, data)
	}

	return lines, nil
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

	fmt.Println("Part1 Solution: ", totalSum)
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
	fmt.Println("Part2 Solution: ", totalSum)
}
