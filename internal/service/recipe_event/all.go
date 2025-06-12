package recipeevent

import (
	"context"

	recipeevent "github.com/go-jedi/foodgramm_backend/internal/domain/recipe_event"
)

func (s *serv) All(ctx context.Context) ([]recipeevent.Recipe, error) {
	return s.recipeEventRepository.All(ctx)
}
