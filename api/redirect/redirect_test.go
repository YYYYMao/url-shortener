package redirect

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	"urlshortener/repositories/model"
	repo "urlshortener/repositories/urls"
	"urlshortener/utils"
	"urlshortener/utils/redis"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRedirectService struct {
	mock.Mock
}

func (m *MockRedirectService) GetUrl(c context.Context, urlId string) (string, error) {
	args := m.Called(c, urlId)
	return args.String(0), args.Error(1)
}

func Test_RedirectSuccess(t *testing.T) {

	urlId := "oByA54204"
	url := "http://www.google.com"

	mockRedirectService := new(MockRedirectService)
	mockRedirectService.On("GetUrl", context.Background(), urlId).Return(url, nil)

	rr := httptest.NewRecorder()

	router := gin.Default()
	redirectController := &RedirectController{
		RedirectService: mockRedirectService,
	}

	router.GET("/:url_id", redirectController.Redirect)

	request, err := http.NewRequest(http.MethodGet, "/"+urlId, nil)
	assert.NoError(t, err)

	router.ServeHTTP(rr, request)

	assert.NoError(t, err)

	assert.Equal(t, http.StatusFound, rr.Code)
	mockRedirectService.AssertExpectations(t)
}

func Test_RedirectFail(t *testing.T) {

	urlId := "test123"

	mockRedirectService := new(MockRedirectService)
	mockRedirectService.On("GetUrl", context.Background(), urlId).Return("", errors.New("not found"))

	rr := httptest.NewRecorder()

	router := gin.Default()
	redirectController := &RedirectController{
		RedirectService: mockRedirectService,
	}

	router.GET("/:url_id", redirectController.Redirect)

	request, err := http.NewRequest(http.MethodGet, "/"+urlId, nil)
	assert.NoError(t, err)

	router.ServeHTTP(rr, request)
	assert.Equal(t, http.StatusNotFound, rr.Code)
	mockRedirectService.AssertExpectations(t)
}

func Test_GetUrlSuccessWithCache(t *testing.T) {

	urlId := utils.RandStringRunes(6)
	url := "http://www.google.com"
	ctx := context.Background()

	data := cacheKey{
		Exp: (time.Now().Unix() + 3600),
		Url: url,
	}

	key, _ := json.Marshal(data)

	urlRepo := new(repo.MockUrlRepo)
	cacheRepo := new(redis.MockRepo)
	cacheRepo.On("Get", ctx, urlId).Return(string(key), nil)

	redirectService := NewRedirectService(urlRepo, cacheRepo)
	result, err := redirectService.GetUrl(ctx, urlId)
	assert.NoError(t, err)
	assert.Equal(t, url, result)
}

func Test_GetUrlExpiredWithCache(t *testing.T) {

	urlId := utils.RandStringRunes(6)
	url := "http://www.google.com"
	ctx := context.Background()

	data := cacheKey{
		Exp: time.Now().Unix(),
		Url: url,
	}

	key, _ := json.Marshal(data)

	urlRepo := new(repo.MockUrlRepo)
	cacheRepo := new(redis.MockRepo)
	cacheRepo.On("Get", ctx, urlId).Return(string(key), nil)

	redirectService := NewRedirectService(urlRepo, cacheRepo)
	result, err := redirectService.GetUrl(ctx, urlId)
	assert.Equal(t, "url expired", err.Error())
	assert.Equal(t, "", result)
}

func Test_GetUrlSuccessWithDB(t *testing.T) {

	urlId := utils.RandStringRunes(6)
	url := "http://www.google.com"
	ctx := context.Background()

	expireAt := time.Now().Add(time.Hour)
	urlInfo := &model.Urls{
		UrlId:    urlId,
		ExpireAt: expireAt,
		Url:      url,
	}

	data := cacheKey{
		Exp: expireAt.Unix(),
		Url: url,
	}

	key, _ := json.Marshal(data)

	urlRepo := new(repo.MockUrlRepo)
	cacheRepo := new(redis.MockRepo)
	cacheRepo.On("Get", ctx, urlId).Return("", errors.New(""))
	cacheRepo.On("Set", ctx, urlId, key, expirePeriod).Return(nil)
	urlRepo.On("SelectByUrlId", urlId).Return(urlInfo, nil)

	redirectService := NewRedirectService(urlRepo, cacheRepo)
	result, err := redirectService.GetUrl(ctx, urlId)
	assert.NoError(t, err)
	assert.Equal(t, url, result)
}

func Test_GetUrlExpiredWithDB(t *testing.T) {

	urlId := utils.RandStringRunes(6)
	url := "http://www.google.com"
	ctx := context.Background()

	urlInfo := &model.Urls{
		UrlId:    urlId,
		ExpireAt: time.Now().Add(-1 * time.Hour),
		Url:      url,
	}

	urlRepo := new(repo.MockUrlRepo)
	cacheRepo := new(redis.MockRepo)
	cacheRepo.On("Get", ctx, urlId).Return("", errors.New(""))
	urlRepo.On("SelectByUrlId", urlId).Return(urlInfo, nil)

	redirectService := NewRedirectService(urlRepo, cacheRepo)
	result, err := redirectService.GetUrl(ctx, urlId)
	assert.Equal(t, "not found", err.Error())
	assert.Equal(t, "", result)
}

func Test_GetUrlNotFound(t *testing.T) {
	urlId := utils.RandStringRunes(6)
	ctx := context.Background()

	urlInfo := &model.Urls{}

	urlRepo := new(repo.MockUrlRepo)
	cacheRepo := new(redis.MockRepo)
	cacheRepo.On("Get", ctx, urlId).Return("", errors.New(""))
	urlRepo.On("SelectByUrlId", urlId).Return(urlInfo, nil)

	redirectService := NewRedirectService(urlRepo, cacheRepo)
	result, err := redirectService.GetUrl(ctx, urlId)
	assert.Equal(t, "not found", err.Error())
	assert.Equal(t, "", result)
}

func Test_GetUrlInvalid(t *testing.T) {

	urlId := "123"
	ctx := context.Background()

	urlRepo := new(repo.MockUrlRepo)
	cacheRepo := new(redis.MockRepo)

	redirectService := NewRedirectService(urlRepo, cacheRepo)
	result, err := redirectService.GetUrl(ctx, urlId)
	assert.Equal(t, "not found", err.Error())
	assert.Equal(t, "", result)
}
