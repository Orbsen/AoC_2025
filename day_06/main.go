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
	solutionPart2 := getSolutionPart2(data, operator)

	fmt.Println("solutions part 1 =", solutionPart1)
	fmt.Println("solutions part 2 =", solutionPart2)
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

func getSolutionPart2(data []byte, operator []string) int {
	var result = 0

	blocks := parseBlocks(data)

	for index, block := range blocks {
		rearrangedNumbers := getRearrangedNumbersForPart2(block)

		innerResult := 1
		if operator[index] == additionOperator {
			for _, value := range rearrangedNumbers {
				innerResult += value
			}
			innerResult = innerResult - 1
		} else {
			for _, value := range rearrangedNumbers {
				innerResult *= value
			}
		}
		result += innerResult
	}

	return result
}

func getRearrangedNumbersForPart2(block []string) []int {

	rows := len(block)
	cols := len(block[0])

	reArrangedConvertedNumbers := make([][]string, cols)

	for i := 0; i < cols; i++ {
		reArrangedConvertedNumbers[i] = make([]string, rows)
		for j := 0; j < rows; j++ {
			reArrangedConvertedNumbers[i][j] = string(block[j][i])
		}
	}

	var calculatedNumbers []int
	for _, row := range reArrangedConvertedNumbers {
		maxLength := getMaxLength(row)
		transformRow := getTransformedRow(row, maxLength)

		calculatedNumbers = append(calculatedNumbers, transformRow...)
	}

	return calculatedNumbers
}

func getMaxLength(row []string) int {
	maxLength := 0
	for _, value := range row {
		if len(value) > maxLength {
			maxLength = len(value)
		}
	}
	return maxLength
}

func getTransformedRow(input []string, maxLen int) []int {
	output := make([]int, 0, maxLen)

	for col := maxLen - 1; col >= 0; col-- {
		var sb strings.Builder

		for _, rowStr := range input {
			if col < len(rowStr) {
				sb.WriteByte(rowStr[col])
			}
		}

		rawString := sb.String()
		cleanString := strings.TrimSpace(rawString)

		if len(cleanString) > 0 {
			number, err := strconv.Atoi(cleanString)
			if err == nil {
				output = append(output, number)
			}
		}
	}

	return output
}

func parseBlocks(data []byte) [][]string {
	lines := strings.Split(string(data), "\n")

	var grid []string
	maxLen := 0
	for _, line := range lines {
		if strings.ContainsAny(line, "*+") || len(strings.TrimSpace(line)) == 0 {
			continue
		}
		grid = append(grid, line)
		if len(line) > maxLen {
			maxLen = len(line)
		}
	}

	var blocks [][]string
	var currentBlock []string

	for range grid {
		currentBlock = append(currentBlock, "")
	}

	startCol := 0
	for col := 0; col <= maxLen; col++ {
		isEmptyCol := true

		if col < maxLen {
			for _, line := range grid {
				if col < len(line) && line[col] != ' ' {
					isEmptyCol = false
					break
				}
			}
		}

		if isEmptyCol || col == maxLen {
			if col > startCol {
				var blockLines []string
				for i, line := range grid {

					end := col
					if end > len(line) {
						end = len(line)
					}
					start := startCol
					if start > len(line) {
						start = len(line)
					}

					blockLines = append(blockLines, line[start:end])

					for len(blockLines[i]) < (col - startCol) {
						blockLines[i] += " "
					}
				}
				blocks = append(blocks, blockLines)
			}
			startCol = col + 1
		}
	}
	return blocks
}
