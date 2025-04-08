package repository

import (
	"context"

	"github.com/go-jedi/foodgrammm-backend/internal/domain/parser"
	"github.com/go-jedi/foodgrammm-backend/internal/domain/product"
	"github.com/go-jedi/foodgrammm-backend/internal/domain/promocode"
	"github.com/go-jedi/foodgrammm-backend/internal/domain/recipe"
	recipeevent "github.com/go-jedi/foodgrammm-backend/internal/domain/recipe_event"
	recipeofdays "github.com/go-jedi/foodgrammm-backend/internal/domain/recipe_of_days"
	recipetypes "github.com/go-jedi/foodgrammm-backend/internal/domain/recipe_types"
	"github.com/go-jedi/foodgrammm-backend/internal/domain/subscription"
	"github.com/go-jedi/foodgrammm-backend/internal/domain/user"
)

//go:generate mockery --name=UserRepository --output=mocks --case=underscore
type UserRepository interface {
	Create(ctx context.Context, dto user.CreateDTO) (user.User, error)
	All(ctx context.Context) ([]user.User, error)
	List(ctx context.Context, dto user.ListDTO) (user.ListResponse, error)
	GetUserCount(ctx context.Context) (int64, error)
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

//go:generate mockery --name=ProductRepository --output=mocks --case=underscore
type ProductRepository interface {
	AddExcludeProductsByTelegramID(ctx context.Context, dto product.AddExcludeProductsByTelegramIDDTO) (product.UserExcludedProducts, error)
	GetExcludeProductsByTelegramID(ctx context.Context, telegramID string) (product.UserExcludedProducts, error)
	DeleteExcludeProductsByTelegramID(ctx context.Context, telegramID string, prod string) (product.UserExcludedProducts, error)
}

//go:generate mockery --name=SubscriptionRepository --output=mocks --case=underscore
type SubscriptionRepository interface {
	Create(ctx context.Context, telegramID string) error
	GetByTelegramID(ctx context.Context, telegramID string) (subscription.Subscription, error)
	ExistsByTelegramID(ctx context.Context, telegramID string) (bool, error)
}

//go:generate mockery --name=RecipeRepository --output=mocks --case=underscore
type RecipeRepository interface {
	CreateRecipe(ctx context.Context, accessType subscription.AccessType, data parser.ParsedRecipe) (recipe.Recipes, error)
	GetRecipesByTelegramID(ctx context.Context, telegramID string) ([]recipe.Recipes, error)
	GetListRecipesByTelegramID(ctx context.Context, dto recipe.GetListRecipesByTelegramIDDTO) (recipe.GetListRecipesByTelegramIDResponse, error)
	GetFreeRecipesByTelegramID(ctx context.Context, telegramID string) (recipe.UserFreeRecipes, error)
	ExistsFreeRecipesByTelegramID(ctx context.Context, telegramID string) (bool, error)
}

//go:generate mockery --name=RecipeOfDaysRepository --output=mocks --case=underscore
type RecipeOfDaysRepository interface {
	Create(ctx context.Context, data parser.ParsedRecipeOfDays) error
	Get(ctx context.Context) (recipeofdays.Recipe, error)
}

//go:generate mockery --name=RecipeTypesRepository --output=mocks --case=underscore
type RecipeTypesRepository interface {
	Create(ctx context.Context, dto recipetypes.CreateDTO) (recipetypes.RecipeTypes, error)
	All(ctx context.Context) ([]recipetypes.RecipeTypes, error)
	GetByID(ctx context.Context, recipeTypeID int64) (recipetypes.RecipeTypes, error)
	Exists(ctx context.Context, title string) (bool, error)
	ExistsByRecipeTypeID(ctx context.Context, recipeTypeID int64) (bool, error)
	ExistsExceptCurrent(ctx context.Context, recipeTypeID int64, title string) (bool, error)
	Update(ctx context.Context, dto recipetypes.UpdateDTO) (recipetypes.RecipeTypes, error)
	DeleteByID(ctx context.Context, recipeTypeID int64) (int64, error)
}

//go:generate mockery --name=RecipeEventRepository --output=mocks --case=underscore
type RecipeEventRepository interface {
	Create(ctx context.Context, dto recipeevent.CreateDTO) (recipeevent.Recipe, error)
	All(ctx context.Context) ([]recipeevent.Recipe, error)
	AllByTypeID(ctx context.Context, typeID int64) ([]recipeevent.Recipe, error)
	GetByID(ctx context.Context, recipeID int64) (recipeevent.Recipe, error)
	Update(ctx context.Context, dto recipeevent.UpdateDTO) (recipeevent.Recipe, error)
	DeleteByID(ctx context.Context, recipeID int64) (int64, error)
}

//go:generate mockery --name=PromoCodeRepository --output=mocks --case=underscore
type PromoCodeRepository interface {
	Create(ctx context.Context, dto promocode.CreateDTO) (promocode.PromoCode, error)
	Apply(ctx context.Context, dto promocode.ApplyDTO) (promocode.ApplyResponse, error)
	CheckPromoCodeAvailability(ctx context.Context, code string) (bool, error)
	IsPromoCodeValidForUser(ctx context.Context, code string, telegramID string) (bool, error)
	Exists(ctx context.Context, code string) (bool, error)
}
