package user

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-jedi/foodgrammm-backend/internal/domain/user"
)

func (r *repo) GetByTelegramID(ctx context.Context, telegramID string) (user.User, error) {
	ctxTimeout, cancel := context.WithTimeout(ctx, time.Duration(r.db.QueryTimeout)*time.Second)
	defer cancel()

	var u user.User

	q := `
		SELECT *
		FROM users
		WHERE telegram_id = $1;
	`

	if err := r.db.Pool.QueryRow(
		ctxTimeout, q, telegramID,
	).Scan(
		&u.ID, &u.TelegramID, &u.Username,
		&u.FirstName, &u.LastName,
		&u.CreatedAt, &u.UpdatedAt,
	); err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			r.logger.Error("request timed out while get user by telegramID", "err", err)
			return user.User{}, fmt.Errorf("the request timed out: %w", err)
		}
		r.logger.Error("failed to get user by telegramID", "err", err)
		return user.User{}, fmt.Errorf("could not get user by telegramID: %w", err)
	}

	return u, nil
}
