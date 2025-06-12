package user

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-jedi/foodgramm_backend/internal/domain/user"
)

func (r *repo) List(ctx context.Context, dto user.ListDTO) (user.ListResponse, error) {
	r.logger.Debug("[get list users with pagination] execute repository")

	ctxTimeout, cancel := context.WithTimeout(ctx, time.Duration(r.db.QueryTimeout)*time.Second)
	defer cancel()

	var lr user.ListResponse

	q := `
		WITH total AS (
			SELECT COUNT(*) AS total_count
			FROM users
		),
		paged_data AS (
			SELECT *
			FROM users
			ORDER BY id
			LIMIT $1 OFFSET ($2 - 1) * $1
		)
		SELECT
			(SELECT total_count FROM total) AS total_count,
			CEIL((SELECT total_count FROM total)::DECIMAL / $1) AS total_pages,
			$2 AS current_page,
			$1 AS size,
			json_agg(paged_data) AS data
		FROM paged_data;
	`

	if err := r.db.Pool.QueryRow(
		ctxTimeout, q,
		dto.Size, dto.Page,
	).Scan(
		&lr.TotalCount, &lr.TotalPages,
		&lr.CurrentPage, &lr.Size, &lr.Data,
	); err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			r.logger.Error("request timed out while get list users", "err", err)
			return user.ListResponse{}, fmt.Errorf("the request timed out: %w", err)
		}
		r.logger.Error("failed to get list users", "err", err)
		return user.ListResponse{}, fmt.Errorf("could not get list users: %w", err)
	}

	return lr, nil
}
