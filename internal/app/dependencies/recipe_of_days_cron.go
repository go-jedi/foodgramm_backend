package dependencies

import (
	"context"

	recipeofdays "github.com/go-jedi/foodgrammm-backend/internal/adapters/http/cron/recipe_of_days"
)

func (d *Dependencies) RecipeOfDaysCron(ctx context.Context) *recipeofdays.Cron {
	if d.recipeOfDaysCron == nil {
		d.recipeOfDaysCron = recipeofdays.NewCron(
			ctx,
			d.RecipeOfDaysService(),
			d.worker,
			d.logger,
		)
	}

	return d.recipeOfDaysCron
}
