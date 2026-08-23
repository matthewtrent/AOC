package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
		data, err := reader.ReadString(',')

		fmt.Println(data)
		if len(data) > 0 {
			totalSum = totalSum + doubleNum(data)
		}
		if err != nil {

			fmt.Println("totalSum: ", totalSum)
			fmt.Println(err)
			return
		}

		//fmt.Println(len(data))
		//fmt.Println(data)
		//fmt.Println(totalSum)
	}
}

func doubleNum(data string) int {
	numSum := 0

	ranges := strings.Split(data[:len(data)-1], "-")

	startInt, _ := strconv.Atoi(ranges[0])
	endInt, _ := strconv.Atoi(ranges[1])

	//fmt.Println(startInt)
	//fmt.Println(endInt)
	for startInt <= endInt {
		currNum := strconv.Itoa(startInt)

		// break the number into sections,
		maxLength := (len(currNum) / 2)

		for i := 1; i <= maxLength; i++ {
			foundPattern := false
			// Check if length % i == 0
			//	if not then it cant seperate evenly and no reason to check sections
			if len(currNum)%i != 0 {
				// fmt.Println("Cant break ", currNum, " into even spaces of ", i)
				continue
			}

			potentialSequence := currNum[:i]

			// j is the number of potential matching blocks
			for j := 1; j <= (len(currNum)-1)/i; j++ {
				// fmt.Println((len(currNum) - 1) / i)
				// fmt.Println(potentialSequence, " ", currNum[i*j:i*(j+1)])
				if potentialSequence != currNum[i*j:i*(j+1)] {
					// fmt.Println(currNum)
					break
				}

				if j == ((len(currNum) - 1) / i) {
					// fmt.Println(currNum)
					numSum += startInt
					foundPattern = true
					break
				}
			}

			if foundPattern {
				break
			}

			//	if it is
		}

		startInt++
	}

	//fmt.Println(doubleNumSum)

	return numSum
}
