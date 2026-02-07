package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		showMainMenu()
		choice := readInput(reader)

		switch choice {
		case "1":
			searchRecipesFlow(reader)
		case "2":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func showMainMenu() {
	fmt.Println("\n")
	fmt.Println("  ____           _                __  __       _       _               ")
	fmt.Println(" |  _ \\ ___  ___(_)_ __   ___   |  \\/  | __ _| |_ ___| |__   ___ _ __ ")
	fmt.Println(" | |_) / _ \\/ __| | '_ \\ / _ \\  | |\\/| |/ _` | __/ __| '_ \\ / _ \\ '__|")
	fmt.Println(" |  _ <  __/ (__| | |_) |  __/  | |  | | (_| | || (__| | | |  __/ |   ")
	fmt.Println(" |_| \\_\\___|\\___|_| .__/ \\___|  |_|  |_|\\__,_|\\__\\__|_| |_|\\___|_|   ")
	fmt.Println("                  |_|                                                  ")
	fmt.Println("")
	fmt.Println(strings.Repeat("=", 72))
	fmt.Println("1. Search for recipes")
	fmt.Println("2. Exit")
	fmt.Println(strings.Repeat("=", 72))
	fmt.Print("\nChoose an option (or 2 to exit): ")
}

func searchRecipesFlow(reader *bufio.Reader) {
	// Get ingredients
	fmt.Println("\nEnter your ingredients separated by comma:")
	fmt.Print("> ")

	input := readInput(reader)

	if input == "exit" || input == "2" {
		return
	}

	ingredients := parseIngredients(input)

	if len(ingredients) == 0 {
		fmt.Println("No ingredients entered!")
		return
	}

	// Search recipes
	fmt.Println("\nSearching for recipes...")
	recipes, err := searchRecipes(ingredients[0])
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	if len(recipes) == 0 {
		fmt.Println("No recipes found. Try different ingredients!")
		return
	}

	// Show recipes and let user browse
	browseRecipes(recipes, reader)
}

func browseRecipes(recipes []Meal, reader *bufio.Reader) {
	for {
		fmt.Printf("\n" + strings.Repeat("-", 60))
		fmt.Printf("\nFound %d recipe(s):\n\n", len(recipes))

		for i, recipe := range recipes {
			fmt.Printf("%d. %s\n", i+1, recipe.Name)
		}

		fmt.Println("\n0. Back to main menu")
		fmt.Print("\nChoose a recipe number (or 0 to exit): ")

		choice := readInput(reader)

		if choice == "0" || choice == "exit" {
			return
		}

		recipeNum, err := strconv.Atoi(choice)
		if err != nil || recipeNum < 1 || recipeNum > len(recipes) {
			fmt.Println("Invalid choice. Please try again.")
			continue
		}

		showRecipeDetails(recipes[recipeNum-1].ID, reader)
	}
}

func showRecipeDetails(mealID string, reader *bufio.Reader) {
	fmt.Println("\nFetching recipe details...")

	recipe, err := getRecipeDetails(mealID)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Printf("RECIPE: %s\n", recipe.Name)
	fmt.Println(strings.Repeat("=", 60))

	fmt.Println("\nINSTRUCTIONS:")
	fmt.Println(recipe.Instructions)

	if recipe.Source != "" {
		fmt.Printf("\nFull recipe: %s\n", recipe.Source)
	}
	fmt.Printf("Image: %s\n", recipe.Thumbnail)

	fmt.Println("\n" + strings.Repeat("-", 60))
	fmt.Println("Press Enter to go back (or type 'exit' to return to main menu)...")

	input := readInput(reader)
	if input == "exit" {
		return
	}
}

func readInput(reader *bufio.Reader) string {
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(strings.ToLower(input))
}

func parseIngredients(input string) []string {
	ingredientList := strings.Split(input, ",")

	var ingredients []string
	for _, ing := range ingredientList {
		cleaned := strings.TrimSpace(strings.ToLower(ing))
		if cleaned != "" {
			ingredients = append(ingredients, cleaned)
		}
	}

	return ingredients
}
