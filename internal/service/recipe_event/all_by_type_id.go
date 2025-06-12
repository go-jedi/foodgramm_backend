package recipeevent

import (
	"context"

	recipeevent "github.com/go-jedi/foodgramm_backend/internal/domain/recipe_event"
)

func (s *serv) AllByTypeID(ctx context.Context, typeID int64) ([]recipeevent.Recipe, error) {
	return s.recipeEventRepository.AllByTypeID(ctx, typeID)
}
