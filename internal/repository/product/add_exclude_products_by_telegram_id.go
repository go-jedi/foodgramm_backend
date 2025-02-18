package product

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-jedi/foodgrammm-backend/internal/domain/product"
)

func (r *repo) AddExcludeProductsByTelegramID(ctx context.Context, dto product.AddExcludeProductsByTelegramIDDTO) (product.UserExcludedProducts, error) {
	r.logger.Debug("[add exclude products by telegram id] execute repository")

	ctxTimeout, cancel := context.WithTimeout(ctx, time.Duration(r.db.QueryTimeout)*time.Second)
	defer cancel()

	q := `
		UPDATE user_excluded_products SET
			products = $1 
        WHERE telegram_id = $2
		RETURNING *;
	`

	var uep product.UserExcludedProducts

	if err := r.db.Pool.QueryRow(
		ctxTimeout, q,
		dto.Products, dto.TelegramID,
	).Scan(
		&uep.ID, &uep.UserID, &uep.TelegramID,
		&uep.Products, &uep.CreatedAt, &uep.UpdatedAt,
	); err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			r.logger.Error("request timed out while add exclude products by telegram id", "err", err)
			return product.UserExcludedProducts{}, fmt.Errorf("the request timed out: %w", err)
		}
		r.logger.Error("failed to add exclude products by telegram id", "err", err)
		return product.UserExcludedProducts{}, fmt.Errorf("could not add exclude products by telegram id: %w", err)
	}

	return uep, nil
}
