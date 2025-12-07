package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const additionOperator = "+"
const multiplicationOperator = "*"

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	numbers, operator := parseInput(data)

	solutionPart1 := getSolutionPart1(numbers, operator)

	fmt.Println("solutions part 1 =", solutionPart1)
}

func parseInput(data []byte) ([][]int, []string) {
	var numbers [][]int
	var innerNumbers []int
	var stringValue []string

	lines := strings.Split(string(data), "\n")

	for index, line := range lines {
		stringValue = strings.Fields(line)

		_, err := strconv.Atoi(stringValue[0])
		if err != nil {
			continue
		}

		for _, str := range stringValue {
			number, _ := strconv.Atoi(str)
			innerNumbers = append(innerNumbers, number)
		}

		if index < len(lines)-1 {
			numbers = append(numbers, innerNumbers)
			innerNumbers = []int{}
		}
	}
	return numbers, stringValue

}

func getSolutionPart1(numbers [][]int, operator []string) int {
	count := 0

	for index, char := range operator {
		innerCount := 1
		if char == additionOperator {
			for i := 0; i < len(numbers); i++ {
				innerCount += numbers[i][index]
			}
			innerCount = innerCount - 1
		} else {
			for i := 0; i < len(numbers); i++ {
				innerCount *= numbers[i][index]
			}
		}
		count += innerCount
	}

	return count
}
