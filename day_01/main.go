package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	//data, err := os.ReadFile("input_test.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	solutionPart1, solutionPart2 := part1(lines)

	fmt.Println("solution part 1 = ", solutionPart1)
	fmt.Println("solution part 2 =", solutionPart2)
}

func part1(lines []string) (int, int) {
	var dialPos int = 50
	var countZeros int = 0
	var countRollOvers int = 0

	for _, line := range lines {
		direction := line[0:1]
		amount, err := strconv.Atoi(line[1:])

		if err != nil {
			panic(err)
		}

		if direction == "R" {
			countRollOvers += (dialPos + amount) / 100

			dialPos = dialRight(dialPos, amount)
		} else {
			if dialPos == 0 {
				countRollOvers += amount / 100
			} else if amount < dialPos {
				// no roll overs
			} else {
				countRollOvers += 1 + (amount-dialPos)/100
			}

			dialPos = dialLeft(dialPos, amount)
		}

		if dialPos == 0 {
			countZeros++
		}
	}

	return countZeros, countRollOvers
}

func dialRight(dialPos int, amount int) int {

	if amount > 100 {
		amount = amount % 100
	}

	var returnValue = dialPos + amount

	if returnValue > 99 {
		returnValue = returnValue % 100
	}

	return returnValue
}

func dialLeft(dialPos int, amount int) int {

	if amount > 100 {
		amount = amount % 100
	}

	var returnValue = dialPos - amount

	if returnValue < 0 {
		returnValue = 100 + returnValue
	}

	return returnValue
}
