package user

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/go-jedi/foodgrammm-backend/internal/domain/user"
	loggermocks "github.com/go-jedi/foodgrammm-backend/pkg/logger/mocks"
	"github.com/go-jedi/foodgrammm-backend/pkg/postgres"
	poolsmocks "github.com/go-jedi/foodgrammm-backend/pkg/postgres/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUpdate(t *testing.T) {
	type in struct {
		ctx context.Context
		dto user.UpdateDTO
	}

	type want struct {
		user user.User
		err  error
	}

	var (
		ctx          = context.TODO()
		id           = gofakeit.Int64()
		telegramID   = gofakeit.UUID()
		username     = gofakeit.Username()
		firstName    = gofakeit.FirstName()
		lastName     = gofakeit.LastName()
		createdAt    = gofakeit.Date()
		updatedAt    = gofakeit.Date()
		queryTimeout = int64(2)
		dto          = user.UpdateDTO{
			ID:         id,
			TelegramID: telegramID,
			Username:   username,
			FirstName:  firstName,
			LastName:   lastName,
		}
		testUser = user.User{
			ID:         id,
			TelegramID: telegramID,
			Username:   username,
			FirstName:  firstName,
			LastName:   lastName,
			CreatedAt:  createdAt,
			UpdatedAt:  updatedAt,
		}
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
					dto.Username,
					dto.FirstName,
					dto.LastName,
					dto.ID,
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
				m.On("Debug", "[update user] execute repository")
			},
			in: in{
				ctx: ctx,
				dto: dto,
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
					dto.TelegramID,
					dto.Username,
					dto.FirstName,
					dto.LastName,
					dto.ID,
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
				m.On("Debug", "[update user] execute repository")
				m.On("Error", "request timed out while update user", "err", context.DeadlineExceeded)
			},
			in: in{
				ctx: ctx,
				dto: dto,
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
					dto.TelegramID,
					dto.Username,
					dto.FirstName,
					dto.LastName,
					dto.ID,
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
				m.On("Debug", "[update user] execute repository")
				m.On("Error", "failed to update user", "err", errors.New("database error"))
			},
			in: in{
				ctx: ctx,
				dto: dto,
			},
			want: want{
				user: user.User{},
				err:  errors.New("could not update user: database error"),
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

			assert.Equal(t, test.want.user, result)

			mockPool.AssertExpectations(t)
			mockLogger.AssertExpectations(t)
			mockRow.AssertExpectations(t)
		})
	}
}
