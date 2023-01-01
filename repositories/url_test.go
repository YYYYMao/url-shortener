package repo

import (
	"testing"
	"time"
	"urlshortener/repositories/model"
)

func Test_Create(t *testing.T) {
	m := new(MockUrlRepo)

	url := &model.Urls{
		UrlId:    "test123",
		ExpireAt: time.Now(),
		Url:      "https:\\www.google.com",
	}

	m.On("Create", url.UrlId, url.Url, url.ExpireAt).Return(nil)
	m.Create(url.UrlId, url.Url, url.ExpireAt)

	m.AssertExpectations(t)
}

func Test_SelectByUrlId(t *testing.T) {
	m := new(MockUrlRepo)
	url := &model.Urls{
		UrlId:    "test123",
		ExpireAt: time.Now(),
		Url:      "https:\\www.google.com",
	}
	m.On("SelectByUrlId", "test123").Return(url, nil)
	m.SelectByUrlId("test123")

	m.AssertExpectations(t)
}

func Test_IsUrlIdExist(t *testing.T) {
	m := new(MockUrlRepo)
	m.On("IsUrlIdExist", "test123").Return(true, nil)
	m.IsUrlIdExist("test123")

	m.AssertExpectations(t)
}
