package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Distance struct {
	CoordsOne string
	CoordsTwo string
	Distance  int
}

func main() {
	lines, err := readFile()
	if err != nil {
		fmt.Println(err)
		return
	}

	// for _, line := range lines {
	// 	fmt.Println(line)
	// }

	sortedDist := sortedDirections(lines)
	partOne(sortedDist)
	partTwo(lines, sortedDist)
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

		data = data[:len(data)-1]
		lines = append(lines, data)
	}

	return lines, nil
}

func partOne(sortedDist []Distance) {
	// create my sorted Array of lengths
	circuits := make([][]string, 0, 100)

	for i := range 1000 {
		currNode := sortedDist[i]
		firstIndex := -1
		secondIndex := -1
		for j, circuit := range circuits {
			for _, box := range circuit {
				if currNode.CoordsOne == box {
					firstIndex = j
				}

				if currNode.CoordsTwo == box {
					secondIndex = j
				}
			}
		}

		if firstIndex == -1 && secondIndex == -1 {
			circuits = append(circuits, []string{currNode.CoordsOne, currNode.CoordsTwo})
		}
		if firstIndex == -1 && secondIndex != -1 {
			circuits[secondIndex] = append(circuits[secondIndex], currNode.CoordsOne)
		}
		if firstIndex != -1 && secondIndex == -1 {
			circuits[firstIndex] = append(circuits[firstIndex], currNode.CoordsTwo)
		}
		if firstIndex != -1 && secondIndex != -1 && firstIndex != secondIndex {
			circuits[firstIndex] = append(circuits[firstIndex], circuits[secondIndex]...)
			circuits[secondIndex] = []string{}
		}
	}

	top1 := 0
	top2 := 0
	top3 := 0

	for _, circuit := range circuits {
		//fmt.Println(circuit)
		if len(circuit) > top1 {
			top3 = top2
			top2 = top1
			top1 = len(circuit)
			continue
		}
		if len(circuit) > top2 {
			top3 = top2
			top2 = len(circuit)
			continue
		}
		if len(circuit) > top3 {
			top3 = len(circuit)
			continue
		}
	}

	fmt.Println("Part 1: ", top1*top2*top3)
}

func partTwo(lines []string, sortedDist []Distance) {
	// create my sorted Array of lengths
	circuits := make([][]string, 0, 100)
	println(len(lines))

	finalXmult := 0
	for i := range len(sortedDist) {
		currNode := sortedDist[i]
		firstIndex := -1
		secondIndex := -1
		for j, circuit := range circuits {
			for _, box := range circuit {
				if currNode.CoordsOne == box {
					firstIndex = j
				}

				if currNode.CoordsTwo == box {
					secondIndex = j
				}
			}
		}

		if firstIndex == -1 && secondIndex == -1 {
			circuits = append(circuits, []string{currNode.CoordsOne, currNode.CoordsTwo})
		}
		if firstIndex == -1 && secondIndex != -1 {
			circuits[secondIndex] = append(circuits[secondIndex], currNode.CoordsOne)
			if len(circuits[secondIndex]) == len(lines) {
				finalXmult = multiplyXDist(currNode)
				break
			}
		}
		if firstIndex != -1 && secondIndex == -1 {
			circuits[firstIndex] = append(circuits[firstIndex], currNode.CoordsTwo)
			if len(circuits[firstIndex]) == len(lines) {
				finalXmult = multiplyXDist(currNode)
				break
			}
		}
		if firstIndex != -1 && secondIndex != -1 && firstIndex != secondIndex {
			circuits[firstIndex] = append(circuits[firstIndex], circuits[secondIndex]...)
			circuits[secondIndex] = []string{}

			if len(circuits[firstIndex]) == len(lines) {
				finalXmult = multiplyXDist(currNode)
				break
			}
		}
	}

	fmt.Println("Part 2: ", finalXmult)
}

func multiplyXDist(currNode Distance) int {
	coordOne := strings.Split(currNode.CoordsOne, ",")
	coordTwo := strings.Split(currNode.CoordsTwo, ",")
	return mustAtoi(coordOne[0]) * mustAtoi(coordTwo[0])
}

func squared(num int) int {
	return num * num
}

func mustAtoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

func sortedDirections(lines []string) []Distance {
	distances := make([]Distance, 0, squared(len(lines)))

	for i, coordOne := range lines {
		for j := range i {
			coordTwo := lines[j]
			// Distance is 0, no reason to store the nodes
			if coordOne == coordTwo {
				continue
			}
			coordsOne := strings.Split(coordOne, ",")
			coordsTwo := strings.Split(coordTwo, ",")

			distance := squared((mustAtoi(coordsOne[0]) - mustAtoi(coordsTwo[0]))) +
				squared((mustAtoi(coordsOne[1]) - mustAtoi(coordsTwo[1]))) +
				squared((mustAtoi(coordsOne[2]) - mustAtoi(coordsTwo[2])))

			distances = append(distances, Distance{
				CoordsOne: coordOne,
				CoordsTwo: coordTwo,
				Distance:  distance,
			})
		}
	}

	sort.Slice(distances, func(i, j int) bool {
		return distances[i].Distance < distances[j].Distance
	})
	return distances
}
