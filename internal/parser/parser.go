package parser

import "github.com/go-jedi/foodgrammm-backend/internal/parser/recipe"

type Parser struct {
	Recipe *recipe.Recipe
}

func NewParser() *Parser {
	return &Parser{
		recipe.NewRecipe(),
	}
}
