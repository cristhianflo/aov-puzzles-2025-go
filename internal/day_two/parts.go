package day_two

import (
	"strconv"
)

func PartOne(rangeList *RangeList) int {

	var sum int
	for _, numberRange := range rangeList.Ranges {
		count := numberRange.FirstID
		for count <= numberRange.LastID {
			countStr := strconv.Itoa(count)
			if len(countStr)%2 == 1 {
				count++
				continue
			}
			mid := len(countStr) / 2

			if countStr[:mid] == countStr[mid:] {
				sum = sum + count
				rangeList.InvalidIDs++
			}

			count++
		}

	}

	return sum
}
