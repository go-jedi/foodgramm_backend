package recipeofdays

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-jedi/foodgrammm-backend/internal/domain/parser"
	"github.com/go-jedi/foodgrammm-backend/pkg/apperrors"
)

func (r *repo) Create(ctx context.Context, data parser.ParsedRecipeOfDays) error {
	ctxTimeout, cancel := context.WithTimeout(ctx, time.Duration(r.db.QueryTimeout)*time.Second)
	defer cancel()

	q := `
		INSERT INTO recipe_of_days(
		    title,
		    content,
		    description
		) VALUES ($1, $2, $3);
	`

	commandTag, err := r.db.Pool.Exec(
		ctxTimeout, q,
		data.Title, data.Content, data.Description,
	)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			r.logger.Error("request timed out while create recipe of days", "err", err)
			return fmt.Errorf("the request timed out: %w", err)
		}
		r.logger.Error("failed to create recipe of days", "err", err)
		return fmt.Errorf("could not create recipe of days: %w", err)
	}

	if commandTag.RowsAffected() == 0 {
		return apperrors.ErrNoRowsWereAffected
	}

	return nil
}
