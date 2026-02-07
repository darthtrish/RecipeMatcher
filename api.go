package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func searchRecipes(ingredient string) ([]Meal, error) {
	url := fmt.Sprintf("https://www.themealdb.com/api/json/v1/1/filter.php?i=%s", ingredient)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result MealDBResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return result.Meals, nil
}

func getRecipeDetails(mealID string) (*Meal, error) {
	url := fmt.Sprintf("https://www.themealdb.com/api/json/v1/1/lookup.php?i=%s", mealID)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result MealDBResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	if len(result.Meals) == 0 {
		return nil, fmt.Errorf("recipe not found")
	}

	return &result.Meals[0], nil
}
