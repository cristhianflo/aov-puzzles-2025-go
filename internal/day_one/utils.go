package day_one

import (
	"fmt"
	"strconv"
)

func ParseStep(s string) (rune, int, error) {
	if len(s) < 2 {
		return 0, 0, fmt.Errorf("invalid step: %s", s)
	}

	dir := rune(s[0])
	num, err := strconv.Atoi(s[1:])
	if err != nil {
		return 0, 0, err
	}

	return dir, num, nil
}
