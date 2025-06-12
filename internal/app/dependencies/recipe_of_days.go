package dependencies

import (
	"context"

	recipeofdayscron "github.com/go-jedi/foodgramm_backend/internal/adapters/cron/recipe_of_days"
	recipeofdays "github.com/go-jedi/foodgramm_backend/internal/adapters/http/handlers/recipe_of_days"
	"github.com/go-jedi/foodgramm_backend/internal/repository"
	recipeOfDaysRepository "github.com/go-jedi/foodgramm_backend/internal/repository/recipe_of_days"
	"github.com/go-jedi/foodgramm_backend/internal/service"
	recipeOfDaysService "github.com/go-jedi/foodgramm_backend/internal/service/recipe_of_days"
)

func (d *Dependencies) RecipeOfDaysRepository() repository.RecipeOfDaysRepository {
	if d.recipeOfDaysRepository == nil {
		d.recipeOfDaysRepository = recipeOfDaysRepository.NewRepository(
			d.logger,
			d.db,
		)
	}

	return d.recipeOfDaysRepository
}

func (d *Dependencies) RecipeOfDaysService() service.RecipeOfDaysService {
	if d.recipeOfDaysService == nil {
		d.recipeOfDaysService = recipeOfDaysService.NewService(
			d.RecipeOfDaysRepository(),
			d.client,
			d.templates,
			d.parser,
			d.logger,
			d.cache,
		)
	}

	return d.recipeOfDaysService
}

func (d *Dependencies) RecipeOfDaysHandler() *recipeofdays.Handler {
	if d.recipeOfDaysHandler == nil {
		d.recipeOfDaysHandler = recipeofdays.NewHandler(
			d.RecipeOfDaysService(),
			d.cookie,
			d.middleware,
			d.engine,
			d.logger,
			d.validator,
		)
	}

	return d.recipeOfDaysHandler
}

func (d *Dependencies) RecipeOfDaysCron(ctx context.Context) *recipeofdayscron.Cron {
	if d.recipeOfDaysCron == nil {
		d.recipeOfDaysCron = recipeofdayscron.NewCron(
			ctx,
			d.RecipeOfDaysService(),
			d.worker,
			d.logger,
		)
	}

	return d.recipeOfDaysCron
}
