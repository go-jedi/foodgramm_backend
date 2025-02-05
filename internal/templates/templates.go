package templates

import "github.com/go-jedi/foodgrammm-backend/internal/templates/recipe"

type Templates struct {
	Recipe *recipe.Template
}

func NewTemplates() *Templates {
	return &Templates{
		Recipe: recipe.NewRecipe(),
	}
}
