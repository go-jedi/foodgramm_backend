package product

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-jedi/foodgrammm-backend/internal/domain/product"
)

func (r *repo) ExcludeProductsByID(ctx context.Context, dto product.ExcludeProductsByIDDTO) (product.ExcludeProductsByIDResponse, error) {
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
		RETURNING user_id, products;
	`

	var ep product.ExcludeProductsByIDResponse

	if err := r.db.Pool.QueryRow(
		ctxTimeout, q,
		dto.Products, dto.UserID,
	).Scan(
		&ep.UserID, &ep.Products,
	); err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			r.logger.Error("request timed out while exclude products by id", "err", err)
			return product.ExcludeProductsByIDResponse{}, fmt.Errorf("the request timed out: %w", err)
		}
		r.logger.Error("failed to exclude products by id", "err", err)
		return product.ExcludeProductsByIDResponse{}, fmt.Errorf("could not exclude products by id: %w", err)
	}

	return ep, nil
}
