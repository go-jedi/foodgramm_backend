package user

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/go-jedi/foodgramm_backend/internal/domain/user"
	loggermocks "github.com/go-jedi/foodgramm_backend/pkg/logger/mocks"
	"github.com/go-jedi/foodgramm_backend/pkg/postgres"
	poolsmocks "github.com/go-jedi/foodgramm_backend/pkg/postgres/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetByTelegramID(t *testing.T) {
	type in struct {
		ctx        context.Context
		telegramID string
	}

	type want struct {
		user user.User
		err  error
	}

	var (
		ctx        = context.TODO()
		telegramID = gofakeit.UUID()
		testUser   = user.User{
			ID:         gofakeit.Int64(),
			TelegramID: telegramID,
			Username:   gofakeit.Username(),
			FirstName:  gofakeit.FirstName(),
			LastName:   gofakeit.LastName(),
			CreatedAt:  gofakeit.Date(),
			UpdatedAt:  gofakeit.Date(),
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
					mock.AnythingOfType("*string"),
					mock.AnythingOfType("*string"),
					mock.AnythingOfType("*string"),
					mock.AnythingOfType("*time.Time"),
					mock.AnythingOfType("*time.Time"),
				).Run(func(args mock.Arguments) {
					id := args.Get(0).(*int64)
					*id = testUser.ID

					telegramID := args.Get(1).(*string)
					*telegramID = testUser.TelegramID

					username := args.Get(2).(*string)
					*username = testUser.Username

					firstName := args.Get(3).(*string)
					*firstName = testUser.FirstName

					lastName := args.Get(4).(*string)
					*lastName = testUser.LastName

					createdAt := args.Get(5).(*time.Time)
					*createdAt = testUser.CreatedAt

					updatedAt := args.Get(6).(*time.Time)
					*updatedAt = testUser.UpdatedAt
				}).Return(nil)
			},
			mockLoggerBehavior: func(m *loggermocks.ILogger) {
				m.On("Debug", "[get user by telegram id] execute repository")
			},
			in: in{
				ctx:        ctx,
				telegramID: telegramID,
			},
			want: want{
				user: testUser,
				err:  nil,
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
					mock.AnythingOfType("*string"),
					mock.AnythingOfType("*string"),
					mock.AnythingOfType("*string"),
					mock.AnythingOfType("*time.Time"),
					mock.AnythingOfType("*time.Time"),
				).Run(func(args mock.Arguments) {
					id := args.Get(0).(*int64)
					*id = testUser.ID

					telegramID := args.Get(1).(*string)
					*telegramID = testUser.TelegramID

					username := args.Get(2).(*string)
					*username = testUser.Username

					firstName := args.Get(3).(*string)
					*firstName = testUser.FirstName

					lastName := args.Get(4).(*string)
					*lastName = testUser.LastName

					createdAt := args.Get(5).(*time.Time)
					*createdAt = testUser.CreatedAt

					updatedAt := args.Get(6).(*time.Time)
					*updatedAt = testUser.UpdatedAt
				}).Return(context.DeadlineExceeded)
			},
			mockLoggerBehavior: func(m *loggermocks.ILogger) {
				m.On("Debug", "[get user by telegram id] execute repository")
				m.On("Error", "request timed out while get user by telegram id", "err", context.DeadlineExceeded)
			},
			in: in{
				ctx:        ctx,
				telegramID: telegramID,
			},
			want: want{
				user: user.User{},
				err:  errors.New("the request timed out: context deadline exceeded"),
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
					mock.AnythingOfType("*string"),
					mock.AnythingOfType("*string"),
					mock.AnythingOfType("*string"),
					mock.AnythingOfType("*time.Time"),
					mock.AnythingOfType("*time.Time"),
				).Run(func(args mock.Arguments) {
					id := args.Get(0).(*int64)
					*id = testUser.ID

					telegramID := args.Get(1).(*string)
					*telegramID = testUser.TelegramID

					username := args.Get(2).(*string)
					*username = testUser.Username

					firstName := args.Get(3).(*string)
					*firstName = testUser.FirstName

					lastName := args.Get(4).(*string)
					*lastName = testUser.LastName

					createdAt := args.Get(5).(*time.Time)
					*createdAt = testUser.CreatedAt

					updatedAt := args.Get(6).(*time.Time)
					*updatedAt = testUser.UpdatedAt
				}).Return(errors.New("database error"))
			},
			mockLoggerBehavior: func(m *loggermocks.ILogger) {
				m.On("Debug", "[get user by telegram id] execute repository")
				m.On("Error", "failed to get user by telegram id", "err", errors.New("database error"))
			},
			in: in{
				ctx:        ctx,
				telegramID: telegramID,
			},
			want: want{
				user: user.User{},
				err:  errors.New("could not get user by telegram id: database error"),
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

			result, err := repository.GetByTelegramID(test.in.ctx, test.in.telegramID)

			if test.want.err != nil {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), test.want.err.Error())
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, test.want.user, result)

			mockPool.AssertExpectations(t)
			mockLogger.AssertExpectations(t)
			mockRow.AssertExpectations(t)
		})
	}
}
