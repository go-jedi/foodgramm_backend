package product

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-jedi/foodgrammm-backend/internal/domain/product"
)

func (r *repo) AddExcludeProductsByUserID(ctx context.Context, dto product.AddExcludeProductsByUserIDDTO) (product.UserExcludedProducts, error) {
	ctxTimeout, cancel := context.WithTimeout(ctx, time.Duration(r.db.QueryTimeout)*time.Second)
	defer cancel()

	q := `
		UPDATE user_excluded_products_table SET
			products = ARRAY_CAT(products, (
            	SELECT ARRAY_AGG(DISTINCT product)
            	FROM UNNEST($1::TEXT[]) AS product
            	WHERE product NOT IN (SELECT UNNEST(products))
        	)
		)
        WHERE user_id = $2
		RETURNING *;
	`

	var uep product.UserExcludedProducts

	if err := r.db.Pool.QueryRow(
		ctxTimeout, q,
		dto.Products, dto.UserID,
	).Scan(
		&uep.ID, &uep.UserID, &uep.TelegramID,
		&uep.Products, &uep.CreatedAt, &uep.UpdatedAt,
	); err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			r.logger.Error("request timed out while add exclude products by user id", "err", err)
			return product.UserExcludedProducts{}, fmt.Errorf("the request timed out: %w", err)
		}
		r.logger.Error("failed to add exclude products by user id", "err", err)
		return product.UserExcludedProducts{}, fmt.Errorf("could not add exclude products by user id: %w", err)
	}

	return uep, nil
}
