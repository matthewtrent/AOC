package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Index struct {
	X int
	Y int
}

func main() {
	fileName := "input.txt"
	total := 0

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

	var lines []string
	firstLine := true

	for {
		data, err := reader.ReadString('\n')
		//fmt.Println(data)

		if err == io.EOF {
			break
		}

		// Not optimal, ideal would set a large size and not require runtime expansion
		// Though I dont know how len() would work with it
		if len(data) > 0 {
			// Remove new line character
			data = data[:len(data)-1]
			data = "." + data + "."
			if firstLine {
				lines = append(lines, strings.Repeat(".", len(data)))
				firstLine = false
			}
			lines = append(lines, data)
		}

		if err != nil {
			fmt.Println(err)
			return
		}
	}

	lines = append(lines, strings.Repeat(".", len(lines[0])))

	for {
		foundIter := 0
		foundIter, lines = openRolls(lines, 4)
		if foundIter == 0 {
			break
		}

		total += foundIter
	}

	// for _, line := range lines {
	// 	fmt.Println(line)
	// }

	fmt.Println(total)
}

func openRolls(wall []string, closedRolls int) (int, []string) {
	total := 0
	var takenRoll []Index
	// Go through each row of the wall
	for i := 1; i < len(wall)-1; i++ {
		// fmt.Println(wall[i-1])
		// fmt.Println(wall[i])
		// fmt.Println(wall[i+1])
		for j := 1; j < len(wall[i])-1; j++ {

			//fmt.Printf("currRune: %c\n", wall[i][j])

			subTotal := 0
			if wall[i][j] == '@' {
				if wall[i-1][j-1] == '@' {
					subTotal += 1
				}
				if wall[i][j-1] == '@' {
					subTotal += 1
				}
				if wall[i+1][j-1] == '@' {
					subTotal += 1
				}
				if wall[i-1][j] == '@' {
					subTotal += 1
				}
				if wall[i+1][j] == '@' {
					subTotal += 1
				}
				if wall[i-1][j+1] == '@' {
					subTotal += 1
				}
				if wall[i][j+1] == '@' {
					subTotal += 1
				}
				if wall[i+1][j+1] == '@' {
					subTotal += 1
				}
				if subTotal < closedRolls {
					//fmt.Println("Less than 4 @ found")
					total += 1
					takenRoll = append(takenRoll, Index{
						X: i,
						Y: j,
					})
				}
			}

		}
	}

	for _, index := range takenRoll {
		line := []rune(wall[index.X])
		line[index.Y] = '.'
		wall[index.X] = string(line)
	}

	return total, wall
}
