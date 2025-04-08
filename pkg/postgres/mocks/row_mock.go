package mocks

import (
	"testing"

	"github.com/stretchr/testify/mock"
)

type MockRow struct {
	mock.Mock
}

func NewMockRow(t *testing.T) *MockRow {
	m := &MockRow{}
	m.Test(t)
	t.Cleanup(func() { m.AssertExpectations(t) })
	return m
}

func (m *MockRow) Scan(dest ...interface{}) error {
	args := m.Called(dest...)
	return args.Error(0)
}
