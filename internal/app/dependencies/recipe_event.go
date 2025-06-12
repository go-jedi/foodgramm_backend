package dependencies

import (
	recipeevent "github.com/go-jedi/foodgramm_backend/internal/adapters/http/handlers/recipe_event"
	"github.com/go-jedi/foodgramm_backend/internal/repository"
	recipeEventRepository "github.com/go-jedi/foodgramm_backend/internal/repository/recipe_event"
	"github.com/go-jedi/foodgramm_backend/internal/service"
	recipeEventService "github.com/go-jedi/foodgramm_backend/internal/service/recipe_event"
)

func (d *Dependencies) RecipeEventRepository() repository.RecipeEventRepository {
	if d.recipeEventRepository == nil {
		d.recipeEventRepository = recipeEventRepository.NewRepository(
			d.logger,
			d.db,
		)
	}

	return d.recipeEventRepository
}

func (d *Dependencies) RecipeEventService() service.RecipeEventService {
	if d.recipeEventService == nil {
		d.recipeEventService = recipeEventService.NewService(
			d.RecipeEventRepository(),
			d.RecipeTypesRepository(),
			d.logger,
			d.cache,
		)
	}

	return d.recipeEventService
}

func (d *Dependencies) RecipeEventHandler() *recipeevent.Handler {
	if d.recipeEventHandler == nil {
		d.recipeEventHandler = recipeevent.NewHandler(
			d.RecipeEventService(),
			d.cookie,
			d.middleware,
			d.engine,
			d.logger,
			d.validator,
		)
	}

	return d.recipeEventHandler
}
