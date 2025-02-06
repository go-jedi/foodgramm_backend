package recipe

import (
	"context"
	"errors"
	"time"

	"github.com/go-jedi/foodgrammm-backend/internal/domain/recipe"
	"github.com/go-jedi/foodgrammm-backend/internal/domain/subscription"
	"github.com/go-jedi/foodgrammm-backend/internal/domain/templates"
	recipetemplate "github.com/go-jedi/foodgrammm-backend/internal/templates/recipe"
	"github.com/go-jedi/foodgrammm-backend/pkg/apperrors"
)

var (
	ErrUserNotSubscriptionOrFreeRecipes = errors.New("user does not have a subscription or free recipes")
)

func (s *serv) GenerateRecipe(ctx context.Context, dto recipe.GenerateRecipeDTO) (recipe.GenerateRecipeResponse, error) {
	// check user exists by telegram id.
	ieu, err := s.userRepository.ExistsByTelegramID(ctx, dto.TelegramID)
	if err != nil {
		return recipe.GenerateRecipeResponse{}, err
	}

	if !ieu {
		return recipe.GenerateRecipeResponse{}, apperrors.ErrUserDoesNotExist
	}

	// check user subscribed or have free recipes.
	at, err := s.isSubscribedOrFreeRecipes(ctx, dto.TelegramID)
	if err != nil {
		return recipe.GenerateRecipeResponse{}, err
	}

	// apply data to need template.
	data, err := s.applyDataToTemplate(dto)
	if err != nil {
		return recipe.GenerateRecipeResponse{}, err
	}

	// send data for openai service by http request.
	result, err := s.client.OpenAI.Send(ctx, data)
	if err != nil {
		return recipe.GenerateRecipeResponse{}, err
	}

	// parse data from openai.
	parsedData, err := s.parser.Recipe.ParseRecipe(dto.TelegramID, string(result))
	if err != nil {
		return recipe.GenerateRecipeResponse{}, err
	}

	// save parsed data in cache.
	if err := s.saveParsedDataToCache(ctx, dto.TelegramID, parsedData); err != nil {
		return recipe.GenerateRecipeResponse{}, err
	}

	// have free recipes add count.
	if at == subscription.FreeRecipesAccess {
		if _, err := s.recipeRepository.AddFreeRecipesCountByTelegramID(ctx, dto.TelegramID); err != nil {
			return recipe.GenerateRecipeResponse{}, err
		}
	}

	return parsedData, nil
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

// saveParsedDataToCache save parsed data to cache.
func (s *serv) saveParsedDataToCache(ctx context.Context, telegramID string, parsedData recipe.GenerateRecipeResponse) error {
	const expirationMin = 60

	ic := recipe.InCache{
		TelegramID: telegramID,
		Title:      parsedData.Title,
		Content:    parsedData.Content,
		CreatedAt:  parsedData.CreatedAt,
		UpdatedAt:  parsedData.UpdatedAt,
	}

	if err := s.cache.Recipe.Set(
		ctx,
		telegramID,
		ic,
		time.Duration(expirationMin)*time.Minute,
	); err != nil {
		return err
	}

	return nil
}
