package redis

import (
	"context"
	"testing"
	"time"
)

func Test_Set(t *testing.T) {
	m := new(MockRepo)

	m.On("Set", context.Background(), "test123", "key", time.Hour).Return(nil)
	m.Set(context.Background(), "test123", "key", time.Hour)

	m.AssertExpectations(t)
}

func Test_Get(t *testing.T) {
	m := new(MockRepo)

	m.On("Get", context.Background(), "test123").Return("key", nil)
	m.Get(context.Background(), "test123")

	m.AssertExpectations(t)
}
