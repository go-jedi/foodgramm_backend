package dependencies

import (
	recipetypes "github.com/go-jedi/foodgrammm-backend/internal/adapters/http/handlers/recipe_types"
	"github.com/go-jedi/foodgrammm-backend/internal/repository"
	recipeTypesRepository "github.com/go-jedi/foodgrammm-backend/internal/repository/recipe_types"
	"github.com/go-jedi/foodgrammm-backend/internal/service"
	recipeTypesService "github.com/go-jedi/foodgrammm-backend/internal/service/recipe_types"
)

func (d *Dependencies) RecipeTypesRepository() repository.RecipeTypesRepository {
	if d.recipeTypesRepository == nil {
		d.recipeTypesRepository = recipeTypesRepository.NewRepository(
			d.logger,
			d.db,
		)
	}

	return d.recipeTypesRepository
}

func (d *Dependencies) RecipeTypesService() service.RecipeTypesService {
	if d.recipeTypesService == nil {
		d.recipeTypesService = recipeTypesService.NewService(
			d.RecipeTypesRepository(),
			d.logger,
			d.cache,
		)
	}

	return d.recipeTypesService
}

func (d *Dependencies) RecipeTypesHandler() *recipetypes.Handler {
	if d.recipeTypesHandler == nil {
		d.recipeTypesHandler = recipetypes.NewHandler(
			d.RecipeTypesService(),
			d.cookie,
			d.middleware,
			d.engine,
			d.logger,
			d.validator,
		)
	}

	return d.recipeTypesHandler
}
