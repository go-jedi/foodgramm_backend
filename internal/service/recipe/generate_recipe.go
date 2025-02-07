package recipe

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-jedi/foodgrammm-backend/internal/domain/recipe"
	"github.com/go-jedi/foodgrammm-backend/internal/domain/subscription"
	"github.com/go-jedi/foodgrammm-backend/internal/domain/templates"
	recipetemplate "github.com/go-jedi/foodgrammm-backend/internal/templates/recipe"
	"github.com/go-jedi/foodgrammm-backend/pkg/apperrors"
)

var (
	ErrUserNotSubscriptionOrFreeRecipes = errors.New("user does not have a subscription or free recipes")
)

func (s *serv) GenerateRecipe(ctx context.Context, dto recipe.GenerateRecipeDTO) (recipe.Recipes, error) {
	fmt.Println("dto:", dto)

	// check user exists by telegram id.
	ieu, err := s.userRepository.ExistsByTelegramID(ctx, dto.TelegramID)
	if err != nil {
		return recipe.Recipes{}, err
	}

	if !ieu {
		return recipe.Recipes{}, apperrors.ErrUserDoesNotExist
	}

	// check user subscribed or have free recipes.
	at, err := s.isSubscribedOrFreeRecipes(ctx, dto.TelegramID)
	if err != nil {
		return recipe.Recipes{}, err
	}

	// apply data to need template.
	data, err := s.applyDataToTemplate(dto)
	if err != nil {
		return recipe.Recipes{}, err
	}

	fmt.Println("data:", data)

	// send data for openai service by http request.
	result, err := s.client.OpenAI.Send(ctx, data)
	if err != nil {
		return recipe.Recipes{}, err
	}

	fmt.Println("result:", result)

	// parse recipe from openai.
	parsedRecipe, err := s.parser.Recipe.ParseRecipe(dto.TelegramID, string(result))
	if err != nil {
		return recipe.Recipes{}, err
	}

	// create parsed recipe in database and have free recipe add count and return response.
	return s.recipeRepository.CreateRecipe(ctx, at, parsedRecipe)
}

// isSubscribedOrFreeRecipes check is subscribed or have free recipes.
func (s *serv) isSubscribedOrFreeRecipes(ctx context.Context, telegramID string) (subscription.AccessType, error) {
	ies, err := s.subscriptionRepository.ExistsByTelegramID(ctx, telegramID)
	if err != nil {
		return subscription.NoAccess, err
	}

	if ies {
		return subscription.SubscriptionAccess, nil
	}

	rie, err := s.recipeRepository.ExistsFreeRecipesByTelegramID(ctx, telegramID)
	if err != nil {
		return subscription.NoAccess, err
	}

	if !rie {
		return subscription.NoAccess, ErrUserNotSubscriptionOrFreeRecipes
	}

	return subscription.FreeRecipesAccess, nil
}

// applyDataToTemplate apply data to templates.
func (s *serv) applyDataToTemplate(dto recipe.GenerateRecipeDTO) (templates.ApplyDataToTemplateResponse, error) {
	str, err := s.templates.Recipe.Generate(recipetemplate.GenerateRecipe{
		Type:                  dto.Type,
		Products:              dto.Products,
		NonConsumableProducts: dto.NonConsumableProducts,
		Name:                  dto.Name,
		AmountCalories:        dto.AmountCalories,
		AvailableProducts:     dto.AvailableProducts,
	})
	if err != nil {
		return templates.ApplyDataToTemplateResponse{}, err
	}

	return templates.ApplyDataToTemplateResponse{Content: str}, nil
}
