package recipe

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-jedi/foodgrammm-backend/internal/domain/parser"
	"github.com/go-jedi/foodgrammm-backend/internal/domain/recipe"
	"github.com/go-jedi/foodgrammm-backend/internal/domain/subscription"
	jsoniter "github.com/json-iterator/go"
)

func (r *repo) CreateRecipe(ctx context.Context, accessType subscription.AccessType, parsedData parser.ParsedRecipe) (recipe.Recipes, error) {
	ctxTimeout, cancel := context.WithTimeout(ctx, time.Duration(r.db.QueryTimeout)*time.Second)
	defer cancel()

	q := `SELECT * FROM public.recipe_create($1, $2);`

	rawData, err := jsoniter.Marshal(parsedData)
	if err != nil {
		return recipe.Recipes{}, err
	}

	var rc recipe.Recipes

	if err := r.db.Pool.QueryRow(
		ctxTimeout, q,
		accessType, rawData,
	).Scan(
		&rc.ID, &rc.TelegramID, &rc.Title,
		&rc.Content, &rc.CreatedAt, &rc.UpdatedAt,
	); err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			r.logger.Error("request timed out while creating the recipe", "err", err)
			return recipe.Recipes{}, fmt.Errorf("the request timed out: %w", err)
		}
		r.logger.Error("failed to create recipe", "err", err)
		return recipe.Recipes{}, fmt.Errorf("could not create recipe: %w", err)
	}

	return rc, nil
}
