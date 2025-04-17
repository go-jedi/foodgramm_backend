package recipeevent

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	recipeevent "github.com/go-jedi/foodgrammm-backend/internal/domain/recipe_event"
	loggermocks "github.com/go-jedi/foodgrammm-backend/pkg/logger/mocks"
	"github.com/go-jedi/foodgrammm-backend/pkg/postgres"
	poolsmocks "github.com/go-jedi/foodgrammm-backend/pkg/postgres/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUpdate(t *testing.T) {
	type in struct {
		ctx context.Context
		dto recipeevent.UpdateDTO
	}

	type want struct {
		recipeEvent recipeevent.Recipe
		err         error
	}

	var (
		ctx          = context.TODO()
		queryTimeout = int64(2)
		id           = gofakeit.Int64()
		typeID       = gofakeit.Int64()
		Title        = gofakeit.BookTitle()
		dto          = recipeevent.UpdateDTO{
			ID:     id,
			Title:  Title,
			TypeID: typeID,
		}
		testRecipeEvent = recipeevent.Recipe{
			ID:     id,
			TypeID: typeID,
			Title:  Title,
			Content: [][]recipeevent.Content{
				{
					{
						ID:                gofakeit.Int64(),
						Type:              gofakeit.UUID(),
						Title:             gofakeit.BookTitle(),
						RecipePreparation: gofakeit.BookTitle(),
						Calories:          gofakeit.BookTitle(),
						Bzhu:              gofakeit.BookTitle(),
						Ingredients:       gofakeit.NiceColors(),
						MethodPreparation: gofakeit.NiceColors(),
					},
				},
			},
			CreatedAt: gofakeit.Date(),
			UpdatedAt: gofakeit.Date(),
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
					dto.TypeID,
					dto.Title,
					dto.ID,
				).Return(row)

				row.On("Scan",
					mock.AnythingOfType("*int64"),
					mock.AnythingOfType("*int64"),
					mock.AnythingOfType("*string"),
					mock.AnythingOfType("*[][]recipeevent.Content"),
					mock.AnythingOfType("*time.Time"),
					mock.AnythingOfType("*time.Time"),
				).Run(func(args mock.Arguments) {
					id := args.Get(0).(*int64)
					*id = testRecipeEvent.ID

					typeID := args.Get(1).(*int64)
					*typeID = testRecipeEvent.TypeID

					title := args.Get(2).(*string)
					*title = testRecipeEvent.Title

					content := args.Get(3).(*[][]recipeevent.Content)
					*content = testRecipeEvent.Content

					createdAt := args.Get(4).(*time.Time)
					*createdAt = testRecipeEvent.CreatedAt

					updatedAt := args.Get(5).(*time.Time)
					*updatedAt = testRecipeEvent.UpdatedAt
				}).Return(nil)
			},
			mockLoggerBehavior: func(m *loggermocks.ILogger) {
				m.On("Debug", "[update recipe event] execute repository")
			},
			in: in{
				ctx: ctx,
				dto: dto,
			},
			want: want{
				recipeEvent: testRecipeEvent,
				err:         nil,
			},
		},
		{
			name: "timeout error",
			mockPoolBehavior: func(m *poolsmocks.IPool, row *poolsmocks.RowMock) {
				m.On("QueryRow",
					mock.Anything,
					mock.Anything,
					dto.TypeID,
					dto.Title,
					dto.ID,
				).Return(row)

				row.On("Scan",
					mock.AnythingOfType("*int64"),
					mock.AnythingOfType("*int64"),
					mock.AnythingOfType("*string"),
					mock.AnythingOfType("*[][]recipeevent.Content"),
					mock.AnythingOfType("*time.Time"),
					mock.AnythingOfType("*time.Time"),
				).Run(func(args mock.Arguments) {
					id := args.Get(0).(*int64)
					*id = testRecipeEvent.ID

					typeID := args.Get(1).(*int64)
					*typeID = testRecipeEvent.TypeID

					title := args.Get(2).(*string)
					*title = testRecipeEvent.Title

					content := args.Get(3).(*[][]recipeevent.Content)
					*content = testRecipeEvent.Content

					createdAt := args.Get(4).(*time.Time)
					*createdAt = testRecipeEvent.CreatedAt

					updatedAt := args.Get(5).(*time.Time)
					*updatedAt = testRecipeEvent.UpdatedAt
				}).Return(context.DeadlineExceeded)
			},
			mockLoggerBehavior: func(m *loggermocks.ILogger) {
				m.On("Debug", "[update recipe event] execute repository")
				m.On("Error", "request timed out while update recipe event", "err", context.DeadlineExceeded)
			},
			in: in{
				ctx: ctx,
				dto: dto,
			},
			want: want{
				recipeEvent: recipeevent.Recipe{},
				err:         errors.New("the request timed out: context deadline exceeded"),
			},
		},
		{
			name: "database error",
			mockPoolBehavior: func(m *poolsmocks.IPool, row *poolsmocks.RowMock) {
				m.On("QueryRow",
					mock.Anything,
					mock.Anything,
					dto.TypeID,
					dto.Title,
					dto.ID,
				).Return(row)

				row.On("Scan",
					mock.AnythingOfType("*int64"),
					mock.AnythingOfType("*int64"),
					mock.AnythingOfType("*string"),
					mock.AnythingOfType("*[][]recipeevent.Content"),
					mock.AnythingOfType("*time.Time"),
					mock.AnythingOfType("*time.Time"),
				).Run(func(args mock.Arguments) {
					id := args.Get(0).(*int64)
					*id = testRecipeEvent.ID

					typeID := args.Get(1).(*int64)
					*typeID = testRecipeEvent.TypeID

					title := args.Get(2).(*string)
					*title = testRecipeEvent.Title

					content := args.Get(3).(*[][]recipeevent.Content)
					*content = testRecipeEvent.Content

					createdAt := args.Get(4).(*time.Time)
					*createdAt = testRecipeEvent.CreatedAt

					updatedAt := args.Get(5).(*time.Time)
					*updatedAt = testRecipeEvent.UpdatedAt
				}).Return(errors.New("database error"))
			},
			mockLoggerBehavior: func(m *loggermocks.ILogger) {
				m.On("Debug", "[update recipe event] execute repository")
				m.On("Error", "failed to update recipe event", "err", errors.New("database error"))
			},
			in: in{
				ctx: ctx,
				dto: dto,
			},
			want: want{
				recipeEvent: recipeevent.Recipe{},
				err:         errors.New("could not update recipe event: database error"),
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

			assert.Equal(t, test.want.recipeEvent, result)

			mockPool.AssertExpectations(t)
			mockLogger.AssertExpectations(t)
			mockRow.AssertExpectations(t)
		})
	}
}
