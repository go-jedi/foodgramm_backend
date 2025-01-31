package product

import (
	"context"

	"github.com/go-jedi/foodgrammm-backend/internal/domain/product"
	"github.com/go-jedi/foodgrammm-backend/pkg/apperrors"
)

func (s *serv) GetExcludeProductsByUserID(ctx context.Context, userID int64) (product.UserExcludedProducts, error) {
	ie, err := s.userRepository.ExistsByID(ctx, userID)
	if err != nil {
		return product.UserExcludedProducts{}, err
	}

	if !ie {
		return product.UserExcludedProducts{}, apperrors.ErrUserDoesNotExist
	}

	return s.productRepository.GetExcludeProductsByUserID(ctx, userID)
}
