package product

import (
	"context"

	"github.com/go-jedi/foodgrammm-backend/internal/domain/product"
	"github.com/go-jedi/foodgrammm-backend/pkg/apperrors"
)

func (s *serv) AddAllergiesByTelegramID(ctx context.Context, dto product.AddAllergiesByTelegramIDDTO) (product.UserExcludedProducts, error) {
	ie, err := s.userRepository.ExistsByTelegramID(ctx, dto.TelegramID)
	if err != nil {
		return product.UserExcludedProducts{}, err
	}

	if !ie {
		return product.UserExcludedProducts{}, apperrors.ErrUserDoesNotExist
	}

	return s.productRepository.AddAllergiesByTelegramID(ctx, dto)
}
