package day_five

import (
	"bufio"
	"os"
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
	Ranges           []IngredientRange
	Ingredients      []Ingredient
	FreshIngredients int
}

func NewCafeteria(path string) (*Cafeteria, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var sections [][]string
	var current []string
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			sections = append(sections, current)
			current = []string{}
			continue
		}

		current = append(current, line)
	}

	sections = append(sections, current)
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

	return &Cafeteria{
		Ranges:           ranges,
		Ingredients:      ingredients,
		FreshIngredients: 0,
	}, nil
}

func (c *Cafeteria) CountFreshIngredient() {
	c.FreshIngredients++
}

func (c *Cafeteria) IsFresh(ingredientRange IngredientRange, ingredient Ingredient) bool {
	return ingredient.ID >= ingredientRange.FirstID && ingredient.ID <= ingredientRange.LastID
}
