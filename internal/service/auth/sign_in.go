package auth

import (
	"context"

	"github.com/go-jedi/foodgrammm-backend/internal/domain/auth"
	"github.com/go-jedi/foodgrammm-backend/internal/domain/user"
)

func (s *serv) SignIn(ctx context.Context, dto auth.SignInDTO) (auth.SignInResp, error) {
	ie, err := s.userRepository.Exists(ctx, dto.TelegramID, dto.Username)
	if err != nil {
		return auth.SignInResp{}, err
	}

	if !ie {
		nu, err := s.SignInCreateUser(ctx, dto)
		if err != nil {
			return auth.SignInResp{}, err
		}

		return s.SignInGenerateTokens(nu.TelegramID)
	}

	return s.SignInGenerateTokens(dto.TelegramID)
}

// SignInCreateUser create new user.
func (s *serv) SignInCreateUser(ctx context.Context, dto auth.SignInDTO) (user.User, error) {
	createDTO := user.CreateDTO{
		TelegramID: dto.TelegramID,
		Username:   dto.Username,
		FirstName:  dto.FirstName,
		LastName:   dto.LastName,
	}

	return s.userRepository.Create(ctx, createDTO)
}

// SignInGenerateTokens generate tokens access and refresh.
func (s *serv) SignInGenerateTokens(telegramID string) (auth.SignInResp, error) {
	tokens, err := s.jwt.Generate(telegramID)
	if err != nil {
		return auth.SignInResp{}, err
	}

	return auth.SignInResp{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
		AccessExpAt:  tokens.AccessExpAt,
		RefreshExpAt: tokens.RefreshExpAt,
	}, nil
}
