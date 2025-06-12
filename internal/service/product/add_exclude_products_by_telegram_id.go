package product

import (
	"context"

	"github.com/go-jedi/foodgramm_backend/internal/domain/product"
	"github.com/go-jedi/foodgramm_backend/pkg/apperrors"
	"github.com/go-jedi/foodgramm_backend/pkg/utils"
)

func (s *serv) AddExcludeProductsByTelegramID(ctx context.Context, dto product.AddExcludeProductsByTelegramIDDTO) (product.UserExcludedProducts, error) {
	s.logger.Debug("[add exclude products by telegram id] execute service")

	ie, err := s.userRepository.ExistsByTelegramID(ctx, dto.TelegramID)
	if err != nil {
		return product.UserExcludedProducts{}, err
	}

	if !ie {
		return product.UserExcludedProducts{}, apperrors.ErrUserDoesNotExist
	}

	dto.Products = utils.RemoveDuplicates(dto.Products)

	return s.productRepository.AddExcludeProductsByTelegramID(ctx, dto)
}
