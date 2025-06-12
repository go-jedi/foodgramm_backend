package user

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/go-jedi/foodgramm_backend/internal/domain/user"
	loggermocks "github.com/go-jedi/foodgramm_backend/pkg/logger/mocks"
	"github.com/go-jedi/foodgramm_backend/pkg/postgres"
	poolsmocks "github.com/go-jedi/foodgramm_backend/pkg/postgres/mocks"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAll(t *testing.T) {
	type in struct {
		ctx context.Context
	}

	type want struct {
		users []user.User
		err   error
	}

	var (
		ctx        = context.TODO()
		id         = gofakeit.Int64()
		telegramID = gofakeit.UUID()
		username   = gofakeit.Username()
		firstname  = gofakeit.FirstName()
		lastname   = gofakeit.LastName()
		createdAt  = time.Now()
		updatedAt  = time.Now()
		testUsers  = []user.User{
			{
				ID:         id,
				TelegramID: telegramID,
				Username:   username,
				FirstName:  firstname,
				LastName:   lastname,
				CreatedAt:  createdAt,
				UpdatedAt:  updatedAt,
			},
		}

		queryTimeout = int64(2)
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
				rows := &poolsmocks.RowsMock{
					NextFunc: func() bool {
						return len(testUsers) > 0
					},
					ScanFunc: func(dest ...any) error {
						if len(testUsers) == 0 {
							return pgx.ErrNoRows
						}

						u := testUsers[0]
						testUsers = testUsers[1:]

						*dest[0].(*int64) = u.ID
						*dest[1].(*string) = u.TelegramID
						*dest[2].(*string) = u.Username
						*dest[3].(*string) = u.FirstName
						*dest[4].(*string) = u.LastName
						*dest[5].(*time.Time) = u.CreatedAt
						*dest[6].(*time.Time) = u.UpdatedAt

						return nil
					},
					ErrFunc: func() error {
						return nil
					},
					CloseFunc: func() {},
					CommandTagFunc: func() pgconn.CommandTag {
						return pgconn.NewCommandTag("SELECT")
					},
					FieldDescriptionsFunc: func() []pgconn.FieldDescription {
						return []pgconn.FieldDescription{
							{Name: "id", DataTypeOID: 20},
							{Name: "telegram_id", DataTypeOID: 20},
							{Name: "username", DataTypeOID: 25},
							{Name: "first_name", DataTypeOID: 25},
							{Name: "last_name", DataTypeOID: 25},
							{Name: "created_at", DataTypeOID: 1114},
							{Name: "updated_at", DataTypeOID: 1114},
						}
					},
					ValuesFunc: func() ([]any, error) {
						if len(testUsers) == 0 {
							return nil, pgx.ErrNoRows
						}
						u := testUsers[0]
						return []any{
							u.ID,
							u.TelegramID,
							u.Username,
							u.FirstName,
							u.LastName,
							u.CreatedAt,
							u.UpdatedAt,
						}, nil
					},
					RawValuesFunc: func() [][]byte {
						return nil
					},
				}
				m.On("Query", mock.Anything, mock.Anything).Return(rows, nil)
			},
			mockLoggerBehavior: func(m *loggermocks.ILogger) {
				m.On("Debug", "[get all users] execute repository", mock.Anything)
			},
			in: in{ctx: ctx},
			want: want{
				users: testUsers,
				err:   nil,
			},
		},
		{
			name: "query error",
			mockPoolBehavior: func(m *poolsmocks.IPool) {
				m.On("Query", mock.Anything, mock.Anything).Return(nil, pgx.ErrNoRows)
			},
			mockLoggerBehavior: func(m *loggermocks.ILogger) {
				m.On("Debug", "[get all users] execute repository", mock.Anything)
				m.On("Error", "failed to get users", "err", pgx.ErrNoRows)
			},
			in: in{ctx: ctx},
			want: want{
				users: nil,
				err:   fmt.Errorf("could not get users: %w", pgx.ErrNoRows),
			},
		},
		{
			name: "context timeout",
			mockPoolBehavior: func(m *poolsmocks.IPool) {
				m.On("Query", mock.Anything, mock.Anything).Return(nil, context.DeadlineExceeded)
			},
			mockLoggerBehavior: func(m *loggermocks.ILogger) {
				m.On("Debug", "[get all users] execute repository", mock.Anything)
				m.On("Error", "request timed out while get users", "err", context.DeadlineExceeded)
			},
			in: in{ctx: ctx},
			want: want{
				users: nil,
				err:   fmt.Errorf("the request timed out: %w", context.DeadlineExceeded),
			},
		},
		{
			name: "scan error",
			mockPoolBehavior: func(m *poolsmocks.IPool) {
				rows := &poolsmocks.RowsMock{
					NextFunc: func() bool {
						return true
					},
					ScanFunc: func(dest ...any) error {
						return pgx.ErrNoRows
					},
					ErrFunc: func() error {
						return nil
					},
					CloseFunc: func() {},
				}
				m.On("Query", mock.Anything, mock.Anything).Return(rows, nil)
			},
			mockLoggerBehavior: func(m *loggermocks.ILogger) {
				m.On("Debug", "[get all users] execute repository", mock.Anything)
				m.On("Error", "failed to scan row to get all users", "err", pgx.ErrNoRows)
			},
			in: in{ctx: ctx},
			want: want{
				users: nil,
				err:   fmt.Errorf("failed to scan row to get all users: %w", pgx.ErrNoRows),
			},
		},
		{
			name: "rows error",
			mockPoolBehavior: func(m *poolsmocks.IPool) {
				rows := &poolsmocks.RowsMock{
					NextFunc: func() bool {
						return false
					},
					ScanFunc: func(dest ...any) error {
						return nil
					},
					ErrFunc: func() error {
						return pgx.ErrNoRows
					},
					CloseFunc: func() {},
				}
				m.On("Query", mock.Anything, mock.Anything).Return(rows, nil)
			},
			mockLoggerBehavior: func(m *loggermocks.ILogger) {
				m.On("Debug", "[get all users] execute repository", mock.Anything)
				m.On("Error", "failed to get all users", "err", pgx.ErrNoRows)
			},
			in: in{ctx: ctx},
			want: want{
				users: nil,
				err:   fmt.Errorf("failed to get all users: %w", pgx.ErrNoRows),
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

			result, err := repository.All(test.in.ctx)

			if test.want.err != nil {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), test.want.err.Error())
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, test.want.users, result)

			mockPool.AssertExpectations(t)
			mockLogger.AssertExpectations(t)
		})
	}
}
