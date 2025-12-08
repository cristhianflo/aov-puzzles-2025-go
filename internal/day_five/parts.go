package day_five

func PartOne(cafeteria *Cafeteria) int {
	freshIngredients := 0

	for _, ingredient := range cafeteria.Ingredients {
		for _, ingredientRange := range cafeteria.Ranges {
			if cafeteria.IsFresh(ingredientRange, ingredient) {
				freshIngredients++
				break
			}

		}

	}
	return freshIngredients
}
