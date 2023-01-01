package repo

import (
	"time"
	"urlshortener/repositories/model"

	"github.com/stretchr/testify/mock"
)

type MockUrlRepo struct {
	mock.Mock
}

func (m *MockUrlRepo) Create(urlId string, url string, expireAt time.Time) (err error) {
	args := m.Called(urlId, url, expireAt)
	return args.Error(0)
}

func (m *MockUrlRepo) SelectByUrlId(urlId string) (*model.Urls, error) {
	args := m.Called(urlId)
	return args.Get(0).(*model.Urls), args.Error(1)
}

func (m *MockUrlRepo) IsUrlIdExist(urlId string) bool {
	args := m.Called(urlId)
	return args.Bool(0)
}
