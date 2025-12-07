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
const paperRoll = "@"

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

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	paddedInput := padInput(lines)

	countPaperRolls = getCountPaperRolls(paddedInput)

	fmt.Println("solution part 1 = ", countPaperRolls)
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

func getCountPaperRolls(paddedInput [][]string) int {
	count := 0

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
			}
		}
	}

	return count
}
