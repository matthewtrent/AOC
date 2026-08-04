package main

import (
	"bufio"
	"container/list"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Ranges struct {
	Lower int
	Upper int
}

func main() {
	fileName := "input.txt"
	userArgs := os.Args[1:]
	freshCount := 0

	if len(userArgs) > 0 {
		fileName = userArgs[0]
	}

	fmt.Println("fileName: ", fileName)

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	reader := bufio.NewReader(file)

	ranges := list.New()

	// fill in ranges

	for {
		data, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		// cut off the new line character
		data = data[:len(data)-1]

		if len(data) == 0 {
			break
		}

		bounds := strings.Split(data, "-")

		lower, _ := strconv.Atoi(bounds[0])
		upper, _ := strconv.Atoi(bounds[1])

		currRange := Ranges{
			Lower: lower,
			Upper: upper,
		}

		condenseRange(currRange, ranges)
	}

	// for e := ranges.Front(); e != nil; e = e.Next() {
	// 	selectedRange := e.Value.(Ranges)
	// 	fmt.Println(selectedRange.Lower, "-", selectedRange.Upper)
	// }

	for {
		data, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println("Part1: ", freshCount)
			break
		}

		if err != nil && err != io.EOF {
			fmt.Println(err)
			return
		}

		// cut off the new line character
		data = data[:len(data)-1]

		if len(data) == 0 {
			break
		}

		num, _ := strconv.Atoi(data)

		freshCount += checkInRange(num, ranges)
	}

	fmt.Println("Part2: ", countRange(ranges))
}

func condenseRange(bounds Ranges, ranges *list.List) {
	currRange := bounds

	for e := ranges.Front(); e != nil; {
		next := e.Next()
		selectedRange := e.Value.(Ranges)
		if currRange.Lower > selectedRange.Upper || currRange.Upper < selectedRange.Lower {
			e = next
			continue
		}

		ranges.Remove(e)

		if currRange.Lower > selectedRange.Lower {
			currRange.Lower = selectedRange.Lower
		}
		if currRange.Upper < selectedRange.Upper {
			currRange.Upper = selectedRange.Upper
		}

		e = next
	}

	ranges.PushBack(currRange)
}

func checkInRange(num int, ranges *list.List) int {
	for e := ranges.Front(); e != nil; e = e.Next() {
		selectedRange := e.Value.(Ranges)
		if num >= selectedRange.Lower && num <= selectedRange.Upper {
			return 1
		}
	}

	return 0
}

func countRange(ranges *list.List) int {
	totalSum := 0
	for e := ranges.Front(); e != nil; e = e.Next() {
		selectedRange := e.Value.(Ranges)
		totalSum += selectedRange.Upper - selectedRange.Lower + 1
	}
	return totalSum
}
