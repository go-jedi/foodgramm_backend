package payment

import (
	"context"

	"github.com/go-jedi/foodgrammm-backend/internal/domain/payment"
	"github.com/go-jedi/foodgrammm-backend/pkg/apperrors"
	"github.com/go-jedi/foodgrammm-backend/pkg/utils"
)

func (s *serv) Create(ctx context.Context, dto payment.CreateDTO) (string, error) {
	s.logger.Debug("[create a payment link] execute service")

	ie, err := s.userRepository.ExistsByTelegramID(ctx, dto.TelegramID)
	if err != nil {
		return "", err
	}

	if !ie {
		return "", apperrors.ErrUserDoesNotExist
	}

	rd, err := s.prepareDataToRequest(dto)
	if err != nil {
		return "", err
	}

	return s.client.Payment.GetLink(ctx, rd)
}

func (s *serv) prepareDataToRequest(data payment.CreateDTO) (payment.GetLinkBody, error) {
	const base = 10

	bi, err := utils.StringToBigInt(data.TelegramID, base)
	if err != nil {
		return payment.GetLinkBody{}, err
	}

	return payment.GetLinkBody{
		TelegramID: bi,
		Type:       data.Type,
		Price:      data.Price,
	}, nil
}
