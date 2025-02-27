package recipeevent

import (
	"context"

	recipeevent "github.com/go-jedi/foodgrammm-backend/internal/domain/recipe_event"
	"github.com/go-jedi/foodgrammm-backend/pkg/apperrors"
)

func (s *serv) Create(ctx context.Context, dto recipeevent.CreateDTO) (recipeevent.Recipe, error) {
	ie, err := s.recipeTypesRepository.ExistsByRecipeTypeID(ctx, dto.TypeID)
	if err != nil {
		return recipeevent.Recipe{}, err
	}

	if !ie {
		return recipeevent.Recipe{}, apperrors.ErrRecipeTypeDoesNotExists
	}

	return s.recipeEventRepository.Create(ctx, dto)
}
