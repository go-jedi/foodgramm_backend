package recipe

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/go-jedi/foodgrammm-backend/internal/domain/parser"
	"github.com/go-jedi/foodgrammm-backend/internal/domain/recipe"
	"github.com/go-jedi/foodgrammm-backend/internal/domain/subscription"
	loggermocks "github.com/go-jedi/foodgrammm-backend/pkg/logger/mocks"
	"github.com/go-jedi/foodgrammm-backend/pkg/postgres"
	poolsmocks "github.com/go-jedi/foodgrammm-backend/pkg/postgres/mocks"
	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateRecipe(t *testing.T) {
	type in struct {
		ctx        context.Context
		accessType subscription.AccessType
		parsedData parser.ParsedRecipe
	}

	type want struct {
		recipe recipe.Recipes
		err    error
	}

	var (
		ctx        = context.TODO()
		accessType = subscription.SubscriptionAccess
		telegramID = gofakeit.UUID()
		title      = gofakeit.BookTitle()
		parsedData = parser.ParsedRecipe{
			TelegramID: telegramID,
			Title:      title,
			Content: [][]parser.Content{
				{
					{
						ID:                1,
						Type:              "type",
						Title:             "title",
						RecipePreparation: "recipe_preparation",
						Calories:          "calories",
						Bzhu:              "bzhu",
						Ingredients:       []string{"ingredients"},
						MethodPreparation: []string{"method_preparation"},
					},
				},
			},
		}
		testRecipe = recipe.Recipes{
			ID:         gofakeit.Int64(),
			TelegramID: telegramID,
			Title:      title,
			Content: [][]recipe.Content{
				{
					{
						ID:                1,
						Type:              "type",
						Title:             "title",
						RecipePreparation: "recipe_preparation",
						Calories:          "calories",
						Bzhu:              "bzhu",
						Ingredients:       []string{"ingredients"},
						MethodPreparation: []string{"method_preparation"},
					},
				},
			},
			CreatedAt: gofakeit.Date(),
			UpdatedAt: gofakeit.Date(),
		}
		queryTimeout = int64(2)
	)

	tests := []struct {
		name               string
		mockPoolBehavior   func(m *poolsmocks.IPool, row *poolsmocks.RowMock)
		mockLoggerBehavior func(m *loggermocks.ILogger)
		in                 in
		want               want
	}{
		{
			name: "ok",
			mockPoolBehavior: func(m *poolsmocks.IPool, row *poolsmocks.RowMock) {
				rawData, _ := jsoniter.Marshal(parsedData)

				m.On("QueryRow",
					mock.Anything,
					`SELECT * FROM public.recipe_create($1, $2);`,
					accessType,
					rawData,
				).Return(row)

				row.On("Scan",
					mock.AnythingOfType("*int64"),
					mock.AnythingOfType("*string"),
					mock.AnythingOfType("*string"),
					mock.AnythingOfType("*[][]recipe.Content"),
					mock.AnythingOfType("*time.Time"),
					mock.AnythingOfType("*time.Time"),
				).Run(func(args mock.Arguments) {
					id := args.Get(0).(*int64)
					*id = testRecipe.ID

					telegramID := args.Get(1).(*string)
					*telegramID = testRecipe.TelegramID

					title := args.Get(2).(*string)
					*title = testRecipe.Title

					content := args.Get(3).(*[][]recipe.Content)
					*content = testRecipe.Content

					createdAt := args.Get(4).(*time.Time)
					*createdAt = testRecipe.CreatedAt

					updatedAt := args.Get(5).(*time.Time)
					*updatedAt = testRecipe.UpdatedAt
				}).Return(nil)
			},
			mockLoggerBehavior: func(m *loggermocks.ILogger) {
				m.On("Debug", "[create recipe] execute repository")
			},
			in: in{
				ctx:        ctx,
				accessType: accessType,
				parsedData: parsedData,
			},
			want: want{
				recipe: testRecipe,
				err:    nil,
			},
		},
		{
			name: "timeout error",
			mockPoolBehavior: func(m *poolsmocks.IPool, row *poolsmocks.RowMock) {
				rawData, _ := jsoniter.Marshal(parsedData)

				m.On("QueryRow",
					mock.Anything,
					`SELECT * FROM public.recipe_create($1, $2);`,
					accessType,
					rawData,
				).Return(row)

				row.On("Scan",
					mock.AnythingOfType("*int64"),
					mock.AnythingOfType("*string"),
					mock.AnythingOfType("*string"),
					mock.AnythingOfType("*[][]recipe.Content"),
					mock.AnythingOfType("*time.Time"),
					mock.AnythingOfType("*time.Time"),
				).Run(func(args mock.Arguments) {
					id := args.Get(0).(*int64)
					*id = testRecipe.ID

					telegramID := args.Get(1).(*string)
					*telegramID = testRecipe.TelegramID

					title := args.Get(2).(*string)
					*title = testRecipe.Title

					content := args.Get(3).(*[][]recipe.Content)
					*content = testRecipe.Content

					createdAt := args.Get(4).(*time.Time)
					*createdAt = testRecipe.CreatedAt

					updatedAt := args.Get(5).(*time.Time)
					*updatedAt = testRecipe.UpdatedAt
				}).Return(context.DeadlineExceeded)
			},
			mockLoggerBehavior: func(m *loggermocks.ILogger) {
				m.On("Debug", "[create recipe] execute repository")
				m.On("Error", "request timed out while creating the recipe", "err", context.DeadlineExceeded)
			},
			in: in{
				ctx:        ctx,
				accessType: accessType,
				parsedData: parsedData,
			},
			want: want{
				recipe: recipe.Recipes{},
				err:    errors.New("the request timed out"),
			},
		},
		{
			name: "database error",
			mockPoolBehavior: func(m *poolsmocks.IPool, row *poolsmocks.RowMock) {
				rawData, _ := jsoniter.Marshal(parsedData)

				m.On("QueryRow",
					mock.Anything,
					`SELECT * FROM public.recipe_create($1, $2);`,
					accessType,
					rawData,
				).Return(row)

				row.On("Scan",
					mock.AnythingOfType("*int64"),
					mock.AnythingOfType("*string"),
					mock.AnythingOfType("*string"),
					mock.AnythingOfType("*[][]recipe.Content"),
					mock.AnythingOfType("*time.Time"),
					mock.AnythingOfType("*time.Time"),
				).Run(func(args mock.Arguments) {
					id := args.Get(0).(*int64)
					*id = testRecipe.ID

					telegramID := args.Get(1).(*string)
					*telegramID = testRecipe.TelegramID

					title := args.Get(2).(*string)
					*title = testRecipe.Title

					content := args.Get(3).(*[][]recipe.Content)
					*content = testRecipe.Content

					createdAt := args.Get(4).(*time.Time)
					*createdAt = testRecipe.CreatedAt

					updatedAt := args.Get(5).(*time.Time)
					*updatedAt = testRecipe.UpdatedAt
				}).Return(errors.New("database error"))
			},
			mockLoggerBehavior: func(m *loggermocks.ILogger) {
				m.On("Debug", "[create recipe] execute repository")
				m.On("Error", "failed to create recipe", "err", errors.New("database error"))
			},
			in: in{
				ctx:        ctx,
				accessType: accessType,
				parsedData: parsedData,
			},
			want: want{
				recipe: recipe.Recipes{},
				err:    errors.New("could not create recipe: database error"),
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockPool := poolsmocks.NewIPool(t)
			mockRow := poolsmocks.NewMockRow(t)
			mockLogger := loggermocks.NewILogger(t)

			if test.mockPoolBehavior != nil {
				test.mockPoolBehavior(mockPool, mockRow)
			}
			if test.mockLoggerBehavior != nil {
				test.mockLoggerBehavior(mockLogger)
			}

			pg := &postgres.Postgres{
				Pool:         mockPool,
				QueryTimeout: queryTimeout,
			}

			repository := NewRepository(mockLogger, pg)

			result, err := repository.CreateRecipe(test.in.ctx, test.in.accessType, test.in.parsedData)

			if test.want.err != nil {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), test.want.err.Error())
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, test.want.recipe, result)

			mockPool.AssertExpectations(t)
			mockLogger.AssertExpectations(t)
			mockRow.AssertExpectations(t)
		})
	}
}
