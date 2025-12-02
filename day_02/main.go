package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type IdRange struct {
	from int
	to   int
}

func main() {
	data, err := os.ReadFile("input.txt")
	//data, err := os.ReadFile("input_test.txt")
	if err != nil {
		panic(err)
	}

	rangeStrings := strings.Split(string(data), ",")
	idRanges := parseIdRanges(rangeStrings)

	invalidIds := findInvalidIds(idRanges)

	var sumInvalidIds int = 0
	for _, id := range invalidIds {
		sumInvalidIds += id
	}

	fmt.Println("invalid ids sum = ", sumInvalidIds)
}

func parseIdRanges(rangeStrings []string) []IdRange {
	var idRanges []IdRange

	for _, rangeString := range rangeStrings {
		ranges := strings.Split(rangeString, "-")

		startString := strings.TrimSpace(ranges[0])
		endString := strings.TrimSpace(ranges[1])

		startID, err := strconv.Atoi(startString)
		if err != nil {
			panic(err)
		}

		endID, err := strconv.Atoi(endString)
		if err != nil {
			panic(err)
		}

		idRanges = append(idRanges, IdRange{
			from: startID,
			to:   endID,
		})
	}

	return idRanges
}

func findInvalidIds(idRanges []IdRange) []int {
	var invalidIds []int

	maxEndID := getMaxEndId(idRanges)
	maxDigits := len(strconv.Itoa(maxEndID))

	for fullLength := 2; fullLength <= maxDigits; fullLength += 2 {
		blockLength := fullLength / 2

		var blockStart int
		var blockEnd int

		switch blockLength {
		case 1:
			blockStart = 1
			blockEnd = 9
		case 2:
			blockStart = 10
			blockEnd = 99
		case 3:
			blockStart = 100
			blockEnd = 999
		case 4:
			blockStart = 1000
			blockEnd = 9999
		case 5:
			blockStart = 10000
			blockEnd = 99999
		default:
			break
		}

		for blockValue := blockStart; blockValue <= blockEnd; blockValue++ {
			blockString := strconv.Itoa(blockValue)

			fullString := blockString + blockString

			fullId, err := strconv.Atoi(fullString)
			if err != nil {
				panic(err)
			}

			if fullId > maxEndID {
				break
			}

			if inRange(fullId, idRanges) {
				invalidIds = append(invalidIds, fullId)
			}
		}
	}

	return invalidIds
}

func inRange(id int, idRanges []IdRange) bool {
	for _, ranges := range idRanges {
		if id >= ranges.from && id <= ranges.to {
			return true
		}
	}
	return false
}

func getMaxEndId(idRanges []IdRange) int {
	var maxEndID int
	for i, ranges := range idRanges {
		if i == 0 || ranges.to > maxEndID {
			maxEndID = ranges.to
		}
	}
	return maxEndID
}
