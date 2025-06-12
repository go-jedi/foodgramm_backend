package templates

import (
	"github.com/go-jedi/foodgramm_backend/internal/templates/recipe"
	recipeofdays "github.com/go-jedi/foodgramm_backend/internal/templates/recipe_of_days"
)

type Templates struct {
	Recipe       *recipe.Template
	RecipeOfDays *recipeofdays.Template
}

func NewTemplates() *Templates {
	return &Templates{
		Recipe:       recipe.NewTemplate(),
		RecipeOfDays: recipeofdays.NewTemplate(),
	}
}
