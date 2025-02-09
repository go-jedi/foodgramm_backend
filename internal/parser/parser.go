package parser

import (
	"github.com/go-jedi/foodgrammm-backend/internal/parser/recipe"
	recipeofdays "github.com/go-jedi/foodgrammm-backend/internal/parser/recipe_of_days"
)

type Parser struct {
	Recipe       *recipe.Parser
	RecipeOfDays *recipeofdays.Parser
}

func NewParser() *Parser {
	return &Parser{
		Recipe:       recipe.NewRecipe(),
		RecipeOfDays: recipeofdays.NewRecipe(),
	}
}
