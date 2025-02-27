package recipeevent

import (
	"context"

	recipeevent "github.com/go-jedi/foodgrammm-backend/internal/domain/recipe_event"
)

func (s *serv) GetByID(ctx context.Context, recipeID int64) (recipeevent.Recipe, error) {
	return s.recipeEventRepository.GetByID(ctx, recipeID)
}
