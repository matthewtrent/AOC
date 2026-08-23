package main

import (
	"bufio"
	"container/list"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/bartolsthoorn/gohighs/highs"
)

type Machine struct {
	NumLights int
	LightSet  Lights
	ButtonSet []Button
	Joltages  StartJoltage
}

type Lights struct {
	Raw    string
	BitRep int
}

type Button struct {
	Raw    string
	BitRep int
	array  []int
}

type StartJoltage struct {
	Raw      string
	Joltages []int
}

type LightState struct {
	Light       int
	CurrentIter int
}

func main() {
	lines, err := readFile()
	if err != nil {
		fmt.Println(err)
		return
	}

	machines := make([]Machine, 0)

	for _, line := range lines {
		machines = append(machines, strToMachine(line))
	}

	partOne(machines)
	partTwo(machines)
}

func partOne(machines []Machine) {
	totalMin := 0
	for _, machine := range machines {
		totalMin += findMin(machine.LightSet.BitRep, machine.ButtonSet)
	}

	println("Part 1: ", totalMin)
}

func partTwo(machines []Machine) {
	totalMin := 0
	for _, machine := range machines {
		totalMin += lpMin(machine)
	}
	println("Part 2: ", totalMin)
}

func lpMin(machine Machine) int {
	rowCount := machine.NumLights
	colCount := len(machine.ButtonSet)

	colCosts := make([]float64, colCount)
	colLower := make([]float64, colCount)

	for i := range colCount {
		colCosts[i] = 1
	}

	// Creating the variables
	varTypes := make([]highs.VariableType, colCount)
	for i := range varTypes {
		varTypes[i] = highs.Integer
	}

	model := highs.Model{
		ColCosts: colCosts,
		ColLower: colLower,
		VarTypes: varTypes,
	}

	for i := range rowCount {
		row := make([]float64, 0)
		for _, button := range machine.ButtonSet {
			row = append(row, float64(button.array[i]))
		}
		model.AddEqRow(row, float64(machine.Joltages.Joltages[i]))
	}

	solution, err := model.Solve(highs.WithOutput(false))
	if err != nil {
		log.Fatal(err)
	}

	if solution.IsOptimal() {
		return int(solution.Objective)
	}

	return 0
}

func findMin(startVal int, buttons []Button) int {
	queue := list.New()
	queue.PushBack(LightState{
		startVal,
		0,
	})
	for e := queue.Front(); e != nil; e = e.Next() {
		curr := e.Value.(LightState)
		for _, button := range buttons {
			newLight := curr.Light ^ button.BitRep
			if newLight == 0 {
				return curr.CurrentIter + 1
			}

			queue.PushBack(LightState{
				newLight,
				curr.CurrentIter + 1,
			})
		}
	}
	return 0
}

func printMachine(machine Machine) {
	fmt.Print("[", machine.LightSet.Raw, "] ")
	for _, button := range machine.ButtonSet {
		fmt.Print("(", button.Raw, ") ")
	}
	fmt.Println("{", machine.Joltages.Raw, "}")

	fmt.Printf("[%b] ", machine.LightSet.BitRep)

	for _, button := range machine.ButtonSet {
		fmt.Printf("(%b) ", button.BitRep)
	}
	fmt.Println()
}

// Each line should be in the [lights] (button) (button) ... {Joltage}
func strToMachine(line string) Machine {
	var sectionLine []string
	var lights Lights
	var joltage StartJoltage
	button := make([]Button, 0)
	totalLights := 0

	for char := range strings.SplitSeq(line, "") {
		switch char {
		case "[", "(", "{":
			sectionLine = make([]string, 0)
		case "]":
			totalLights = len(sectionLine)
			lights = parseLights(sectionLine)
		case ")":
			button = append(button, parseButton(sectionLine, totalLights))
		case "}":
			joltage = parseJoltage(sectionLine)
		case " ":
			// Do Nothing
		default:
			sectionLine = append(sectionLine, char)
		}
	}
	return Machine{
		totalLights,
		lights,
		button,
		joltage,
	}
}

func parseLights(data []string) Lights {
	val := 0

	for i, char := range data {
		if char == "#" {
			val += 1 << i
		}
	}

	return Lights{
		strings.Join(data, ""),
		val,
	}
}

func parseButton(data []string, numLights int) Button {
	var sb strings.Builder
	val := 0
	array := make([]int, numLights)
	for _, char := range data {
		if char != "," {
			sb.WriteString(char)
		} else {
			num, _ := strconv.Atoi(sb.String())
			if num < numLights {
				val += 1 << num
				array[num] = 1
				sb.Reset()
			}
		}
	}

	num, _ := strconv.Atoi(sb.String())
	if num < numLights {
		val += 1 << num
		array[num] = 1
		sb.Reset()
	}
	return Button{
		strings.Join(data, ""),
		val,
		array,
	}
}

func parseJoltage(data []string) StartJoltage {
	var sb strings.Builder
	vals := make([]int, 0)
	for _, char := range data {
		if char != "," {
			sb.WriteString(char)
		} else {
			num, _ := strconv.Atoi(sb.String())
			vals = append(vals, num)
			sb.Reset()
		}
	}
	num, _ := strconv.Atoi(sb.String())
	vals = append(vals, num)

	return StartJoltage{
		strings.Join(data, ""),
		vals,
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
