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
	if err != nil {
		panic(err)
	}
	ranges, ids := parseInput(data)

	countFreshIds := getCountofFreshIds(ranges, ids)

	fmt.Println("solution part 1 = ", countFreshIds)
}

func parseInput(data []byte) ([]IdRange, []int) {
	lines := strings.Split(string(data), "\n")
	var ranges []IdRange
	var ids []int

	for _, line := range lines {
		if strings.Contains(line, "-") {
			parts := strings.Split(line, "-")
			from, _ := strconv.Atoi(parts[0])
			to, _ := strconv.Atoi(parts[1])

			ranges = append(ranges, IdRange{from: from, to: to})
		} else {
			id, _ := strconv.Atoi(line)
			ids = append(ids, id)
		}
	}

	ids = ids[1:]

	return ranges, ids
}

func getCountofFreshIds(idRanges []IdRange, ids []int) int {
	count := 0

	for _, id := range ids {
		for _, idRange := range idRanges {
			if id >= idRange.from && id <= idRange.to {
				count++
				break
			}
		}
	}
	return count
}
