package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	lines, err := readFile()
	if err != nil {
		fmt.Println(err)
		return
	}

	sol(lines)
}

func sol(lines []string) {
	for i := range lines {
		floor := 0
		firstBasement := -1
		for j, char := range lines[i] {
			switch char {
				case '(':
					floor++
				case ')':
					floor--
				default:
			}

			if(floor == -1 && firstBasement == -1) {
				firstBasement = j + 1
			}
		}
		fmt.Println("Line ", i, " has goes to floor ", floor)
		fmt.Println("Line ", i, " goes to -1 at ", firstBasement)
	}
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
