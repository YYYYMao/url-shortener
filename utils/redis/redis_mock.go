package redis

import (
	"context"
	"time"

	"github.com/stretchr/testify/mock"
)

type MockRepo struct {
	mock.Mock
}

func (m *MockRepo) Set(ctx context.Context, key string, value interface{}, exp time.Duration) error {
	args := m.Called(ctx, key, value, exp)
	return args.Error(0)
}

func (m *MockRepo) Get(ctx context.Context, key string) (string, error) {
	args := m.Called(ctx, key)
	return args.String(0), args.Error(1)
}
