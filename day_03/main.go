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

	fmt.Println("solution part 1 = ", solution1)
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
