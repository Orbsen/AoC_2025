package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	solution1 := getSolutionPart1(lines)
	solution2 := getSolutionPart2(lines)

	fmt.Println("solution part 1 = ", solution1)
	fmt.Println("solution part 2 = ", solution2)
}

func getSolutionPart1(lines []string) int {
	sumJoltages := 0

	for _, line := range lines {
		firstValue, index := getHighestValueAndIndex(line, false)

		croppedLine := line[index+1:]

		secondValue, _ := getHighestValueAndIndex(croppedLine, true)

		combinedValue, _ := strconv.Atoi(firstValue + secondValue)

		sumJoltages += combinedValue
	}

	return sumJoltages
}

func getSolutionPart2(lines []string) int {
	sumMuchLargerJoltages := 0
	var highestJoltage int

	for _, line := range lines {
		highestJoltage = getHighestJoltage(line, 12)
		sumMuchLargerJoltages += highestJoltage
	}

	return sumMuchLargerJoltages
}

func getHighestValueAndIndex(
	line string,
	useLastValue bool,
) (string, int) {
	var digits []int
	highestValue := 0
	highestIndex := 0

	for _, char := range line {
		digits = append(digits, int(char-'0'))
	}

	if !useLastValue {
		digits = digits[:len(digits)-1]
	}

	for i := 0; i < len(digits); i++ {

		if highestValue < digits[i] {
			highestValue = digits[i]
			highestIndex = i
		}
	}

	return strconv.Itoa(highestValue), highestIndex
}

func getHighestJoltage(line string, length int) int {
	var digits []int
	var croppedArray []int
	indexPointer := 0
	outputString := ""

	for _, char := range line {
		digits = append(digits, int(char-'0'))
	}

	for i := length; i > 0; i-- {
		inputString := ""
		croppedArray = digits[indexPointer:]
		croppedArray = croppedArray[:len(croppedArray)-(i-1)]

		for _, digit := range croppedArray {
			inputString += strconv.Itoa(digit)
		}

		index := 0
		highestValue := ""
		highestValue, index = getHighestValueAndIndex(inputString, true)

		outputString += highestValue
		indexPointer += index + 1
	}

	outputInt, _ := strconv.Atoi(outputString)

	return outputInt
}
