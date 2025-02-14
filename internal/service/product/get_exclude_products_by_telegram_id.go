package product

import (
	"context"

	"github.com/go-jedi/foodgrammm-backend/internal/domain/product"
	"github.com/go-jedi/foodgrammm-backend/pkg/apperrors"
)

func (s *serv) GetExcludeProductsByTelegramID(ctx context.Context, telegramID string) (product.UserExcludedProducts, error) {
	s.logger.Debug("[GetExcludeProductsByTelegramID] execute service")

	ie, err := s.userRepository.ExistsByTelegramID(ctx, telegramID)
	if err != nil {
		return product.UserExcludedProducts{}, err
	}

	if !ie {
		return product.UserExcludedProducts{}, apperrors.ErrUserDoesNotExist
	}

	return s.productRepository.GetExcludeProductsByTelegramID(ctx, telegramID)
}
