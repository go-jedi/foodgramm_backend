package user

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/go-jedi/foodgramm_backend/pkg/apperrors"
	loggermocks "github.com/go-jedi/foodgramm_backend/pkg/logger/mocks"
	"github.com/go-jedi/foodgramm_backend/pkg/postgres"
	poolsmocks "github.com/go-jedi/foodgramm_backend/pkg/postgres/mocks"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDeleteByTelegramID(t *testing.T) {
	type in struct {
		ctx        context.Context
		telegramID string
	}

	type want struct {
		telegramID string
		err        error
	}

	var (
		ctx               = context.TODO()
		telegramID        = gofakeit.UUID()
		queryTimeout      = int64(2)
		successCommandTag = pgconn.NewCommandTag("DELETE 0 1")
		noRowsCommandTag  = pgconn.NewCommandTag("DELETE 0 0")
		emptyCommandTag   = pgconn.NewCommandTag("")
	)

	tests := []struct {
		name               string
		mockPoolBehavior   func(m *poolsmocks.IPool)
		mockLoggerBehavior func(m *loggermocks.ILogger)
		in                 in
		want               want
	}{
		{
			name: "ok",
			mockPoolBehavior: func(m *poolsmocks.IPool) {
				m.On("Exec", mock.Anything, mock.Anything, telegramID).
					Return(successCommandTag, nil)
			},
			mockLoggerBehavior: func(m *loggermocks.ILogger) {
				m.On("Debug", "[delete user by telegram id] execute repository", mock.Anything)
			},
			in: in{
				ctx:        ctx,
				telegramID: telegramID,
			},
			want: want{
				telegramID: telegramID,
				err:        nil,
			},
		},
		{
			name: "timeout error",
			mockPoolBehavior: func(m *poolsmocks.IPool) {
				m.On("Exec", mock.Anything, mock.Anything, telegramID).
					Return(emptyCommandTag, context.DeadlineExceeded)
			},
			mockLoggerBehavior: func(m *loggermocks.ILogger) {
				m.On("Debug", "[delete user by telegram id] execute repository", mock.Anything)
				m.On("Error", "request timed out while delete user by telegramID", "err", mock.Anything)
			},
			in: in{
				ctx:        ctx,
				telegramID: telegramID,
			},
			want: want{
				telegramID: "",
				err:        fmt.Errorf("the request timed out: %w", context.DeadlineExceeded),
			},
		},
		{
			name: "database error",
			mockPoolBehavior: func(m *poolsmocks.IPool) {
				m.On("Exec", mock.Anything, mock.Anything, telegramID).
					Return(emptyCommandTag, errors.New("some db error"))
			},
			mockLoggerBehavior: func(m *loggermocks.ILogger) {
				m.On("Debug", "[delete user by telegram id] execute repository", mock.Anything)
				m.On("Error", "failed to delete user by telegramID", "err", mock.Anything)
			},
			in: in{
				ctx:        ctx,
				telegramID: telegramID,
			},
			want: want{
				telegramID: "",
				err:        errors.New("could not delete user by telegramID: some db error"),
			},
		},
		{
			name: "no rows affected",
			mockPoolBehavior: func(m *poolsmocks.IPool) {
				m.On("Exec", mock.Anything, mock.Anything, telegramID).
					Return(noRowsCommandTag, nil)
			},
			mockLoggerBehavior: func(m *loggermocks.ILogger) {
				m.On("Debug", "[delete user by telegram id] execute repository", mock.Anything)
			},
			in: in{
				ctx:        ctx,
				telegramID: telegramID,
			},
			want: want{
				telegramID: "",
				err:        apperrors.ErrNoRowsWereAffected,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockPool := poolsmocks.NewIPool(t)
			mockLogger := loggermocks.NewILogger(t)

			if test.mockPoolBehavior != nil {
				test.mockPoolBehavior(mockPool)
			}
			if test.mockLoggerBehavior != nil {
				test.mockLoggerBehavior(mockLogger)
			}

			pg := &postgres.Postgres{
				Pool:         mockPool,
				QueryTimeout: queryTimeout,
			}

			repository := NewRepository(mockLogger, pg)

			result, err := repository.DeleteByTelegramID(test.in.ctx, test.in.telegramID)

			if test.want.err != nil {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), test.want.err.Error())
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, test.want.telegramID, result)

			mockPool.AssertExpectations(t)
			mockLogger.AssertExpectations(t)
		})
	}
}
