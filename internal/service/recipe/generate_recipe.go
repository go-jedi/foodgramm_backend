package recipe

import (
	"context"
	"errors"
	"strings"

	"github.com/go-jedi/foodgrammm-backend/internal/domain/parser"
	"github.com/go-jedi/foodgrammm-backend/internal/domain/recipe"
	recipescraper "github.com/go-jedi/foodgrammm-backend/internal/domain/recipe_scraper"
	"github.com/go-jedi/foodgrammm-backend/internal/domain/subscription"
	"github.com/go-jedi/foodgrammm-backend/internal/domain/templates"
	recipetemplate "github.com/go-jedi/foodgrammm-backend/internal/templates/recipe"
	"github.com/go-jedi/foodgrammm-backend/pkg/apperrors"
)

var ErrUserNotSubscriptionOrFreeRecipes = errors.New("user does not have a subscription or free recipes")

func (s *serv) GenerateRecipe(ctx context.Context, dto recipe.GenerateRecipeDTO) (recipe.Recipes, error) {
	s.logger.Debug("[generate recipe] execute service")

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

	var pr parser.ParsedRecipe

	switch dto.Type {
	case 1:
		// generate recipe with scraper.
		pr, err = s.generateScraper(ctx, dto)
		if err != nil {
			return recipe.Recipes{}, err
		}
	default:
		// generate recipe with AI.
		pr, err = s.generateAI(ctx, dto)
		if err != nil {
			return recipe.Recipes{}, err
		}
	}

	// create parsed recipe in database and have free recipe add count and return response.
	return s.recipeRepository.CreateRecipe(ctx, at, pr)
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

// generateScraper generate recipe with scraper.
func (s *serv) generateScraper(ctx context.Context, data recipe.GenerateRecipeDTO) (parser.ParsedRecipe, error) {
	body := recipescraper.GetBody{
		TelegramID: data.TelegramID,
		Type:       data.Type,
	}

	if data.NonConsumableProducts != nil {
		body.NonConsumableProducts = strings.Split(*data.NonConsumableProducts, ", ")
	}

	if len(data.Products) > 0 {
		body.NonConsumableProducts = append(body.NonConsumableProducts, data.Products...)
	}

	return s.client.RecipeScraper.Get(ctx, body)
}

// generateAI generate recipe with ai.
func (s *serv) generateAI(ctx context.Context, data recipe.GenerateRecipeDTO) (parser.ParsedRecipe, error) {
	// apply data to need template.
	t, err := s.applyDataToTemplate(data)
	if err != nil {
		return parser.ParsedRecipe{}, err
	}

	// send data for openai service by http request.
	result, err := s.client.OpenAI.Send(ctx, t)
	if err != nil {
		return parser.ParsedRecipe{}, err
	}

	// parse recipe from openai.
	return s.parser.Recipe.ParseRecipe(data.TelegramID, string(result))
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
