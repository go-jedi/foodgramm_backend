package user

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-jedi/foodgramm_backend/internal/domain/user"
	jsoniter "github.com/json-iterator/go"
)

func (r *repo) Create(ctx context.Context, dto user.CreateDTO) (user.User, error) {
	r.logger.Debug("[create a new user] execute repository")

	ctxTimeout, cancel := context.WithTimeout(ctx, time.Duration(r.db.QueryTimeout)*time.Second)
	defer cancel()

	q := `SELECT * FROM public.user_create($1);`

	rawData, err := jsoniter.Marshal(dto)
	if err != nil {
		r.logger.Error("failed to marshal user data", "err", err)
		return user.User{}, fmt.Errorf("could not marshal user data: %w", err)
	}

	var nu user.User

	if err := r.db.Pool.QueryRow(
		ctxTimeout, q,
		rawData,
	).Scan(
		&nu.ID, &nu.TelegramID, &nu.Username,
		&nu.FirstName, &nu.LastName,
		&nu.CreatedAt, &nu.UpdatedAt,
	); err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			r.logger.Error("request timed out while creating the user", "err", err)
			return user.User{}, fmt.Errorf("the request timed out: %w", err)
		}
		r.logger.Error("failed to create user", "err", err)
		return user.User{}, fmt.Errorf("could not create user: %w", err)
	}

	return nu, nil
}
