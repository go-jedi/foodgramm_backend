package user

import (
	"context"
	"errors"
	"time"

	"github.com/go-jedi/foodgrammm-backend/internal/domain/user"
	"github.com/go-jedi/foodgrammm-backend/pkg/apperrors"
)

func (r *repo) Create(ctx context.Context, dto user.CreateDTO) (user.User, error) {
	ctxTimeout, cancel := context.WithTimeout(ctx, time.Duration(r.db.QueryTimeout)*time.Second)
	defer cancel()

	var nu user.User

	q := `
		INSERT INTO users(
			telegram_id,
		    username,
		    first_name,
		    last_name
		) VALUES($1, $2, $3, $4)
		RETURNING *;
	`

	if err := r.db.Pool.QueryRow(ctxTimeout, q, dto.TelegramID, dto.Username, dto.FirstName, dto.LastName).Scan(
		&nu.ID, &nu.TelegramID, &nu.Username, &nu.FirstName,
		&nu.LastName, &nu.CreatedAt, &nu.UpdatedAt,
	); err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return user.User{}, apperrors.ErrRequestTimedOut
		}
		return user.User{}, err
	}

	return nu, nil
}
