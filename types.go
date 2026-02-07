package main

type MealDBResponse struct {
	Meals []Meal `json:"meals"`
}

type Meal struct {
	ID           string `json:"idMeal"`
	Name         string `json:"strMeal"`
	Instructions string `json:"strInstructions"`
	Thumbnail    string `json:"strMealThumb"`
	Source       string `json:"strSource"`
}
