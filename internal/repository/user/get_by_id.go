package user

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-jedi/foodgrammm-backend/internal/domain/user"
)

func (r *repo) GetByID(ctx context.Context, userID int64) (user.User, error) {
	r.logger.Debug("[GetByID] execute repository")

	ctxTimeout, cancel := context.WithTimeout(ctx, time.Duration(r.db.QueryTimeout)*time.Second)
	defer cancel()

	var u user.User

	q := `
		SELECT *
		FROM users
		WHERE id = $1;
	`

	if err := r.db.Pool.QueryRow(
		ctxTimeout, q, userID,
	).Scan(
		&u.ID, &u.TelegramID, &u.Username,
		&u.FirstName, &u.LastName,
		&u.CreatedAt, &u.UpdatedAt,
	); err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			r.logger.Error("request timed out while get user by id", "err", err)
			return user.User{}, fmt.Errorf("the request timed out: %w", err)
		}
		r.logger.Error("failed to get user by id", "err", err)
		return user.User{}, fmt.Errorf("could not get user by id: %w", err)
	}

	return u, nil
}
