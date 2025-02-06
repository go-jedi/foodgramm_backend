package service

import (
	"context"

	"github.com/go-jedi/foodgrammm-backend/internal/domain/auth"
	"github.com/go-jedi/foodgrammm-backend/internal/domain/payment"
	"github.com/go-jedi/foodgrammm-backend/internal/domain/product"
	"github.com/go-jedi/foodgrammm-backend/internal/domain/recipe"
	"github.com/go-jedi/foodgrammm-backend/internal/domain/subscription"
	"github.com/go-jedi/foodgrammm-backend/internal/domain/user"
)

type AuthService interface {
	SignIn(ctx context.Context, dto auth.SignInDTO) (auth.SignInResp, error)
	Check(ctx context.Context, dto auth.CheckDTO) (auth.CheckResponse, error)
	Refresh(ctx context.Context, dto auth.RefreshDTO) (auth.RefreshResponse, error)
}

type UserService interface {
	Create(ctx context.Context, dto user.CreateDTO) (user.User, error)
	All(ctx context.Context) ([]user.User, error)
	List(ctx context.Context, dto user.ListDTO) (user.ListResponse, error)
	GetByID(ctx context.Context, userID int64) (user.User, error)
	GetByTelegramID(ctx context.Context, telegramID string) (user.User, error)
	Exists(ctx context.Context, telegramID string, username string) (bool, error)
	ExistsByID(ctx context.Context, userID int64) (bool, error)
	ExistsByTelegramID(ctx context.Context, telegramID string) (bool, error)
	Update(ctx context.Context, dto user.UpdateDTO) (user.User, error)
	DeleteByID(ctx context.Context, id int64) (int64, error)
	DeleteByTelegramID(ctx context.Context, telegramID string) (string, error)
}

type ProductService interface {
	AddExcludeProductsByTelegramID(ctx context.Context, dto product.AddExcludeProductsByTelegramIDDTO) (product.UserExcludedProducts, error)
	GetExcludeProductsByTelegramID(ctx context.Context, telegramID string) (product.UserExcludedProducts, error)
	DeleteExcludeProductsByTelegramID(ctx context.Context, telegramID string, prod string) (product.UserExcludedProducts, error)
}

type SubscriptionService interface {
	GetByTelegramID(ctx context.Context, telegramID string) (subscription.Subscription, error)
	ExistsByTelegramID(ctx context.Context, telegramID string) (bool, error)
}

type RecipeService interface {
	GenerateRecipe(ctx context.Context, dto recipe.GenerateRecipeDTO) (recipe.Recipes, error)
	GetRecipesByTelegramID(ctx context.Context, telegramID string) ([]recipe.Recipes, error)
	GetFreeRecipesByTelegramID(ctx context.Context, telegramID string) (recipe.UserFreeRecipes, error)
}

type PaymentService interface {
	Create(ctx context.Context, dto payment.CreateDTO) error
	CheckStatus(ctx context.Context, dto payment.CheckStatusDTO) error
}
