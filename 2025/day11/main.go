package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	lines, err := readFile()
	if err != nil {
		fmt.Println(err)
		return
	}

	inOutMap := linesToMap(lines)

	partOne(inOutMap)
	partTwo(inOutMap)
}

func partOne(inOutMap map[string][]string) {
	startString := "you"
	endString := "out"
	numExits := make(map[string]int)
	DFSNumExits(startString, endString, inOutMap, numExits)
	fmt.Println("Part 1: ", numExits["you"])
}

func partTwo(inOutMap map[string][]string) {
	startString := "svr"
	endString := "out"
	totalCrossedExits := DFSNumExits(startString, "fft", inOutMap, make(map[string]int)) *
		DFSNumExits("fft", "dac", inOutMap, make(map[string]int)) *
		DFSNumExits("dac", endString, inOutMap, make(map[string]int))
	fmt.Println("Part 2: ", totalCrossedExits)
}

func DFSNumExits(startString string, endString string, inOutMap map[string][]string, numExits map[string]int) int {
	if startString == endString {
		return 1
	}

	val, ok := numExits[startString]

	if ok {
		return val
	}

	branches := inOutMap[startString]

	totalExits := 0

	for _, branch := range branches {
		totalExits += DFSNumExits(branch, endString, inOutMap, numExits)
	}

	numExits[startString] = totalExits

	return totalExits
}

func linesToMap(lines []string) map[string][]string {
	inOutMap := make(map[string][]string)
	for _, line := range lines {
		input := strings.Split(line, ":")
		output := strings.Split(strings.Trim(input[1], " "), " ")
		inOutMap[input[0]] = output
	}
	return inOutMap
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

		// Cutting off the newLine character at the end, as it is never in actual input
		data = data[:len(data)-1]
		lines = append(lines, data)
	}

	return lines, nil
}
