package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/matthewtrent/aoc/utils"
)

type Present struct {
	PresentShape [][]string
	Area         int
}

type UnderTree struct {
	X             int
	Y             int
	Area          int
	PresentCount  []int
	SumOfPresents int
}

func main() {
	lines, err := utils.ReadFile()
	if err != nil {
		panic(err)
	}

	presents, trees := evalInput(lines)
	partOne(presents, trees)
}

func partOne(presents []Present, trees []UnderTree) {
	totalPossible := 0

	for _, format := range trees {
		if isPossible(presents, format) {
			totalPossible++
		}
	}

	fmt.Println("Part 1: ", totalPossible)
}

func isPossible(presents []Present, floor UnderTree) bool {
	if floor.Area/9 < floor.SumOfPresents {
		return false
	}

	sumOfPresentsArea := 0
	for i := range floor.PresentCount {
		sumOfPresentsArea += presents[i].Area * floor.PresentCount[i]
	}

	fmt.Println(sumOfPresentsArea)

	if floor.Area < sumOfPresentsArea {
		return false
	}

	return true
}

func evalInput(lines []string) ([]Present, []UnderTree) {
	presents := make([]Present, 0)
	floorplans := make([]UnderTree, 0, 100)

	tentativePresent := make([]string, 0)
	wasPresent := false

	for i := range len(lines) {
		line := lines[i]
		if line == "" {
			if wasPresent {
				presents = append(presents, getPresent(tentativePresent))
				wasPresent = false
				tentativePresent = make([]string, 0)
			}
			continue
		}

		if isPresentHeader(line) {
			continue
		}

		if isPresent(line) {
			tentativePresent = append(tentativePresent, line)
			wasPresent = true
		}

		if strings.Contains(line, "x") && strings.Contains(line, ":") {

			tree, err := parseTree(line)
			if err != nil {
				return nil, nil
			}

			floorplans = append(floorplans, tree)
		}

	}

	return presents, floorplans
}

func isPresentHeader(s string) bool {
	if len(s) < 2 || s[len(s)-1] != ':' {
		return false
	}

	_, err := strconv.Atoi(strings.TrimSuffix(s, ":"))
	return err == nil
}

func isPresent(s string) bool {
	if s == "" {
		return false
	}

	for _, r := range s {
		if r != '#' && r != '.' {
			return false
		}
	}

	return true
}

func getPresent(lines []string) Present {
	result := Present{
		PresentShape: make([][]string, len(lines)),
		Area:         0,
	}
	for i, line := range lines {
		section := strings.Split(line, "")

		result.PresentShape[i] = section

		for _, char := range section {
			if char == "#" {
				result.Area++
			}
		}

	}
	return result
}

func parseTree(line string) (UnderTree, error) {
	var x, y int

	// Split at the colon.
	parts := strings.SplitN(line, ":", 2)
	if len(parts) != 2 {
		return UnderTree{}, fmt.Errorf("invalid tree line: %q", line)
	}

	if _, err := fmt.Sscanf(parts[0], "%dx%d", &x, &y); err != nil {
		return UnderTree{}, fmt.Errorf("invalid dimensions: %q", parts[0])
	}

	fields := strings.Fields(parts[1])

	counts := make([]int, len(fields))
	sum := 0

	for i, field := range fields {
		n, err := strconv.Atoi(field)
		if err != nil {
			return UnderTree{}, fmt.Errorf("invalid present count %q: %w", field, err)
		}

		counts[i] = n
		sum += n
	}

	return UnderTree{
		X:             x,
		Y:             y,
		Area:          x * y,
		PresentCount:  counts,
		SumOfPresents: sum,
	}, nil
}
