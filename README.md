# Recipe Matcher

A command-line tool that helps you find recipes based on ingredients you have at home.

**Note:** This is a practice project I built to learn Go programming and work with APIs. :)

## Features

- Search for recipes by ingredient
- Browse recipe results
- View detailed cooking instructions
- Free API - no registration required

## Prerequisites

- Go 1.16 or higher installed
- Internet connection (to fetch recipes)

## Installation

1. Clone this repository:
```bash
git clone https://github.com/darthtrish/recipe-matcher.git
```

2. Navigate to the project folder:
```bash
cd recipe-matcher
```

3. Run the program:
```bash
go run .
```

## Usage

1. Start the program
2. Choose option 1 to search for recipes
3. Enter ingredients separated by commas (e.g., `chicken, rice, tomato`)
4. Browse the recipe list
5. Select a recipe number to view full instructions
6. Press 0 to go back or 2 to exit

## Example
```
Enter your ingredients separated by comma:
> chicken, garlic

Found 12 recipe(s):

1. Garlic Chicken
2. Chicken Alfredo
3. Lemon Garlic Chicken
...

Choose a recipe number to view: 1
```

## Technologies Used

- **Go** - Programming language
- **TheMealDB API** - Free recipe database

## Project Structure
```
recipe-matcher/
├── main.go      # Main program logic and menu
├── api.go       # API calls to TheMealDB
├── types.go     # Data structures
├── go.mod       # Go module file
└── README.md    # This file
```

## What I Learned

- Go basics: loops, functions, structs
- Reading user input with bufio
- Making HTTP requests and working with APIs
- Parsing JSON responses
- Error handling in Go
- Organizing code into multiple files
- Git version control and GitHub

## Future Plans

- [ ] Create a web version with HTML/CSS frontend
- [ ] Add ability to search by multiple ingredients simultaneously
- [ ] Parse grocery store sales PDFs (Lidl, S-market) to suggest recipes based on discounted items

## License

This project is open source and available for anyone to use and learn from.

## Acknowledgments

- Recipe data provided by [TheMealDB](https://www.themealdb.com/)
- Built as a learning project to practice Go programming and API integration
