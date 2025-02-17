package recipe

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-jedi/foodgrammm-backend/internal/domain/recipe"
)

func (r *repo) GetListRecipesByTelegramID(ctx context.Context, dto recipe.GetListRecipesByTelegramIDDTO) (recipe.GetListRecipesByTelegramIDResponse, error) {
	r.logger.Debug("[GetListRecipesByTelegramID] execute repository")

	ctxTimeout, cancel := context.WithTimeout(ctx, time.Duration(r.db.QueryTimeout)*time.Second)
	defer cancel()

	var glr recipe.GetListRecipesByTelegramIDResponse

	q := `
		WITH total AS (
			SELECT COUNT(*) AS total_count
			FROM recipes
			WHERE telegram_id = $1
		),
		paged_data AS (
			SELECT *
			FROM recipes
			WHERE telegram_id = $1
			ORDER BY id
			LIMIT $2 OFFSET ($3 - 1) * $2
		)
		SELECT
			(SELECT total_count FROM total) AS total_count,
			CEIL((SELECT total_count FROM total)::DECIMAL / $2) AS total_pages,
			$3 AS current_page,
			$2 AS size,
			json_agg(paged_data) AS data
		FROM paged_data;
	`

	if err := r.db.Pool.QueryRow(
		ctxTimeout, q,
		dto.TelegramID, dto.Size, dto.Page,
	).Scan(
		&glr.TotalCount, &glr.TotalPages,
		&glr.CurrentPage, &glr.Size, &glr.Data,
	); err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			r.logger.Error("request timed out while get list recipes", "err", err)
			return recipe.GetListRecipesByTelegramIDResponse{}, fmt.Errorf("the request timed out: %w", err)
		}
		r.logger.Error("failed to get list recipes", "err", err)
		return recipe.GetListRecipesByTelegramIDResponse{}, fmt.Errorf("could not get list recipes: %w", err)
	}

	return glr, nil
}
