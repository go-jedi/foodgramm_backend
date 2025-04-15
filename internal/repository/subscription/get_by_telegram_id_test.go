package subscription

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/go-jedi/foodgrammm-backend/internal/domain/subscription"
	loggermocks "github.com/go-jedi/foodgrammm-backend/pkg/logger/mocks"
	"github.com/go-jedi/foodgrammm-backend/pkg/postgres"
	poolsmocks "github.com/go-jedi/foodgrammm-backend/pkg/postgres/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetByTelegramID(t *testing.T) {
	type in struct {
		ctx        context.Context
		telegramID string
	}

	type want struct {
		subscription subscription.Subscription
		err          error
	}

	var (
		ctx              = context.TODO()
		telegramID       = gofakeit.UUID()
		subscribedAt     = time.Now()
		expiresAt        = time.Now().Add(1 * time.Hour)
		testSubscription = subscription.Subscription{
			ID:           gofakeit.Int64(),
			TelegramID:   telegramID,
			SubscribedAt: &subscribedAt,
			ExpiresAt:    &expiresAt,
			IsActive:     gofakeit.Bool(),
			CreatedAt:    gofakeit.Date(),
			UpdatedAt:    gofakeit.Date(),
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
					"SELECT * FROM public.subscription_get_by_telegram_id($1);",
					telegramID,
				).Return(row)

				row.On("Scan",
					mock.AnythingOfType("*int64"),
					mock.AnythingOfType("*string"),
					mock.AnythingOfType("**time.Time"),
					mock.AnythingOfType("**time.Time"),
					mock.AnythingOfType("*bool"),
					mock.AnythingOfType("*time.Time"),
					mock.AnythingOfType("*time.Time"),
				).Run(func(args mock.Arguments) {
					id := args.Get(0).(*int64)
					*id = testSubscription.ID

					telegramID := args.Get(1).(*string)
					*telegramID = testSubscription.TelegramID

					subscribedAt := args.Get(2).(**time.Time)
					*subscribedAt = testSubscription.SubscribedAt

					expiresAt := args.Get(3).(**time.Time)
					*expiresAt = testSubscription.ExpiresAt

					isActive := args.Get(4).(*bool)
					*isActive = testSubscription.IsActive

					createdAt := args.Get(5).(*time.Time)
					*createdAt = testSubscription.CreatedAt

					updatedAt := args.Get(6).(*time.Time)
					*updatedAt = testSubscription.UpdatedAt
				}).Return(nil)
			},
			mockLoggerBehavior: func(m *loggermocks.ILogger) {
				m.On("Debug", "[get subscription by telegram id] execute repository")
			},
			in: in{
				ctx:        ctx,
				telegramID: telegramID,
			},
			want: want{
				subscription: testSubscription,
				err:          nil,
			},
		},
		{
			name: "timeout error",
			mockPoolBehavior: func(m *poolsmocks.IPool, row *poolsmocks.RowMock) {
				m.On("QueryRow",
					mock.Anything,
					"SELECT * FROM public.subscription_get_by_telegram_id($1);",
					telegramID,
				).Return(row)

				row.On("Scan",
					mock.AnythingOfType("*int64"),
					mock.AnythingOfType("*string"),
					mock.AnythingOfType("**time.Time"),
					mock.AnythingOfType("**time.Time"),
					mock.AnythingOfType("*bool"),
					mock.AnythingOfType("*time.Time"),
					mock.AnythingOfType("*time.Time"),
				).Run(func(args mock.Arguments) {
					id := args.Get(0).(*int64)
					*id = testSubscription.ID

					telegramID := args.Get(1).(*string)
					*telegramID = testSubscription.TelegramID

					subscribedAt := args.Get(2).(**time.Time)
					*subscribedAt = testSubscription.SubscribedAt

					expiresAt := args.Get(3).(**time.Time)
					*expiresAt = testSubscription.ExpiresAt

					isActive := args.Get(4).(*bool)
					*isActive = testSubscription.IsActive

					createdAt := args.Get(5).(*time.Time)
					*createdAt = testSubscription.CreatedAt

					updatedAt := args.Get(6).(*time.Time)
					*updatedAt = testSubscription.UpdatedAt
				}).Return(context.DeadlineExceeded)
			},
			mockLoggerBehavior: func(m *loggermocks.ILogger) {
				m.On("Debug", "[get subscription by telegram id] execute repository")
				m.On("Error", "request timed out while get subscription by telegram id", "err", context.DeadlineExceeded)
			},
			in: in{
				ctx:        ctx,
				telegramID: telegramID,
			},
			want: want{
				subscription: subscription.Subscription{},
				err:          errors.New("the request timed out: context deadline exceeded"),
			},
		},
		{
			name: "database error",
			mockPoolBehavior: func(m *poolsmocks.IPool, row *poolsmocks.RowMock) {
				m.On("QueryRow",
					mock.Anything,
					"SELECT * FROM public.subscription_get_by_telegram_id($1);",
					telegramID,
				).Return(row)

				row.On("Scan",
					mock.AnythingOfType("*int64"),
					mock.AnythingOfType("*string"),
					mock.AnythingOfType("**time.Time"),
					mock.AnythingOfType("**time.Time"),
					mock.AnythingOfType("*bool"),
					mock.AnythingOfType("*time.Time"),
					mock.AnythingOfType("*time.Time"),
				).Run(func(args mock.Arguments) {
					id := args.Get(0).(*int64)
					*id = testSubscription.ID

					telegramID := args.Get(1).(*string)
					*telegramID = testSubscription.TelegramID

					subscribedAt := args.Get(2).(**time.Time)
					*subscribedAt = testSubscription.SubscribedAt

					expiresAt := args.Get(3).(**time.Time)
					*expiresAt = testSubscription.ExpiresAt

					isActive := args.Get(4).(*bool)
					*isActive = testSubscription.IsActive

					createdAt := args.Get(5).(*time.Time)
					*createdAt = testSubscription.CreatedAt

					updatedAt := args.Get(6).(*time.Time)
					*updatedAt = testSubscription.UpdatedAt
				}).Return(errors.New("database error"))
			},
			mockLoggerBehavior: func(m *loggermocks.ILogger) {
				m.On("Debug", "[get subscription by telegram id] execute repository")
				m.On("Error", "failed to get subscription by telegram id", "err", errors.New("database error"))
			},
			in: in{
				ctx:        ctx,
				telegramID: telegramID,
			},
			want: want{
				subscription: subscription.Subscription{},
				err:          errors.New("could not get subscription by telegram id: database error"),
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

			assert.Equal(t, test.want.subscription, result)

			mockPool.AssertExpectations(t)
			mockLogger.AssertExpectations(t)
			mockRow.AssertExpectations(t)
		})
	}
}
