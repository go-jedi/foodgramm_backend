package recipeevent

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/go-jedi/foodgrammm-backend/pkg/apperrors"
	loggermocks "github.com/go-jedi/foodgrammm-backend/pkg/logger/mocks"
	"github.com/go-jedi/foodgrammm-backend/pkg/postgres"
	poolsmocks "github.com/go-jedi/foodgrammm-backend/pkg/postgres/mocks"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDeleteByID(t *testing.T) {
	type in struct {
		ctx      context.Context
		recipeID int64
	}

	type want struct {
		recipeID int64
		err      error
	}

	var (
		ctx               = context.TODO()
		recipeID          = gofakeit.Int64()
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
				m.On("Exec", mock.Anything, mock.Anything, recipeID).
					Return(successCommandTag, nil)
			},
			mockLoggerBehavior: func(m *loggermocks.ILogger) {
				m.On("Debug", "[delete recipe event by id] execute repository", mock.Anything)
			},
			in: in{
				ctx:      ctx,
				recipeID: recipeID,
			},
			want: want{
				recipeID: recipeID,
				err:      nil,
			},
		},
		{
			name: "timeout error",
			mockPoolBehavior: func(m *poolsmocks.IPool) {
				m.On("Exec", mock.Anything, mock.Anything, recipeID).
					Return(emptyCommandTag, context.DeadlineExceeded)
			},
			mockLoggerBehavior: func(m *loggermocks.ILogger) {
				m.On("Debug", "[delete recipe event by id] execute repository", mock.Anything)
				m.On("Error", "request timed out while delete recipe event by id", "err", mock.Anything)
			},
			in: in{
				ctx:      ctx,
				recipeID: recipeID,
			},
			want: want{
				recipeID: 0,
				err:      fmt.Errorf("the request timed out: %w", context.DeadlineExceeded),
			},
		},
		{
			name: "database error",
			mockPoolBehavior: func(m *poolsmocks.IPool) {
				m.On("Exec", mock.Anything, mock.Anything, recipeID).
					Return(emptyCommandTag, errors.New("some db error"))
			},
			mockLoggerBehavior: func(m *loggermocks.ILogger) {
				m.On("Debug", "[delete recipe event by id] execute repository", mock.Anything)
				m.On("Error", "failed to delete recipe event by id", "err", mock.Anything)
			},
			in: in{
				ctx:      ctx,
				recipeID: recipeID,
			},
			want: want{
				recipeID: 0,
				err:      errors.New("could not delete recipe event by id: some db error"),
			},
		},
		{
			name: "no rows affected",
			mockPoolBehavior: func(m *poolsmocks.IPool) {
				m.On("Exec", mock.Anything, mock.Anything, recipeID).
					Return(noRowsCommandTag, nil)
			},
			mockLoggerBehavior: func(m *loggermocks.ILogger) {
				m.On("Debug", "[delete recipe event by id] execute repository", mock.Anything)
			},
			in: in{
				ctx:      ctx,
				recipeID: recipeID,
			},
			want: want{
				recipeID: 0,
				err:      apperrors.ErrNoRowsWereAffected,
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

			result, err := repository.DeleteByID(test.in.ctx, test.in.recipeID)

			if test.want.err != nil {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), test.want.err.Error())
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, test.want.recipeID, result)

			mockPool.AssertExpectations(t)
			mockLogger.AssertExpectations(t)
		})
	}
}
