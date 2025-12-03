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
			}

			count++
		}

	}

	return sum
}

func PartTwo(rangeList *RangeList) int {

	var sum int
	for _, numberRange := range rangeList.Ranges {
		count := numberRange.FirstID
		for count <= numberRange.LastID {
			countStr := strconv.Itoa(count)

			for i := 1; i < len(countStr); i++ {

				if len(countStr)%i != 0 {
					continue
				}

				var parts []string
				for j := 0; j < len(countStr); j += i {
					end := min(j+i, len(countStr))
					parts = append(parts, countStr[j:end])
				}

				first := parts[0]
				isInvalid := true
				for _, part := range parts[1:] {
					if first != part {
						isInvalid = false
					}

				}

				if isInvalid == true {
					sum = sum + count
					break
				}
			}
			count++
		}
	}
	return sum
}
