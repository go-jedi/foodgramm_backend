package recipeofdays

import (
	"context"
	"fmt"

	"github.com/go-jedi/foodgrammm-backend/internal/domain/parser"
)

func (s *serv) Create(ctx context.Context) error {
	// get template.
	str, err := s.templates.RecipeOfDays.Generate()
	if err != nil {
		return err
	}

	// send data for openai service by http request.
	result, err := s.client.OpenAI.Send(ctx, str)
	if err != nil {
		return err
	}

	fmt.Println("result:", result)

	return s.recipeOfDaysRepository.Create(ctx, parser.ParsedRecipeOfDays{})
}
