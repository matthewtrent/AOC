package main

import (
	"crypto/md5"
	"encoding/hex"
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
	fmt.Println("Part One: ", hashWithLeadingNumberZero(lines[0], 5))
}

func partTwo(lines []string) {
	fmt.Println("Part Two: ", hashWithLeadingNumberZero(lines[0], 6))
}

func hashWithLeadingNumberZero(secret string, numberZero int) int {
	goal := strings.Repeat("0", numberZero)
	startingHex := ""
	num := 0
	for {
		testSecret := secret + strconv.Itoa(num)
		hash := getMD5Hash(testSecret)
		startingHex = hash[:numberZero]

		if startingHex == goal {
			return num
		}
		num++
	}
}

func getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
