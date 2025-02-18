package recipeofdays

import (
	"context"
)

func (s *serv) Create(ctx context.Context) error {
	s.logger.Debug("[create recipe of the day] execute service")

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

	// parse recipe from openai.
	parsedRecipe, err := s.parser.RecipeOfDays.ParseRecipe(string(result))
	if err != nil {
		return err
	}

	// create parsed recipe in database.
	return s.recipeOfDaysRepository.Create(ctx, parsedRecipe)
}
