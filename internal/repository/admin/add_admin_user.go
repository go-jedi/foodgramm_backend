package admin

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-jedi/foodgramm_backend/internal/domain/admin"
)

func (r *repo) AddAdminUser(ctx context.Context, telegramID string) (admin.Admin, error) {
	r.logger.Debug("[add a new admin user] execute repository")

	ctxTimeout, cancel := context.WithTimeout(ctx, time.Duration(r.db.QueryTimeout)*time.Second)
	defer cancel()

	q := `
		INSERT INTO admins(
			telegram_id
		) VALUES ($1)
		RETURNING *;
	`

	var na admin.Admin

	if err := r.db.Pool.QueryRow(
		ctxTimeout, q,
		telegramID,
	).Scan(
		&na.ID, &na.TelegramID,
		&na.CreatedAt, &na.UpdatedAt,
	); err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			r.logger.Error("request timed out while add admin user", "err", err)
			return admin.Admin{}, fmt.Errorf("the request timed out: %w", err)
		}
		r.logger.Error("failed to add admin user", "err", err)
		return admin.Admin{}, fmt.Errorf("could not add admin user: %w", err)
	}

	return na, nil
}
