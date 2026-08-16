package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type Coordinate struct {
	X int
	Y int
}

type Square struct {
	CornerCoordOne Coordinate
	CornerCoordTwo Coordinate
	Area           int
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

	areaList, coordList := partOne(lines)
	partTwo(coordList, areaList)
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

func partOne(lines []string) ([]Square, []Coordinate) {
	areaList := make([]Square, 0, 1000)
	coordList := make([]Coordinate, len(lines))
	for i, coord1 := range lines {
		coordOne := stringToCoord(coord1)
		coordList[i] = coordOne
		for j := range i {
			coordTwo := coordList[j]

			area := (abs(coordOne.X-coordTwo.X) + 1) * (abs(coordOne.Y-coordTwo.Y) + 1)
			areaList = append(areaList, Square{
				coordOne,
				coordTwo,
				area,
			})
		}
	}

	sort.Slice(areaList, func(i, j int) bool {
		return areaList[j].Area < areaList[i].Area
	})

	fmt.Println("Part 1: ", areaList[0].Area)
	return areaList, coordList
}

func partTwo(coordList []Coordinate, areaList []Square) {
	// Creating an array of used Xs and Ys to condense the coordinates
	xVals := make([]int, len(coordList))
	yVals := make([]int, len(coordList))

	xMap := make(map[int]int)
	yMap := make(map[int]int)

	for i, coord := range coordList {
		xVals[i] = coord.X
		yVals[i] = coord.Y
	}

	slices.Sort(xVals)
	xVals = slices.Compact(xVals)
	slices.Sort(yVals)
	yVals = slices.Compact(yVals)

	for i, val := range xVals {
		xMap[val] = i
	}

	for i, val := range yVals {
		yMap[val] = i
	}

	// Creating the 2d grid for the condensed coordinates
	// The grid is rounded to the nearest 10 on each side
	grid := make([][]int, len(yVals))
	for i := range len(yVals) {
		grid[i] = make([]int, len(xVals))
	}

	for i, currCoord := range coordList {
		var nextCoord Coordinate

		if i == len(coordList)-1 {
			nextCoord = coordList[0]
		} else {
			nextCoord = coordList[i+1]
		}

		// Set the current coord to true
		grid[yMap[currCoord.Y]][xMap[currCoord.X]] = 1

		// creating the outline of the shape
		for j := xMap[currCoord.X]; j < xMap[nextCoord.X]; j++ {
			grid[yMap[currCoord.Y]][j] = 1
		}

		for j := yMap[currCoord.Y]; j < yMap[nextCoord.Y]; j++ {
			grid[j][xMap[currCoord.X]] = 1
		}

		for j := xMap[currCoord.X]; j > xMap[nextCoord.X]; j-- {
			grid[yMap[currCoord.Y]][j] = 1
		}

		for j := yMap[currCoord.Y]; j > yMap[nextCoord.Y]; j-- {
			grid[j][xMap[currCoord.X]] = 1
		}

	}

	// Floodfill the polygon
	// finding the starting coord
	xDif := 0
	yDif := 0
	if coordList[0].X-coordList[1].X == 0 {
		xDif = coordList[0].X - coordList[len(coordList)-1].X
	} else {
		xDif = coordList[0].X - coordList[1].X
	}

	if coordList[0].Y-coordList[1].Y == 0 {
		yDif = coordList[0].Y - coordList[len(coordList)-1].Y
	} else {
		yDif = coordList[0].Y - coordList[1].Y
	}

	xStartCoord := xMap[coordList[0].X]
	yStartCoord := yMap[coordList[0].Y]

	if xDif < 0 {
		xStartCoord++
	} else {
		xStartCoord--
	}
	if yDif < 0 {
		yStartCoord++
	} else {
		yStartCoord--
	}

	floodFill(xStartCoord, yStartCoord, grid)

	maxArea := 0

	for _, square := range areaList {
		if checkSquare(square, grid, xMap, yMap) {
			maxArea = square.Area
			break
		}
	}

	fmt.Println("Part 2: ", maxArea)
}

func checkSquare(square Square, grid [][]int, xMap map[int]int, yMap map[int]int) bool {
	// Checking the alternate corners are valid, if not then just move on

	if grid[yMap[square.CornerCoordOne.Y]][xMap[square.CornerCoordTwo.X]] != 1 {
		return false
	}

	if grid[yMap[square.CornerCoordTwo.Y]][xMap[square.CornerCoordOne.X]] != 1 {
		return false
	}

	// finding min and max for x & ys

	xMin := xMap[square.CornerCoordTwo.X]
	xMax := xMap[square.CornerCoordOne.X]
	yMin := yMap[square.CornerCoordTwo.Y]
	yMax := yMap[square.CornerCoordOne.Y]

	if square.CornerCoordOne.X < square.CornerCoordTwo.X {
		xMin = xMap[square.CornerCoordOne.X]
		xMax = xMap[square.CornerCoordTwo.X]
	}

	if square.CornerCoordOne.Y < square.CornerCoordTwo.Y {
		yMin = yMap[square.CornerCoordOne.Y]
		yMax = yMap[square.CornerCoordTwo.Y]
	}

	temp := make([][]int, len(grid))
	for i := range grid {
		temp[i] = make([]int, len(grid[i]))
		copy(temp[i], grid[i])
	}
	return floodCheck(xMap[square.CornerCoordOne.X], yMap[square.CornerCoordOne.Y], xMin, xMax, yMin, yMax, temp)
}

func floodCheck(x, y, xMin, xMax, yMin, yMax int, grid [][]int) bool {
	if y < yMin || y > yMax || x < xMin || x > xMax {
		return true
	}
	// Checking if the current coord has been visited
	if grid[y][x] == 2 {
		return true
	}

	curValue := grid[y][x]
	valAsBool := curValue != 0

	grid[y][x] = 2
	upVal := floodCheck(x, y+1, xMin, xMax, yMin, yMax, grid)
	downVal := floodCheck(x, y-1, xMin, xMax, yMin, yMax, grid)
	leftVal := floodCheck(x+1, y, xMin, xMax, yMin, yMax, grid)
	rightVal := floodCheck(x-1, y, xMin, xMax, yMin, yMax, grid)
	return upVal && downVal && leftVal && rightVal && valAsBool
}

func floodFill(x, y int, grid [][]int) [][]int {
	// Checking bounds, should not be needed
	if y < 0 || y >= len(grid) || x < 0 || x >= len(grid[y]) {
		return grid
	}
	// Checking if the current coord has been visited
	if grid[y][x] == 1 {
		return grid
	}

	grid[y][x] = 1
	floodFill(x, y+1, grid)
	floodFill(x, y-1, grid)
	floodFill(x-1, y, grid)
	floodFill(x+1, y, grid)
	return grid
}

// func printGrid(grid [][]int) {
// 	for _, row := range grid {
// 		for _, col := range row {
// 			if col == 1 {
// 				fmt.Print("X ")
// 			} else if col == 2 {
// 				fmt.Print("0 ")
// 			} else {
// 				fmt.Print("· ")
// 			}
// 		}
// 		fmt.Println()
// 	}
// }

func stringToCoord(strCord string) Coordinate {
	coordSlice := strings.Split(strCord, ",")
	x, _ := strconv.Atoi(coordSlice[0])
	y, _ := strconv.Atoi(coordSlice[1])
	coord := Coordinate{
		x,
		y,
	}
	return coord
}

func abs(num int) int {
	if num < 0 {
		return -1 * num
	} else {
		return num
	}
}
