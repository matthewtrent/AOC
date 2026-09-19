package main

import (
	"fmt"
	"slices"
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
