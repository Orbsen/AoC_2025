package main

import (
	"fmt"
	"os"
	"strings"
)

type Point struct {
	x int
	y int
}

const placeholder = "."

var directions = []Point{
	{x: 1, y: 0},
	{x: 1, y: 1},
	{x: 1, y: -1},
	{x: 0, y: 1},
	{x: 0, y: -1},
	{x: -1, y: 0},
	{x: -1, y: 1},
	{x: -1, y: -1},
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	countPaperRolls := 0
	maxNumberRemoveablePaperRolls := 0

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	paddedInput := padInput(lines)

	countPaperRolls, _ = getCountPaperRolls(paddedInput)
	maxNumberRemoveablePaperRolls = getMaxNumberRemoveablePaperRolls(paddedInput)

	fmt.Println("solution part 1 = ", countPaperRolls)
	fmt.Println("solution part 2 = ", maxNumberRemoveablePaperRolls)
}

func padInput(lines []string) [][]string {
	var output [][]string
	lengthRow := len(lines[0])

	paddedLine := strings.Repeat(placeholder, lengthRow+2)

	output = append(output, strings.Split(paddedLine, ""))

	for _, line := range lines {
		line = placeholder + line + placeholder
		output = append(output, strings.Split(line, ""))
	}
	output = append(output, strings.Split(paddedLine, ""))

	return output
}

func getCountPaperRolls(paddedInput [][]string) (int, [][]string) {
	count := 0
	paddedOutput := make([][]string, len(paddedInput))

	for i := range paddedInput {
		paddedOutput[i] = make([]string, len(paddedInput[i]))
		copy(paddedOutput[i], paddedInput[i])
	}

	for indexLine, line := range paddedInput {

		if indexLine == 0 || indexLine == len(paddedInput)-1 {
			continue
		}

		for indexRow, char := range line {
			adjacentCounter := 0

			if indexRow == 0 || indexRow == len(line)-1 || char == placeholder {
				continue
			}

			for _, direction := range directions {
				if char == paddedInput[indexLine+direction.x][indexRow+direction.y] {
					adjacentCounter++
				}

				if adjacentCounter == 4 {
					break
				}
			}

			if adjacentCounter < 4 {
				count++
				paddedOutput[indexLine][indexRow] = placeholder
			}
		}
	}

	return count, paddedOutput
}

func getMaxNumberRemoveablePaperRolls(paddedInput [][]string) int {
	maxCount := 0
	var count int

	for {
		count, paddedInput = getCountPaperRolls(paddedInput)
		maxCount += count

		if count == 0 {
			break
		}
	}

	return maxCount
}
