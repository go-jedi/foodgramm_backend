package recipe

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/go-jedi/foodgramm_backend/internal/domain/recipe"
	loggermocks "github.com/go-jedi/foodgramm_backend/pkg/logger/mocks"
	"github.com/go-jedi/foodgramm_backend/pkg/postgres"
	poolsmocks "github.com/go-jedi/foodgramm_backend/pkg/postgres/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetFreeRecipesByTelegramID(t *testing.T) {
	type in struct {
		ctx        context.Context
		telegramID string
	}

	type want struct {
		freeRecipes recipe.UserFreeRecipes
		err         error
	}

	var (
		ctx             = context.TODO()
		telegramID      = gofakeit.UUID()
		testFreeRecipes = recipe.UserFreeRecipes{
			ID:                 gofakeit.Int64(),
			TelegramID:         telegramID,
			FreeRecipesAllowed: gofakeit.Number(0, 3),
			FreeRecipesUsed:    gofakeit.Number(0, 3),
			CreatedAt:          gofakeit.Date(),
			UpdatedAt:          gofakeit.Date(),
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
					telegramID,
				).Return(row)

				row.On("Scan",
					mock.AnythingOfType("*int64"),
					mock.AnythingOfType("*string"),
					mock.AnythingOfType("*int"),
					mock.AnythingOfType("*int"),
					mock.AnythingOfType("*time.Time"),
					mock.AnythingOfType("*time.Time"),
				).Run(func(args mock.Arguments) {
					id := args.Get(0).(*int64)
					*id = testFreeRecipes.ID

					telegramID := args.Get(1).(*string)
					*telegramID = testFreeRecipes.TelegramID

					freeRecipesAllowed := args.Get(2).(*int)
					*freeRecipesAllowed = testFreeRecipes.FreeRecipesAllowed

					freeRecipesUsed := args.Get(3).(*int)
					*freeRecipesUsed = testFreeRecipes.FreeRecipesUsed

					createdAt := args.Get(4).(*time.Time)
					*createdAt = testFreeRecipes.CreatedAt

					updatedAt := args.Get(5).(*time.Time)
					*updatedAt = testFreeRecipes.UpdatedAt
				}).Return(nil)
			},
			mockLoggerBehavior: func(m *loggermocks.ILogger) {
				m.On("Debug", "[get free recipes by telegram id] execute repository")
			},
			in: in{
				ctx:        ctx,
				telegramID: telegramID,
			},
			want: want{
				freeRecipes: testFreeRecipes,
				err:         nil,
			},
		},
		{
			name: "timeout error",
			mockPoolBehavior: func(m *poolsmocks.IPool, row *poolsmocks.RowMock) {
				m.On("QueryRow",
					mock.Anything,
					mock.Anything,
					telegramID,
				).Return(row)

				row.On("Scan",
					mock.AnythingOfType("*int64"),
					mock.AnythingOfType("*string"),
					mock.AnythingOfType("*int"),
					mock.AnythingOfType("*int"),
					mock.AnythingOfType("*time.Time"),
					mock.AnythingOfType("*time.Time"),
				).Run(func(args mock.Arguments) {
					id := args.Get(0).(*int64)
					*id = testFreeRecipes.ID

					telegramID := args.Get(1).(*string)
					*telegramID = testFreeRecipes.TelegramID

					freeRecipesAllowed := args.Get(2).(*int)
					*freeRecipesAllowed = testFreeRecipes.FreeRecipesAllowed

					freeRecipesUsed := args.Get(3).(*int)
					*freeRecipesUsed = testFreeRecipes.FreeRecipesUsed

					createdAt := args.Get(4).(*time.Time)
					*createdAt = testFreeRecipes.CreatedAt

					updatedAt := args.Get(5).(*time.Time)
					*updatedAt = testFreeRecipes.UpdatedAt
				}).Return(context.DeadlineExceeded)
			},
			mockLoggerBehavior: func(m *loggermocks.ILogger) {
				m.On("Debug", "[get free recipes by telegram id] execute repository")
				m.On("Error", "request timed out while get free recipes by telegram id", "err", context.DeadlineExceeded)
			},
			in: in{
				ctx:        ctx,
				telegramID: telegramID,
			},
			want: want{
				freeRecipes: recipe.UserFreeRecipes{},
				err:         errors.New("the request timed out: context deadline exceeded"),
			},
		},
		{
			name: "database error",
			mockPoolBehavior: func(m *poolsmocks.IPool, row *poolsmocks.RowMock) {
				m.On("QueryRow",
					mock.Anything,
					mock.Anything,
					telegramID,
				).Return(row)

				row.On("Scan",
					mock.AnythingOfType("*int64"),
					mock.AnythingOfType("*string"),
					mock.AnythingOfType("*int"),
					mock.AnythingOfType("*int"),
					mock.AnythingOfType("*time.Time"),
					mock.AnythingOfType("*time.Time"),
				).Run(func(args mock.Arguments) {
					id := args.Get(0).(*int64)
					*id = testFreeRecipes.ID

					telegramID := args.Get(1).(*string)
					*telegramID = testFreeRecipes.TelegramID

					freeRecipesAllowed := args.Get(2).(*int)
					*freeRecipesAllowed = testFreeRecipes.FreeRecipesAllowed

					freeRecipesUsed := args.Get(3).(*int)
					*freeRecipesUsed = testFreeRecipes.FreeRecipesUsed

					createdAt := args.Get(4).(*time.Time)
					*createdAt = testFreeRecipes.CreatedAt

					updatedAt := args.Get(5).(*time.Time)
					*updatedAt = testFreeRecipes.UpdatedAt
				}).Return(errors.New("database error"))
			},
			mockLoggerBehavior: func(m *loggermocks.ILogger) {
				m.On("Debug", "[get free recipes by telegram id] execute repository")
				m.On("Error", "failed to get free recipes by telegram id", "err", errors.New("database error"))
			},
			in: in{
				ctx:        ctx,
				telegramID: telegramID,
			},
			want: want{
				freeRecipes: recipe.UserFreeRecipes{},
				err:         errors.New("could not get free recipes by telegram id: database error"),
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

			result, err := repository.GetFreeRecipesByTelegramID(test.in.ctx, test.in.telegramID)

			if test.want.err != nil {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), test.want.err.Error())
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, test.want.freeRecipes, result)

			mockPool.AssertExpectations(t)
			mockLogger.AssertExpectations(t)
			mockRow.AssertExpectations(t)
		})
	}
}
