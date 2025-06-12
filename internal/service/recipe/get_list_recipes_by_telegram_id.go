package recipe

import (
	"context"

	"github.com/go-jedi/foodgramm_backend/internal/domain/recipe"
	"github.com/go-jedi/foodgramm_backend/pkg/apperrors"
)

func (s *serv) GetListRecipesByTelegramID(ctx context.Context, dto recipe.GetListRecipesByTelegramIDDTO) (recipe.GetListRecipesByTelegramIDResponse, error) {
	s.logger.Debug("[get list recipes by telegram id] execute service")

	ie, err := s.userRepository.ExistsByTelegramID(ctx, dto.TelegramID)
	if err != nil {
		return recipe.GetListRecipesByTelegramIDResponse{}, err
	}

	if !ie {
		return recipe.GetListRecipesByTelegramIDResponse{}, apperrors.ErrUserDoesNotExist
	}

	return s.recipeRepository.GetListRecipesByTelegramID(ctx, dto)
}
