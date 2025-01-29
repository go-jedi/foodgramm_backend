package product

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-jedi/foodgrammm-backend/internal/domain/product"
)

func (r *repo) GetExcludeProductsByUserID(ctx context.Context, userID int64) (product.UserExcludedProducts, error) {
	ctxTimeout, cancel := context.WithTimeout(ctx, time.Duration(r.db.QueryTimeout)*time.Second)
	defer cancel()

	var uep product.UserExcludedProducts

	q := `
		SELECT *
		FROM user_excluded_products_table
		WHERE user_id = $1;
	`

	if err := r.db.Pool.QueryRow(ctxTimeout, q, userID).Scan(
		&uep.ID, &uep.UserID, &uep.TelegramID,
		&uep.Products, &uep.CreatedAt, &uep.UpdatedAt,
	); err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			r.logger.Error("request timed out while get exclude products by user id", "err", err)
			return product.UserExcludedProducts{}, fmt.Errorf("the request timed out: %w", err)
		}
		r.logger.Error("failed to get exclude products by user id", "err", err)
		return product.UserExcludedProducts{}, fmt.Errorf("could not get exclude products by user id: %w", err)
	}

	return uep, nil
}
