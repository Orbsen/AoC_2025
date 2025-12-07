package main

import (
	"fmt"
	"os"
	"sort"
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
	countAllFreshIds := getCountOffAllFreshIds(ranges)

	fmt.Println("solution part 1 = ", countFreshIds)
	fmt.Println("solution part 2 = ", countAllFreshIds)
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

func getCountOffAllFreshIds(idRanges []IdRange) int {
	count := 0

	sort.Slice(idRanges, func(i, j int) bool {
		return idRanges[i].from < idRanges[j].from
	})

	mergedRanges := []IdRange{idRanges[0]}

	for _, current := range idRanges[1:] {
		lastMergedIndex := len(mergedRanges) - 1
		lastMerged := mergedRanges[lastMergedIndex]

		if current.from <= lastMerged.to {
			if current.to > lastMerged.to {
				mergedRanges[lastMergedIndex].to = current.to
			}
		} else {
			mergedRanges = append(mergedRanges, current)
		}
	}

	for _, idRange := range mergedRanges {
		result := idRange.to - idRange.from
		count += result + 1
	}
	return count
}
