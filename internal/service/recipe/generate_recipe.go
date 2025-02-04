package recipe

import (
	"context"
	"fmt"

	"github.com/go-jedi/foodgrammm-backend/internal/domain/recipe"
)

func (s *serv) GenerateRecipe(_ context.Context, dto recipe.GenerateRecipeDTO) error {
	fmt.Println("dto:", dto)
	return nil
}
