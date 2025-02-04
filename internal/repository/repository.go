package repository

import (
	"context"

	"github.com/go-jedi/foodgrammm-backend/internal/domain/product"
	"github.com/go-jedi/foodgrammm-backend/internal/domain/recipe"
	"github.com/go-jedi/foodgrammm-backend/internal/domain/subscription"
	"github.com/go-jedi/foodgrammm-backend/internal/domain/user"
)

type UserRepository interface {
	Create(ctx context.Context, dto user.CreateDTO) (user.User, error)
	All(ctx context.Context) ([]user.User, error)
	List(ctx context.Context, dto user.ListDTO) (user.ListResponse, error)
	GetByID(ctx context.Context, userID int64) (user.User, error)
	GetByTelegramID(ctx context.Context, telegramID string) (user.User, error)
	Exists(ctx context.Context, telegramID string, username string) (bool, error)
	ExistsByID(ctx context.Context, userID int64) (bool, error)
	ExistsByTelegramID(ctx context.Context, telegramID string) (bool, error)
	ExistsExceptCurrent(ctx context.Context, id int64, telegramID string, username string) (bool, error)
	Update(ctx context.Context, dto user.UpdateDTO) (user.User, error)
	DeleteByID(ctx context.Context, id int64) (int64, error)
	DeleteByTelegramID(ctx context.Context, telegramID string) (string, error)
}

type ProductRepository interface {
	AddAllergiesByTelegramID(ctx context.Context, dto product.AddAllergiesByTelegramIDDTO) (product.UserExcludedProducts, error)
	GetAllergiesByTelegramID(ctx context.Context, telegramID string) (product.UserExcludedProducts, error)
	AddExcludeProductsByTelegramID(ctx context.Context, dto product.AddExcludeProductsByTelegramIDDTO) (product.UserExcludedProducts, error)
	GetExcludeProductsByTelegramID(ctx context.Context, telegramID string) (product.UserExcludedProducts, error)
	DeleteExcludeProductsByTelegramID(ctx context.Context, telegramID string, prod string) (product.UserExcludedProducts, error)
}

type SubscriptionRepository interface {
	Create(ctx context.Context, telegramID string) (subscription.Subscription, error)
	GetByTelegramID(ctx context.Context, telegramID string) (subscription.Subscription, error)
	ExistsByTelegramID(ctx context.Context, telegramID string) (bool, error)
}

type RecipeRepository interface {
	GetRecipesByTelegramID(ctx context.Context, telegramID string) ([]recipe.Recipes, error)
	AddFreeRecipesCountByTelegramID(ctx context.Context, telegramID string) (recipe.UserFreeRecipes, error)
	GetFreeRecipesByTelegramID(ctx context.Context, telegramID string) (recipe.UserFreeRecipes, error)
}

type PaymentRepository interface{}
