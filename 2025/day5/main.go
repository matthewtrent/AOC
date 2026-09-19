package main

import (
	"container/list"
	"fmt"
	"strconv"
	"strings"

	"github.com/matthewtrent/aoc/utils"
)

type Ranges struct {
	Lower int
	Upper int
}

func main() {
	lines, err := utils.ReadFile()
	if err != nil {
		print(err)
		return
	}

	ranges, ingredients := createRangeAndIngredients(lines)

	partOne(ranges, ingredients)
	partTwo(ranges)
}

func partOne(ranges *list.List, ingredients []string) {
	freshCount := 0

	for _, ingredient := range ingredients {
		val, _ := strconv.Atoi(ingredient)

		freshCount += checkInRange(val, ranges)
	}

	fmt.Println("Part 1: ", freshCount)
}

func partTwo(ranges *list.List) {
	fmt.Println("Part 2: ", countRange(ranges))
}

func createRangeAndIngredients(lines []string) (*list.List, []string) {
	ranges := list.New()
	ingredients := make([]string, 0)

	for _, line := range lines {

		bounds := strings.Split(line, "-")
		// Then it is a range
		if len(bounds) >= 2 {
			lower, _ := strconv.Atoi(bounds[0])
			upper, _ := strconv.Atoi(bounds[1])

			currRange := Ranges{
				Lower: lower,
				Upper: upper,
			}

			condenseRange(currRange, ranges)
		} else {
			if line == "" {
				continue
			}
			ingredients = append(ingredients, line)
		}

	}
	return ranges, ingredients
}

func condenseRange(bounds Ranges, ranges *list.List) {
	currRange := bounds

	for e := ranges.Front(); e != nil; {
		next := e.Next()
		selectedRange := e.Value.(Ranges)
		if currRange.Lower > selectedRange.Upper || currRange.Upper < selectedRange.Lower {
			e = next
			continue
		}

		ranges.Remove(e)

		if currRange.Lower > selectedRange.Lower {
			currRange.Lower = selectedRange.Lower
		}
		if currRange.Upper < selectedRange.Upper {
			currRange.Upper = selectedRange.Upper
		}

		e = next
	}

	ranges.PushBack(currRange)
}

func checkInRange(num int, ranges *list.List) int {
	for e := ranges.Front(); e != nil; e = e.Next() {
		selectedRange := e.Value.(Ranges)
		if num >= selectedRange.Lower && num <= selectedRange.Upper {
			return 1
		}
	}

	return 0
}

func countRange(ranges *list.List) int {
	totalSum := 0
	for e := ranges.Front(); e != nil; e = e.Next() {
		selectedRange := e.Value.(Ranges)
		totalSum += selectedRange.Upper - selectedRange.Lower + 1
	}
	return totalSum
}
