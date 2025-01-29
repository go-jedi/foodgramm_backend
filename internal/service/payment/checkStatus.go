package payment

import (
	"context"

	"github.com/go-jedi/foodgrammm-backend/internal/domain/payment"
)

func (s *serv) CheckStatus(_ context.Context, _ payment.CheckStatusDTO) error {
	return nil
}
