package day_five

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Ingredient struct {
	ID int
}

type IngredientRange struct {
	FirstID int
	LastID  int
}

type Cafeteria struct {
	Ranges      []IngredientRange
	Ingredients []Ingredient
}

func NewCafeteria(path string) (*Cafeteria, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var sections [][]string
	var prevRange []string
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			sections = append(sections, prevRange)
			prevRange = []string{}
			continue
		}

		prevRange = append(prevRange, line)
	}

	sections = append(sections, prevRange)
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	var ranges []IngredientRange
	var ingredients []Ingredient

	for i, sec := range sections {
		for _, line := range sec {
			switch i {
			case 0:
				rangeString := strings.Split(line, "-")
				firstID, err := strconv.Atoi(rangeString[0])
				if err != nil {
					return nil, err
				}
				lastID, err := strconv.Atoi(rangeString[1])
				if err != nil {
					return nil, err
				}
				ranges = append(ranges, IngredientRange{
					FirstID: firstID,
					LastID:  lastID,
				})
			case 1:
				ingredientID, err := strconv.Atoi(line)
				if err != nil {
					return nil, err
				}
				ingredients = append(ingredients, Ingredient{
					ID: ingredientID,
				})
			}
		}
	}

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].FirstID < ranges[j].FirstID
	})

	return &Cafeteria{
		Ranges:      ranges,
		Ingredients: ingredients,
	}, nil
}

func (c *Cafeteria) MergeIngredientRanges() {
	merged := []IngredientRange{c.Ranges[0]}

	for _, currentRange := range c.Ranges[1:] {
		prevRange := merged[len(merged)-1]
		if currentRange.FirstID <= prevRange.LastID {
			merged[len(merged)-1].LastID = max(prevRange.LastID, currentRange.LastID)
		} else {
			merged = append(merged, currentRange)
		}

	}
	c.Ranges = merged
}

func (c *Cafeteria) IsFresh(ingredientRange IngredientRange, ingredient Ingredient) bool {
	return ingredient.ID >= ingredientRange.FirstID && ingredient.ID <= ingredientRange.LastID
}
