package repository

import (
	"context"

	"github.com/go-jedi/foodgrammm-backend/internal/domain/product"
	"github.com/go-jedi/foodgrammm-backend/internal/domain/user"
)

type UserRepository interface {
	Create(ctx context.Context, dto user.CreateDTO) (user.User, error)
	All(ctx context.Context) ([]user.User, error)
	List(ctx context.Context, dto user.ListDTO) (user.ListResponse, error)
	GetByID(ctx context.Context, userID int64) (user.User, error)
	GetByTelegramID(ctx context.Context, telegramID string) (user.User, error)
	Exists(ctx context.Context, telegramID string, username string) (bool, error)
	ExistsExceptCurrent(ctx context.Context, id int64, telegramID string, username string) (bool, error)
	Update(ctx context.Context, dto user.UpdateDTO) (user.User, error)
	DeleteByID(ctx context.Context, id int64) (int64, error)
	DeleteByTelegramID(ctx context.Context, telegramID string) (string, error)
}

type ProductRepository interface {
	AddExcludeProductsByID(ctx context.Context, dto product.AddExcludeProductsByIDDTO) (product.AddExcludeProductsByIDResponse, error)
	AddExcludeProductsByTelegramID(ctx context.Context, dto product.AddExcludeProductsByTelegramIDDTO) (product.AddExcludeProductsByTelegramIDResponse, error)
}
