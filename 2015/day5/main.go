package main

import (
	"fmt"
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
	totalNice := 0

	for _, line := range lines {
		totalNice += checkNaughtOrNiceBasic(line)
	}

	fmt.Println("Part One: ", totalNice)
}

func checkNaughtOrNiceBasic(line string) int {
	totalVowels := 0
	doubleLetters := false
	notExcluded := true

	vowels := "aeiou"

	if strings.ContainsRune(vowels, rune(line[0])) {
		totalVowels++
	}

	for i := 1; i < len(line); i++ {
		charOne := rune(line[i])
		charTwo := rune(line[i-1])

		// Checking vowels
		if strings.ContainsRune(vowels, charOne) {
			totalVowels++
		}

		// Checking double letters
		if charOne == charTwo {
			doubleLetters = true
		}

		switch line[i-1 : i+1] {
		case "ab":
			notExcluded = false
		case "cd":
			notExcluded = false
		case "pq":
			notExcluded = false
		case "xy":
			notExcluded = false
		}

	}

	if totalVowels >= 3 && doubleLetters && notExcluded {
		return 1
	} else {
		return 0
	}
}

func partTwo(lines []string) {
	totalNice := 0

	for _, line := range lines {
		totalNice += checkNaughtOrNiceComplex(line)
	}

	fmt.Println("Part Two: ", totalNice)
}

func checkNaughtOrNiceComplex(line string) int {
	foundDoubleRepeat := false
	foundSeperatedDouble := false

	for i := 0; i <= len(line); i++ {
		// Checking that the same letter appears between a single letter
		endChar := i + 2
		if endChar < len(line) {
			if line[i] == line[endChar] {
				foundSeperatedDouble = true
				break
			}
		}
	}

	endIndex := 2

	for startIndex := 0; startIndex < len(line); startIndex++ {
		if endIndex <= len(line) && !foundDoubleRepeat {
			possiblePair := line[startIndex:endIndex]

			for j := endIndex; j <= len(line); j++ {
				testEnd := j + 2
				if testEnd <= len(line) && possiblePair == line[j:testEnd] {
					foundDoubleRepeat = true
				}
			}
		}
		endIndex++
	}

	if foundDoubleRepeat && foundSeperatedDouble {
		return 1
	} else {
		return 0
	}
}
