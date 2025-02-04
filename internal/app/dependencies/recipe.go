package dependencies

import (
	"github.com/go-jedi/foodgrammm-backend/internal/adapters/http/handlers/recipe"
	"github.com/go-jedi/foodgrammm-backend/internal/repository"
	recipeRepository "github.com/go-jedi/foodgrammm-backend/internal/repository/recipe"
	"github.com/go-jedi/foodgrammm-backend/internal/service"
	recipeService "github.com/go-jedi/foodgrammm-backend/internal/service/recipe"
)

func (d *Dependencies) RecipeRepository() repository.RecipeRepository {
	if d.recipeRepository == nil {
		d.recipeRepository = recipeRepository.NewRepository(
			d.logger,
			d.db,
		)
	}

	return d.recipeRepository
}

func (d *Dependencies) RecipeService() service.RecipeService {
	if d.recipeService == nil {
		d.recipeService = recipeService.NewService(
			d.RecipeRepository(),
			d.UserRepository(),
			d.ProductRepository(),
			d.SubscriptionRepository(),
			d.client,
			d.templates,
			d.logger,
			d.cache,
		)
	}

	return d.recipeService
}

func (d *Dependencies) RecipeHandler() *recipe.Handler {
	if d.recipeHandler == nil {
		d.recipeHandler = recipe.NewHandler(
			d.RecipeService(),
			d.cookie,
			d.middleware,
			d.engine,
			d.logger,
			d.validator,
		)
	}

	return d.recipeHandler
}
