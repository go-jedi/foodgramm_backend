package service

import (
	"context"

	"github.com/go-jedi/foodgrammm-backend/internal/domain/auth"
	"github.com/go-jedi/foodgrammm-backend/internal/domain/payment"
	"github.com/go-jedi/foodgrammm-backend/internal/domain/product"
	"github.com/go-jedi/foodgrammm-backend/internal/domain/promocode"
	"github.com/go-jedi/foodgrammm-backend/internal/domain/recipe"
	recipeevent "github.com/go-jedi/foodgrammm-backend/internal/domain/recipe_event"
	recipeofdays "github.com/go-jedi/foodgrammm-backend/internal/domain/recipe_of_days"
	recipetypes "github.com/go-jedi/foodgrammm-backend/internal/domain/recipe_types"
	"github.com/go-jedi/foodgrammm-backend/internal/domain/subscription"
	"github.com/go-jedi/foodgrammm-backend/internal/domain/user"
)

//go:generate mockery --name=AuthService --output=mocks --case=underscore
type AuthService interface {
	SignIn(ctx context.Context, dto auth.SignInDTO) (auth.SignInResp, error)
	Check(ctx context.Context, dto auth.CheckDTO) (auth.CheckResponse, error)
	Refresh(ctx context.Context, dto auth.RefreshDTO) (auth.RefreshResponse, error)
}

//go:generate mockery --name=UserService --output=mocks --case=underscore
type UserService interface {
	Create(ctx context.Context, dto user.CreateDTO) (user.User, error)
	All(ctx context.Context) ([]user.User, error)
	List(ctx context.Context, dto user.ListDTO) (user.ListResponse, error)
	GetUserCount(ctx context.Context) (int64, error)
	GetByID(ctx context.Context, userID int64) (user.User, error)
	GetByTelegramID(ctx context.Context, telegramID string) (user.User, error)
	Exists(ctx context.Context, telegramID string, username string) (bool, error)
	ExistsByID(ctx context.Context, userID int64) (bool, error)
	ExistsByTelegramID(ctx context.Context, telegramID string) (bool, error)
	Update(ctx context.Context, dto user.UpdateDTO) (user.User, error)
	DeleteByID(ctx context.Context, id int64) (int64, error)
	DeleteByTelegramID(ctx context.Context, telegramID string) (string, error)
}

//go:generate mockery --name=ProductService --output=mocks --case=underscore
type ProductService interface {
	AddExcludeProductsByTelegramID(ctx context.Context, dto product.AddExcludeProductsByTelegramIDDTO) (product.UserExcludedProducts, error)
	GetExcludeProductsByTelegramID(ctx context.Context, telegramID string) (product.UserExcludedProducts, error)
	DeleteExcludeProductsByTelegramID(ctx context.Context, telegramID string, prod string) (product.UserExcludedProducts, error)
}

//go:generate mockery --name=SubscriptionService --output=mocks --case=underscore
type SubscriptionService interface {
	Create(ctx context.Context, telegramID string) error
	GetByTelegramID(ctx context.Context, telegramID string) (subscription.Subscription, error)
	ExistsByTelegramID(ctx context.Context, telegramID string) (bool, error)
}

//go:generate mockery --name=RecipeService --output=mocks --case=underscore
type RecipeService interface {
	GenerateRecipe(ctx context.Context, dto recipe.GenerateRecipeDTO) (recipe.Recipes, error)
	GetRecipesByTelegramID(ctx context.Context, telegramID string) ([]recipe.Recipes, error)
	GetListRecipesByTelegramID(ctx context.Context, dto recipe.GetListRecipesByTelegramIDDTO) (recipe.GetListRecipesByTelegramIDResponse, error)
	GetFreeRecipesByTelegramID(ctx context.Context, telegramID string) (recipe.UserFreeRecipes, error)
}

//go:generate mockery --name=RecipeOfDaysService --output=mocks --case=underscore
type RecipeOfDaysService interface {
	Create(ctx context.Context) error
	Get(ctx context.Context) (recipeofdays.Recipe, error)
}

//go:generate mockery --name=RecipeTypesService --output=mocks --case=underscore
type RecipeTypesService interface {
	Create(ctx context.Context, dto recipetypes.CreateDTO) (recipetypes.RecipeTypes, error)
	All(ctx context.Context) ([]recipetypes.RecipeTypes, error)
	GetByID(ctx context.Context, recipeTypeID int64) (recipetypes.RecipeTypes, error)
	Update(ctx context.Context, dto recipetypes.UpdateDTO) (recipetypes.RecipeTypes, error)
	DeleteByID(ctx context.Context, recipeTypeID int64) (int64, error)
}

//go:generate mockery --name=RecipeEventService --output=mocks --case=underscore
type RecipeEventService interface {
	Create(ctx context.Context, dto recipeevent.CreateDTO) (recipeevent.Recipe, error)
	All(ctx context.Context) ([]recipeevent.Recipe, error)
	AllByTypeID(ctx context.Context, typeID int64) ([]recipeevent.Recipe, error)
	GetByID(ctx context.Context, recipeID int64) (recipeevent.Recipe, error)
	Update(ctx context.Context, dto recipeevent.UpdateDTO) (recipeevent.Recipe, error)
	DeleteByID(ctx context.Context, recipeID int64) (int64, error)
}

//go:generate mockery --name=PaymentService --output=mocks --case=underscore
type PaymentService interface {
	Create(ctx context.Context, dto payment.CreateDTO) (string, error)
}

//go:generate mockery --name=PromoCodeService --output=mocks --case=underscore
type PromoCodeService interface {
	Create(ctx context.Context, dto promocode.CreateDTO) (promocode.PromoCode, error)
	Apply(ctx context.Context, dto promocode.ApplyDTO) (promocode.ApplyResponse, error)
	IsPromoCodeValidForUser(ctx context.Context, dto promocode.IsPromoCodeValidForUserDTO) (bool, error)
}
