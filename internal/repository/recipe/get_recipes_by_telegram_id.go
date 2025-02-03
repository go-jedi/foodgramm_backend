package recipe

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-jedi/foodgrammm-backend/internal/domain/recipe"
)

func (r *repo) GetRecipesByTelegramID(ctx context.Context, telegramID string) ([]recipe.Recipes, error) {
	ctxTimeout, cancel := context.WithTimeout(ctx, time.Duration(r.db.QueryTimeout)*time.Second)
	defer cancel()

	q := `
		SELECT *
		FROM recipes
		WHERE telegram_id = $1
		ORDER BY id;
	`

	rows, err := r.db.Pool.Query(ctxTimeout, q, telegramID)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			r.logger.Error("request timed out while get recipes by telegram id", "err", err)
			return nil, fmt.Errorf("the request timed out: %w", err)
		}
		r.logger.Error("failed to get recipes by telegram id", "err", err)
		return nil, fmt.Errorf("could not get recipes by telegram id: %w", err)
	}
	defer rows.Close()

	var rsu []recipe.Recipes

	for rows.Next() {
		var rs recipe.Recipes

		if err := rows.Scan(
			&rs.ID, &rs.TelegramID, &rs.Title,
			&rs.Content, &rs.CreatedAt, &rs.UpdatedAt,
		); err != nil {
			r.logger.Error("failed to scan row to get recipes by telegram id", "err", err)
			return nil, fmt.Errorf("failed to scan row to get recipes by telegram id: %w", err)
		}

		rsu = append(rsu, rs)
	}

	if err := rows.Err(); err != nil {
		r.logger.Error("failed to get recipes by telegram id", "err", rows.Err())
		return nil, fmt.Errorf("failed to get recipes by telegram id: %w", err)
	}

	return rsu, nil
}
