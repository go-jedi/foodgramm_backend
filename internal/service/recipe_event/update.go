package recipeevent

import (
	"context"

	recipeevent "github.com/go-jedi/foodgrammm-backend/internal/domain/recipe_event"
	"github.com/go-jedi/foodgrammm-backend/pkg/apperrors"
)

func (s *serv) Update(ctx context.Context, dto recipeevent.UpdateDTO) (recipeevent.Recipe, error) {
	ie, err := s.recipeTypesRepository.ExistsByRecipeTypeID(ctx, dto.TypeID)
	if err != nil {
		return recipeevent.Recipe{}, err
	}

	if !ie {
		return recipeevent.Recipe{}, apperrors.ErrRecipeTypeDoesNotExists
	}

	return s.recipeEventRepository.Update(ctx, dto)
}
