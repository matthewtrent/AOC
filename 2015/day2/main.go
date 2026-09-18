package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	lines, err := readFile()
	if err != nil {
		fmt.Println(err)
		return
	}

	partOne(lines)
}

func partOne(lines []string) {
	//line is in the form
	//l x w x h
	//with expected output being surface area + smallest size
	totalPaper := 0
	totalRibbon := 0
	for _, line := range lines {
		// split the string, with seperator x
		vals := strSliceToIntSlice(strings.Split(line, "x"))
		totalPaper += surfaceAreaPlusSmallest(vals[0], vals[1], vals[2])
		totalRibbon += ribbonNeeded(vals[0], vals[1], vals[2])
	}

	fmt.Println("Part 1: ", totalPaper)
	fmt.Println("Part 2: ", totalRibbon)
}

func ribbonNeeded(l, w, h int) int {
	volume := l * w * h
	sides := []int{l, w, h}

	slices.Sort(sides)

	smallestPerim := 2*sides[0] + 2*sides[1]

	return volume + smallestPerim
}

func strSliceToIntSlice(nums []string) []int {
	// Pre-allocate memory for efficiency
	intSlice := make([]int, 0, len(nums))

	for _, str := range nums {
		val, err := strconv.Atoi(str)
		if err != nil {
			// Handle the conversion error appropriately
			fmt.Printf("Error converting %s: %v\n", str, err)
			continue
		}
		intSlice = append(intSlice, val)
	}

	return intSlice
}

func surfaceAreaPlusSmallest(l, w, h int) int {
	sideOne := l * w
	sideTwo := l * h
	sideThree := w * h
	return 2*(sideOne+sideTwo+sideThree) + min(sideOne, sideTwo, sideThree)
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
