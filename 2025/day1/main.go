package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fileName := "input.txt"
	currNum := 50
	numZero := 0

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
		gotToZero := 0
		if err != nil {
			fmt.Println(err)
			fmt.Println("stopped at zero: ", numZero)
			return
		}

		//fmt.Println(len(data))
		//fmt.Print(data)

		currNum, gotToZero = rotateLock(data, currNum)
		numZero += gotToZero
	}
}

/*
* simulates rotating a lock dial, numbered 0-99
*	@Args: data - the rotation direction and the amount rotated, example R55
*				currNum - the current dial number
*	@returns:
*				int - the new current dial number
*				int - the amount of times that the dial landed exactly on 0
* */
func rotateLock(data string, currNum int) (int, int) {
	total0 := 0

	rotateDir := 1 // Default rotate left

	if data[0] == 'L' {
		rotateDir = -1
	}

	// Extract the number of times to rotate
	num := data[1 : len(data)-1]
	numSpin, _ := strconv.Atoi(num)

	for range numSpin {
		currNum += rotateDir
		if currNum == -1 {
			currNum = 99
		}

		currNum = currNum % 100

		if currNum == 0 {
			total0 += 1
		}
	}

	return currNum, total0
}
