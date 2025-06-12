package user

import (
	"context"
	"errors"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/go-jedi/foodgramm_backend/internal/domain/user"
	loggermocks "github.com/go-jedi/foodgramm_backend/pkg/logger/mocks"
	"github.com/go-jedi/foodgramm_backend/pkg/postgres"
	poolsmocks "github.com/go-jedi/foodgramm_backend/pkg/postgres/mocks"
	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestList(t *testing.T) {
	type in struct {
		ctx context.Context
		dto user.ListDTO
	}

	type want struct {
		list user.ListResponse
		err  error
	}

	var (
		ctx = context.TODO()
		dto = user.ListDTO{
			Page: gofakeit.Number(1, 1000),
			Size: gofakeit.Number(1, 1000),
		}
		list = user.ListResponse{
			TotalCount:  gofakeit.Number(1, 1000),
			TotalPages:  gofakeit.Number(1, 1000),
			CurrentPage: gofakeit.Number(1, 1000),
			Size:        gofakeit.Number(1, 1000),
			Data:        jsoniter.RawMessage{},
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
					dto.Size,
					dto.Page,
				).Return(row)

				row.On("Scan",
					mock.AnythingOfType("*int"),
					mock.AnythingOfType("*int"),
					mock.AnythingOfType("*int"),
					mock.AnythingOfType("*int"),
					mock.AnythingOfType("*jsoniter.RawMessage"),
				).Run(func(args mock.Arguments) {
					totalCount := args.Get(0).(*int)
					*totalCount = list.TotalCount

					totalPages := args.Get(1).(*int)
					*totalPages = list.TotalPages

					currentPage := args.Get(2).(*int)
					*currentPage = list.CurrentPage

					size := args.Get(3).(*int)
					*size = list.Size

					data := args.Get(4).(*jsoniter.RawMessage)
					*data = list.Data
				}).Return(nil)
			},
			mockLoggerBehavior: func(m *loggermocks.ILogger) {
				m.On("Debug", "[get list users with pagination] execute repository")
			},
			in: in{
				ctx: ctx,
				dto: dto,
			},
			want: want{
				list: list,
				err:  nil,
			},
		},
		{
			name: "timeout error",
			mockPoolBehavior: func(m *poolsmocks.IPool, row *poolsmocks.RowMock) {
				m.On("QueryRow",
					mock.Anything,
					mock.Anything,
					dto.Size,
					dto.Page,
				).Return(row)

				row.On("Scan",
					mock.AnythingOfType("*int"),
					mock.AnythingOfType("*int"),
					mock.AnythingOfType("*int"),
					mock.AnythingOfType("*int"),
					mock.AnythingOfType("*jsoniter.RawMessage"),
				).Run(func(args mock.Arguments) {
					totalCount := args.Get(0).(*int)
					*totalCount = list.TotalCount

					totalPages := args.Get(1).(*int)
					*totalPages = list.TotalPages

					currentPage := args.Get(2).(*int)
					*currentPage = list.CurrentPage

					size := args.Get(3).(*int)
					*size = list.Size

					data := args.Get(4).(*jsoniter.RawMessage)
					*data = list.Data
				}).Return(context.DeadlineExceeded)
			},
			mockLoggerBehavior: func(m *loggermocks.ILogger) {
				m.On("Debug", "[get list users with pagination] execute repository")
				m.On("Error", "request timed out while get list users", "err", context.DeadlineExceeded)
			},
			in: in{
				ctx: ctx,
				dto: dto,
			},
			want: want{
				list: user.ListResponse{},
				err:  errors.New("the request timed out: context deadline exceeded"),
			},
		},
		{
			name: "database error",
			mockPoolBehavior: func(m *poolsmocks.IPool, row *poolsmocks.RowMock) {
				m.On("QueryRow",
					mock.Anything,
					mock.Anything,
					dto.Size,
					dto.Page,
				).Return(row)

				row.On("Scan",
					mock.AnythingOfType("*int"),
					mock.AnythingOfType("*int"),
					mock.AnythingOfType("*int"),
					mock.AnythingOfType("*int"),
					mock.AnythingOfType("*jsoniter.RawMessage"),
				).Run(func(args mock.Arguments) {
					totalCount := args.Get(0).(*int)
					*totalCount = list.TotalCount

					totalPages := args.Get(1).(*int)
					*totalPages = list.TotalPages

					currentPage := args.Get(2).(*int)
					*currentPage = list.CurrentPage

					size := args.Get(3).(*int)
					*size = list.Size

					data := args.Get(4).(*jsoniter.RawMessage)
					*data = list.Data
				}).Return(errors.New("database error"))
			},
			mockLoggerBehavior: func(m *loggermocks.ILogger) {
				m.On("Debug", "[get list users with pagination] execute repository")
				m.On("Error", "failed to get list users", "err", errors.New("database error"))
			},
			in: in{
				ctx: ctx,
				dto: dto,
			},
			want: want{
				list: user.ListResponse{},
				err:  errors.New("could not get list users: database error"),
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

			result, err := repository.List(test.in.ctx, test.in.dto)

			if test.want.err != nil {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), test.want.err.Error())
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, test.want.list, result)

			mockPool.AssertExpectations(t)
			mockLogger.AssertExpectations(t)
			mockRow.AssertExpectations(t)
		})
	}
}
