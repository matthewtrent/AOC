package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fileName := "input.txt"
	totalSum := 0

	userArgs := os.Args[1:]

	if len(userArgs) > 0 {
		fileName = userArgs[0]
	}

	fmt.Println("fileName: ", fileName)

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		data, err := reader.ReadString('\n')

		//fmt.Println(data)
		if len(data) > 0 {
			data = data[:len(data)-1]
			totalSum = totalSum + largestBattery(data, 12)
		}
		if err != nil {
			fmt.Println("totalSum: ", totalSum)
			fmt.Println(err)
			return
		}
	}
}

func largestBattery(data string, length int) int {
	largest := make([]rune, length)

	for i, curRune := range data {
		runesLeft := len(data) - i
		indexToCheck := 0

		if runesLeft < length {
			indexToCheck = i - (len(data) - length)
		}

		// fmt.Println(i, len(data))
		// fmt.Println(indexToCheck)
		// fmt.Printf("%c %c\n", curRune, largest[indexToCheck])

		replaced := false
		for j := indexToCheck; j < length; j++ {
			if replaced {
				largest[j] = '0'
			} else if curRune > largest[j] {
				largest[j] = curRune
				replaced = true
			}
		}
	}

	val, _ := strconv.Atoi(string(largest))
	fmt.Println(val)

	return val
}
