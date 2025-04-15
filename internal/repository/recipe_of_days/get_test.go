package recipeofdays

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	recipeofdays "github.com/go-jedi/foodgrammm-backend/internal/domain/recipe_of_days"
	loggermocks "github.com/go-jedi/foodgrammm-backend/pkg/logger/mocks"
	"github.com/go-jedi/foodgrammm-backend/pkg/postgres"
	poolsmocks "github.com/go-jedi/foodgrammm-backend/pkg/postgres/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGet(t *testing.T) {
	type in struct {
		ctx context.Context
	}

	type want struct {
		recipeOfDays recipeofdays.Recipe
		err          error
	}

	var (
		ctx              = context.TODO()
		testRecipeOfDays = recipeofdays.Recipe{
			ID:    gofakeit.Int64(),
			Title: gofakeit.BookTitle(),
			Lifehack: recipeofdays.Lifehack{
				Name:        gofakeit.Name(),
				Description: gofakeit.ProductDescription(),
			},
			Content: [][]recipeofdays.Content{
				{
					{
						ID:                gofakeit.Int64(),
						Type:              gofakeit.AnimalType(),
						Title:             gofakeit.BookTitle(),
						RecipePreparation: gofakeit.ProductDescription(),
						Calories:          gofakeit.ProductDescription(),
						Bzhu:              gofakeit.ProductDescription(),
						Ingredients:       gofakeit.NiceColors(),
						MethodPreparation: gofakeit.NiceColors(),
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
				m.On("QueryRow",
					mock.Anything,
					mock.Anything,
				).Return(row)

				row.On("Scan",
					mock.AnythingOfType("*int64"),
					mock.AnythingOfType("*string"),
					mock.AnythingOfType("*recipeofdays.Lifehack"),
					mock.AnythingOfType("*[][]recipeofdays.Content"),
					mock.AnythingOfType("*time.Time"),
					mock.AnythingOfType("*time.Time"),
				).Run(func(args mock.Arguments) {
					id := args.Get(0).(*int64)
					*id = testRecipeOfDays.ID

					title := args.Get(1).(*string)
					*title = testRecipeOfDays.Title

					lifeHack := args.Get(2).(*recipeofdays.Lifehack)
					*lifeHack = testRecipeOfDays.Lifehack

					content := args.Get(3).(*[][]recipeofdays.Content)
					*content = testRecipeOfDays.Content

					createdAt := args.Get(4).(*time.Time)
					*createdAt = testRecipeOfDays.CreatedAt

					updatedAt := args.Get(5).(*time.Time)
					*updatedAt = testRecipeOfDays.UpdatedAt
				}).Return(nil)
			},
			mockLoggerBehavior: func(m *loggermocks.ILogger) {
				m.On("Debug", "[get recipe of the day] execute repository")
			},
			in: in{
				ctx: ctx,
			},
			want: want{
				recipeOfDays: testRecipeOfDays,
				err:          nil,
			},
		},
		{
			name: "timeout error",
			mockPoolBehavior: func(m *poolsmocks.IPool, row *poolsmocks.RowMock) {
				m.On("QueryRow",
					mock.Anything,
					mock.Anything,
				).Return(row)

				row.On("Scan",
					mock.AnythingOfType("*int64"),
					mock.AnythingOfType("*string"),
					mock.AnythingOfType("*recipeofdays.Lifehack"),
					mock.AnythingOfType("*[][]recipeofdays.Content"),
					mock.AnythingOfType("*time.Time"),
					mock.AnythingOfType("*time.Time"),
				).Run(func(args mock.Arguments) {
					id := args.Get(0).(*int64)
					*id = testRecipeOfDays.ID

					title := args.Get(1).(*string)
					*title = testRecipeOfDays.Title

					lifeHack := args.Get(2).(*recipeofdays.Lifehack)
					*lifeHack = testRecipeOfDays.Lifehack

					content := args.Get(3).(*[][]recipeofdays.Content)
					*content = testRecipeOfDays.Content

					createdAt := args.Get(4).(*time.Time)
					*createdAt = testRecipeOfDays.CreatedAt

					updatedAt := args.Get(5).(*time.Time)
					*updatedAt = testRecipeOfDays.UpdatedAt
				}).Return(context.DeadlineExceeded)
			},
			mockLoggerBehavior: func(m *loggermocks.ILogger) {
				m.On("Debug", "[get recipe of the day] execute repository")
				m.On("Error", "request timed out while get recipe of days", "err", context.DeadlineExceeded)
			},
			in: in{
				ctx: ctx,
			},
			want: want{
				recipeOfDays: recipeofdays.Recipe{},
				err:          errors.New("the request timed out: context deadline exceeded"),
			},
		},
		{
			name: "database error",
			mockPoolBehavior: func(m *poolsmocks.IPool, row *poolsmocks.RowMock) {
				m.On("QueryRow",
					mock.Anything,
					mock.Anything,
				).Return(row)

				row.On("Scan",
					mock.AnythingOfType("*int64"),
					mock.AnythingOfType("*string"),
					mock.AnythingOfType("*recipeofdays.Lifehack"),
					mock.AnythingOfType("*[][]recipeofdays.Content"),
					mock.AnythingOfType("*time.Time"),
					mock.AnythingOfType("*time.Time"),
				).Run(func(args mock.Arguments) {
					id := args.Get(0).(*int64)
					*id = testRecipeOfDays.ID

					title := args.Get(1).(*string)
					*title = testRecipeOfDays.Title

					lifeHack := args.Get(2).(*recipeofdays.Lifehack)
					*lifeHack = testRecipeOfDays.Lifehack

					content := args.Get(3).(*[][]recipeofdays.Content)
					*content = testRecipeOfDays.Content

					createdAt := args.Get(4).(*time.Time)
					*createdAt = testRecipeOfDays.CreatedAt

					updatedAt := args.Get(5).(*time.Time)
					*updatedAt = testRecipeOfDays.UpdatedAt
				}).Return(errors.New("database error"))
			},
			mockLoggerBehavior: func(m *loggermocks.ILogger) {
				m.On("Debug", "[get recipe of the day] execute repository")
				m.On("Error", "failed to get recipe of days", "err", errors.New("database error"))
			},
			in: in{
				ctx: ctx,
			},
			want: want{
				recipeOfDays: recipeofdays.Recipe{},
				err:          errors.New("could not get recipe of days: database error"),
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

			result, err := repository.Get(test.in.ctx)

			if test.want.err != nil {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), test.want.err.Error())
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, test.want.recipeOfDays, result)

			mockPool.AssertExpectations(t)
			mockLogger.AssertExpectations(t)
			mockRow.AssertExpectations(t)
		})
	}
}
