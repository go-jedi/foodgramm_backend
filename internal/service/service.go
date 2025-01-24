package service

import (
	"context"

	"github.com/go-jedi/foodgrammm-backend/internal/domain/user"
)

type UserService interface {
	Create(ctx context.Context, dto user.CreateDTO) (user.User, error)
	GetByID(ctx context.Context, userID int64) (user.User, error)
	// List(ctx context.Context) ([]user.User, error)
	// GetByTelegramID(ctx context.Context, telegramID int64) (user.User, error)
	// Exists(ctx context.Context, telegramID int64, username string) (bool, error)
	// Update(ctx context.Context) (user.User, error)
	// Delete(ctx context.Context) (user.User, error)
}

type ProductService interface{}

type RecipeService interface{}

type RecipeIngredientService interface{}
