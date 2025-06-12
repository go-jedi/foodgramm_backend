package user

import (
	"context"
	"errors"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	loggermocks "github.com/go-jedi/foodgramm_backend/pkg/logger/mocks"
	"github.com/go-jedi/foodgramm_backend/pkg/postgres"
	poolsmocks "github.com/go-jedi/foodgramm_backend/pkg/postgres/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetUserCount(t *testing.T) {
	type in struct {
		ctx context.Context
	}

	type want struct {
		count int64
		err   error
	}

	var (
		ctx          = context.TODO()
		count        = gofakeit.Int64()
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
				).Run(func(args mock.Arguments) {
					cnt := args.Get(0).(*int64)
					*cnt = count
				}).Return(nil)
			},
			mockLoggerBehavior: func(m *loggermocks.ILogger) {
				m.On("Debug", "[get user count] execute repository")
			},
			in: in{
				ctx: ctx,
			},
			want: want{
				count: count,
				err:   nil,
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
				).Run(func(args mock.Arguments) {
					cnt := args.Get(0).(*int64)
					*cnt = count
				}).Return(context.DeadlineExceeded)
			},
			mockLoggerBehavior: func(m *loggermocks.ILogger) {
				m.On("Debug", "[get user count] execute repository")
				m.On("Error", "request timed out while get user count", "err", context.DeadlineExceeded)
			},
			in: in{
				ctx: ctx,
			},
			want: want{
				count: 0,
				err:   errors.New("the request timed out"),
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
				).Run(func(args mock.Arguments) {
					cnt := args.Get(0).(*int64)
					*cnt = count
				}).Return(errors.New("database error"))
			},
			mockLoggerBehavior: func(m *loggermocks.ILogger) {
				m.On("Debug", "[get user count] execute repository")
				m.On("Error", "failed to get user count", "err", errors.New("database error"))
			},
			in: in{
				ctx: ctx,
			},
			want: want{
				count: 0,
				err:   errors.New("could not get user count: database error"),
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

			result, err := repository.GetUserCount(test.in.ctx)

			if test.want.err != nil {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), test.want.err.Error())
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, test.want.count, result)

			mockPool.AssertExpectations(t)
			mockLogger.AssertExpectations(t)
			mockRow.AssertExpectations(t)
		})
	}
}
