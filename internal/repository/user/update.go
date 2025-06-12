package user

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-jedi/foodgramm_backend/internal/domain/user"
)

func (r *repo) Update(ctx context.Context, dto user.UpdateDTO) (user.User, error) {
	r.logger.Debug("[update user] execute repository")

	ctxTimeout, cancel := context.WithTimeout(ctx, time.Duration(r.db.QueryTimeout)*time.Second)
	defer cancel()

	var uu user.User

	q := `
		UPDATE users SET
			telegram_id = $1,
			username = $2,
			first_name = $3,
			last_name = $4,
		    updated_at = NOW()
		WHERE id = $5
		RETURNING *;
	`

	if err := r.db.Pool.QueryRow(
		ctxTimeout, q,
		dto.TelegramID, dto.Username,
		dto.FirstName, dto.LastName, dto.ID,
	).Scan(
		&uu.ID, &uu.TelegramID, &uu.Username,
		&uu.FirstName, &uu.LastName,
		&uu.CreatedAt, &uu.UpdatedAt,
	); err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			r.logger.Error("request timed out while update user", "err", err)
			return user.User{}, fmt.Errorf("the request timed out: %w", err)
		}
		r.logger.Error("failed to update user", "err", err)
		return user.User{}, fmt.Errorf("could not update user: %w", err)
	}

	return uu, nil
}
