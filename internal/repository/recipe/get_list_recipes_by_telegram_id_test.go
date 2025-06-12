package recipe

import (
	"context"
	"encoding/json"
	"errors"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/go-jedi/foodgramm_backend/internal/domain/recipe"
	loggermocks "github.com/go-jedi/foodgramm_backend/pkg/logger/mocks"
	"github.com/go-jedi/foodgramm_backend/pkg/postgres"
	poolsmocks "github.com/go-jedi/foodgramm_backend/pkg/postgres/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetListRecipesByTelegramID(t *testing.T) {
	type in struct {
		ctx context.Context
		dto recipe.GetListRecipesByTelegramIDDTO
	}

	type want struct {
		listRecipes recipe.GetListRecipesByTelegramIDResponse
		err         error
	}

	var (
		ctx  = context.TODO()
		page = gofakeit.Number(1, 1000)
		size = gofakeit.Number(1, 30)
		dto  = recipe.GetListRecipesByTelegramIDDTO{
			TelegramID: gofakeit.UUID(),
			Page:       page,
			Size:       size,
		}
		testListRecipes = recipe.GetListRecipesByTelegramIDResponse{
			TotalCount:  gofakeit.Number(1, 1000),
			TotalPages:  gofakeit.Number(1, 1000),
			CurrentPage: page,
			Size:        size,
			Data:        json.RawMessage{},
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
					dto.TelegramID,
					dto.Size,
					dto.Page,
				).Return(row)

				row.On("Scan",
					mock.AnythingOfType("*int"),
					mock.AnythingOfType("*int"),
					mock.AnythingOfType("*int"),
					mock.AnythingOfType("*int"),
					mock.AnythingOfType("*json.RawMessage"),
				).Run(func(args mock.Arguments) {
					totalCount := args.Get(0).(*int)
					*totalCount = testListRecipes.TotalCount

					totalPages := args.Get(1).(*int)
					*totalPages = testListRecipes.TotalPages

					currentPage := args.Get(2).(*int)
					*currentPage = testListRecipes.CurrentPage

					size := args.Get(3).(*int)
					*size = testListRecipes.Size

					data := args.Get(4).(*json.RawMessage)
					*data = testListRecipes.Data
				}).Return(nil)
			},
			mockLoggerBehavior: func(m *loggermocks.ILogger) {
				m.On("Debug", "[get list recipes by telegram id] execute repository")
			},
			in: in{
				ctx: ctx,
				dto: dto,
			},
			want: want{
				listRecipes: testListRecipes,
				err:         nil,
			},
		},
		{
			name: "timeout error",
			mockPoolBehavior: func(m *poolsmocks.IPool, row *poolsmocks.RowMock) {
				m.On("QueryRow",
					mock.Anything,
					mock.Anything,
					dto.TelegramID,
					dto.Size,
					dto.Page,
				).Return(row)

				row.On("Scan",
					mock.AnythingOfType("*int"),
					mock.AnythingOfType("*int"),
					mock.AnythingOfType("*int"),
					mock.AnythingOfType("*int"),
					mock.AnythingOfType("*json.RawMessage"),
				).Run(func(args mock.Arguments) {
					totalCount := args.Get(0).(*int)
					*totalCount = testListRecipes.TotalCount

					totalPages := args.Get(1).(*int)
					*totalPages = testListRecipes.TotalPages

					currentPage := args.Get(2).(*int)
					*currentPage = testListRecipes.CurrentPage

					size := args.Get(3).(*int)
					*size = testListRecipes.Size

					data := args.Get(4).(*json.RawMessage)
					*data = testListRecipes.Data
				}).Return(context.DeadlineExceeded)
			},
			mockLoggerBehavior: func(m *loggermocks.ILogger) {
				m.On("Debug", "[get list recipes by telegram id] execute repository")
				m.On("Error", "request timed out while get list recipes", "err", context.DeadlineExceeded)
			},
			in: in{
				ctx: ctx,
				dto: dto,
			},
			want: want{
				listRecipes: recipe.GetListRecipesByTelegramIDResponse{},
				err:         errors.New("the request timed out: context deadline exceeded"),
			},
		},
		{
			name: "database error",
			mockPoolBehavior: func(m *poolsmocks.IPool, row *poolsmocks.RowMock) {
				m.On("QueryRow",
					mock.Anything,
					mock.Anything,
					dto.TelegramID,
					dto.Size,
					dto.Page,
				).Return(row)

				row.On("Scan",
					mock.AnythingOfType("*int"),
					mock.AnythingOfType("*int"),
					mock.AnythingOfType("*int"),
					mock.AnythingOfType("*int"),
					mock.AnythingOfType("*json.RawMessage"),
				).Run(func(args mock.Arguments) {
					totalCount := args.Get(0).(*int)
					*totalCount = testListRecipes.TotalCount

					totalPages := args.Get(1).(*int)
					*totalPages = testListRecipes.TotalPages

					currentPage := args.Get(2).(*int)
					*currentPage = testListRecipes.CurrentPage

					size := args.Get(3).(*int)
					*size = testListRecipes.Size

					data := args.Get(4).(*json.RawMessage)
					*data = testListRecipes.Data
				}).Return(errors.New("database error"))
			},
			mockLoggerBehavior: func(m *loggermocks.ILogger) {
				m.On("Debug", "[get list recipes by telegram id] execute repository")
				m.On("Error", "failed to get list recipes", "err", errors.New("database error"))
			},
			in: in{
				ctx: ctx,
				dto: dto,
			},
			want: want{
				listRecipes: recipe.GetListRecipesByTelegramIDResponse{},
				err:         errors.New("could not get list recipes: database error"),
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

			result, err := repository.GetListRecipesByTelegramID(test.in.ctx, test.in.dto)

			if test.want.err != nil {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), test.want.err.Error())
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, test.want.listRecipes, result)

			mockPool.AssertExpectations(t)
			mockLogger.AssertExpectations(t)
			mockRow.AssertExpectations(t)
		})
	}
}
