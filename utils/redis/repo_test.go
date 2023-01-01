package redis

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
)

type Mockrepo struct {
	mock.Mock
}

func (m *Mockrepo) Set(ctx context.Context, key string, value interface{}, exp time.Duration) error {
	args := m.Called(ctx, key, value, exp)
	return args.Error(0)
}

func (m *Mockrepo) Get(ctx context.Context, key string) (string, error) {
	args := m.Called(ctx, key)
	return args.String(0), args.Error(1)
}

func Test_Set(t *testing.T) {
	m := new(Mockrepo)

	m.On("Set", context.Background(), "test123", "key", time.Hour).Return(nil)
	m.Set(context.Background(), "test123", "key", time.Hour)

	m.AssertExpectations(t)
}

func Test_Get(t *testing.T) {
	m := new(Mockrepo)

	m.On("Get", context.Background(), "test123").Return("key", nil)
	m.Get(context.Background(), "test123")

	m.AssertExpectations(t)
}
