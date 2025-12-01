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

	solutionPart1 := part1(lines)

	fmt.Println(solutionPart1)
}

func part1(lines []string) int {
	var dialPos int = 50
	var countZeros int = 0

	for _, line := range lines {
		direction := line[0:1]
		amount, err := strconv.Atoi(line[1:])

		if err != nil {
			panic(err)
		}
		//fmt.Println("direction, amount")
		//fmt.Println(direction, amount)

		if direction == "R" {
			dialPos = dialRight(dialPos, amount)
		} else {
			dialPos = dialLeft(dialPos, amount)
		}

		//fmt.Println(dialPos)

		if dialPos == 0 {
			countZeros++
		}
	}

	return countZeros
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
