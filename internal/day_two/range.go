package day_two

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	FirstID int
	LastID  int
}

type RangeList struct {
	Ranges     []Range
	InvalidIDs int
}

func NewRangeList(path string) (*RangeList, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var builder strings.Builder
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		builder.WriteString(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	content := builder.String()
	var numberRanges []Range

	for ranges := range strings.SplitSeq(content, ",") {
		parts := strings.Split(ranges, "-")
		firstID, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, err
		}
		lastID, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, err
		}
		numberRange := Range{
			FirstID: firstID,
			LastID:  lastID,
		}
		numberRanges = append(numberRanges, numberRange)
	}

	return &RangeList{
		Ranges:     numberRanges,
		InvalidIDs: 0,
	}, nil
}
