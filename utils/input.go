// Package utils provides useful repeated functions for advent of code
package utils

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

func ReadFile() ([]string, error) {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		return nil, fmt.Errorf("could not determine caller")
	}

	dir := filepath.Dir(filename)
	fileName := "input.txt"

	if len(os.Args) > 1 {
		fileName = os.Args[1]
	}

	filePath := filepath.Join(dir, fileName)

	fmt.Println("FilePath:", filePath)

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
