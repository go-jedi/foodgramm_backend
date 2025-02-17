package recipe

import (
	"context"

	"github.com/go-jedi/foodgrammm-backend/internal/domain/recipe"
	"github.com/go-jedi/foodgrammm-backend/pkg/apperrors"
)

func (s *serv) GetListRecipesByTelegramID(ctx context.Context, dto recipe.GetListRecipesByTelegramIDDTO) (recipe.GetListRecipesByTelegramIDResponse, error) {
	s.logger.Debug("[GetListRecipesByTelegramID] execute service")

	ie, err := s.userRepository.ExistsByTelegramID(ctx, dto.TelegramID)
	if err != nil {
		return recipe.GetListRecipesByTelegramIDResponse{}, err
	}

	if !ie {
		return recipe.GetListRecipesByTelegramIDResponse{}, apperrors.ErrUserDoesNotExist
	}

	return s.recipeRepository.GetListRecipesByTelegramID(ctx, dto)
}
