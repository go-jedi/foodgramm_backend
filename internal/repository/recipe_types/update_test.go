package recipetypes

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	recipetypes "github.com/go-jedi/foodgramm_backend/internal/domain/recipe_types"
	loggermocks "github.com/go-jedi/foodgramm_backend/pkg/logger/mocks"
	"github.com/go-jedi/foodgramm_backend/pkg/postgres"
	poolsmocks "github.com/go-jedi/foodgramm_backend/pkg/postgres/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUpdate(t *testing.T) {
	type in struct {
		ctx context.Context
		dto recipetypes.UpdateDTO
	}

	type want struct {
		recipeType recipetypes.RecipeTypes
		err        error
	}

	var (
		ctx   = context.TODO()
		id    = gofakeit.Int64()
		title = gofakeit.BookTitle()
		dto   = recipetypes.UpdateDTO{
			ID:    id,
			Title: title,
		}
		testRecipeType = recipetypes.RecipeTypes{
			ID:        id,
			Title:     title,
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
					dto.Title,
					dto.ID,
				).Return(row)

				row.On("Scan",
					mock.AnythingOfType("*int64"),
					mock.AnythingOfType("*string"),
					mock.AnythingOfType("*time.Time"),
					mock.AnythingOfType("*time.Time"),
				).Run(func(args mock.Arguments) {
					id := args.Get(0).(*int64)
					*id = testRecipeType.ID

					title := args.Get(1).(*string)
					*title = testRecipeType.Title

					createdAt := args.Get(2).(*time.Time)
					*createdAt = testRecipeType.CreatedAt

					updatedAt := args.Get(3).(*time.Time)
					*updatedAt = testRecipeType.UpdatedAt
				}).Return(nil)
			},
			mockLoggerBehavior: func(m *loggermocks.ILogger) {
				m.On("Debug", "[update recipe type] execute repository")
			},
			in: in{
				ctx: ctx,
				dto: dto,
			},
			want: want{
				recipeType: testRecipeType,
				err:        nil,
			},
		},
		{
			name: "timeout error",
			mockPoolBehavior: func(m *poolsmocks.IPool, row *poolsmocks.RowMock) {
				m.On("QueryRow",
					mock.Anything,
					mock.Anything,
					dto.Title,
					dto.ID,
				).Return(row)

				row.On("Scan",
					mock.AnythingOfType("*int64"),
					mock.AnythingOfType("*string"),
					mock.AnythingOfType("*time.Time"),
					mock.AnythingOfType("*time.Time"),
				).Run(func(args mock.Arguments) {
					id := args.Get(0).(*int64)
					*id = testRecipeType.ID

					title := args.Get(1).(*string)
					*title = testRecipeType.Title

					createdAt := args.Get(2).(*time.Time)
					*createdAt = testRecipeType.CreatedAt

					updatedAt := args.Get(3).(*time.Time)
					*updatedAt = testRecipeType.UpdatedAt
				}).Return(context.DeadlineExceeded)
			},
			mockLoggerBehavior: func(m *loggermocks.ILogger) {
				m.On("Debug", "[update recipe type] execute repository")
				m.On("Error", "request timed out while update recipe type", "err", context.DeadlineExceeded)
			},
			in: in{
				ctx: ctx,
				dto: dto,
			},
			want: want{
				recipeType: recipetypes.RecipeTypes{},
				err:        errors.New("the request timed out: context deadline exceeded"),
			},
		},
		{
			name: "database error",
			mockPoolBehavior: func(m *poolsmocks.IPool, row *poolsmocks.RowMock) {
				m.On("QueryRow",
					mock.Anything,
					mock.Anything,
					dto.Title,
					dto.ID,
				).Return(row)

				row.On("Scan",
					mock.AnythingOfType("*int64"),
					mock.AnythingOfType("*string"),
					mock.AnythingOfType("*time.Time"),
					mock.AnythingOfType("*time.Time"),
				).Run(func(args mock.Arguments) {
					id := args.Get(0).(*int64)
					*id = testRecipeType.ID

					title := args.Get(1).(*string)
					*title = testRecipeType.Title

					createdAt := args.Get(2).(*time.Time)
					*createdAt = testRecipeType.CreatedAt

					updatedAt := args.Get(3).(*time.Time)
					*updatedAt = testRecipeType.UpdatedAt
				}).Return(errors.New("database error"))
			},
			mockLoggerBehavior: func(m *loggermocks.ILogger) {
				m.On("Debug", "[update recipe type] execute repository")
				m.On("Error", "failed to update recipe type", "err", errors.New("database error"))
			},
			in: in{
				ctx: ctx,
				dto: dto,
			},
			want: want{
				recipeType: recipetypes.RecipeTypes{},
				err:        errors.New("could not update recipe type: database error"),
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

			result, err := repository.Update(test.in.ctx, test.in.dto)

			if test.want.err != nil {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), test.want.err.Error())
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, test.want.recipeType, result)

			mockPool.AssertExpectations(t)
			mockLogger.AssertExpectations(t)
			mockRow.AssertExpectations(t)
		})
	}
}
