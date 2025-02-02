package product

import (
	"context"

	"github.com/go-jedi/foodgrammm-backend/internal/domain/product"
	"github.com/go-jedi/foodgrammm-backend/pkg/apperrors"
)

func (s *serv) DeleteExcludeProductsByTelegramID(ctx context.Context, telegramID string, prod string) (product.UserExcludedProducts, error) {
	ie, err := s.userRepository.ExistsByTelegramID(ctx, telegramID)
	if err != nil {
		return product.UserExcludedProducts{}, err
	}

	if !ie {
		return product.UserExcludedProducts{}, apperrors.ErrUserDoesNotExist
	}

	return s.productRepository.DeleteExcludeProductsByTelegramID(ctx, telegramID, prod)
}
