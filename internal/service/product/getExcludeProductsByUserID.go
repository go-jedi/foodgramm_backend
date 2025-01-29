package product

import (
	"context"

	"github.com/go-jedi/foodgrammm-backend/internal/domain/product"
)

func (s *serv) GetExcludeProductsByUserID(ctx context.Context, userID int64) (product.UserExcludedProducts, error) {
	return s.productRepository.GetExcludeProductsByUserID(ctx, userID)
}
