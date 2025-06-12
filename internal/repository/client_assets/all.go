package clientassets

import (
	"context"
	"errors"
	"fmt"
	"time"

	clientassets "github.com/go-jedi/foodgramm_backend/internal/domain/client_assets"
)

func (r *repo) All(ctx context.Context) ([]clientassets.ClientAssets, error) {
	r.logger.Debug("[get all client assets] execute repository")

	ctxTimeout, cancel := context.WithTimeout(ctx, time.Duration(r.db.QueryTimeout)*time.Second)
	defer cancel()

	q := `
		SELECT *
		FROM client_assets
		ORDER BY id;
	`

	rows, err := r.db.Pool.Query(ctxTimeout, q)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			r.logger.Error("request timed out while get all client assets", "err", err)
			return nil, fmt.Errorf("the request timed out: %w", err)
		}
		r.logger.Error("failed to get all client assets", "err", err)
		return nil, fmt.Errorf("could not get all client assets: %w", err)
	}
	defer rows.Close()

	var ca []clientassets.ClientAssets

	for rows.Next() {
		var a clientassets.ClientAssets

		if err := rows.Scan(
			&a.ID, &a.NameFile, &a.ServerPathFile,
			&a.ClientPathFile, &a.Extension, &a.Quality,
			&a.OldNameFile, &a.OldExtension, &a.CreatedAt, &a.UpdatedAt,
		); err != nil {
			r.logger.Error("failed to scan row to get all client assets", "err", err)
			return nil, fmt.Errorf("failed to scan row to get all client assets: %w", err)
		}

		ca = append(ca, a)
	}

	if err := rows.Err(); err != nil {
		r.logger.Error("failed to get all client assets", "err", rows.Err())
		return nil, fmt.Errorf("failed to get all client assets: %w", err)
	}

	return ca, nil
}
