package urls

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	"urlshortener/repositories/model"
	repo "urlshortener/repositories/urls"
	"urlshortener/utils"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUrlService struct {
	mock.Mock
}

// ClearProfileImage is a mock of UserService.ClearProfileImage
func (m *MockUrlService) Create(url string, expireAt time.Time) (string, error) {
	args := m.Called(url, expireAt)
	return args.String(0), args.Error(1)
}

func Test_CreateUrlSuccess(t *testing.T) {
	if err := godotenv.Load("../../.env"); err != nil {
		fmt.Println("Error loading .env file")
	}

	urlId := "test123"
	url := "http://www.google.com"
	expireAt := "2022-12-30T15:03:43.000Z"
	api := "/api/v1/urls"

	layout := "2006-01-02T15:04:05.000Z"

	exAt, _ := time.Parse(layout, expireAt)

	mockUrlService := new(MockUrlService)
	mockUrlService.On("Create", url, exAt).Return(urlId, nil)

	rr := httptest.NewRecorder()

	router := gin.Default()
	urlController := &UrlController{
		UrlService: mockUrlService,
	}

	router.POST(api, urlController.CreateUrl)

	reqBody, _ := json.Marshal(gin.H{
		"url":      url,
		"expireAt": expireAt,
	})

	request, err := http.NewRequest(http.MethodPost, "/api/v1/urls", bytes.NewBuffer(reqBody))
	assert.NoError(t, err)

	router.ServeHTTP(rr, request)

	responseBody, err := json.Marshal(gin.H{
		"id":       urlId,
		"shortUrl": "http://127.0.0.1:8080/" + urlId,
	})

	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, responseBody, rr.Body.Bytes())
	mockUrlService.AssertExpectations(t)
}

func Test_CreateUrlBadRequest(t *testing.T) {

	url := "http://www.google.com"
	api := "/api/v1/urls"

	mockUrlService := new(MockUrlService)
	rr := httptest.NewRecorder()

	router := gin.Default()
	urlController := &UrlController{
		UrlService: mockUrlService,
	}

	router.POST(api, urlController.CreateUrl)

	reqBody, _ := json.Marshal(gin.H{
		"url": url,
	})

	request, err := http.NewRequest(http.MethodPost, "/api/v1/urls", bytes.NewBuffer(reqBody))
	assert.NoError(t, err)

	router.ServeHTTP(rr, request)

	assert.NoError(t, err)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
	mockUrlService.AssertExpectations(t)
}

func Test_CreateUrlDbFail(t *testing.T) {
	if err := godotenv.Load("../../.env"); err != nil {
		fmt.Println("Error loading .env file")
	}

	url := "http://www.google.com"
	expireAt := "2022-12-30T15:03:43.000Z"
	api := "/api/v1/urls"

	layout := "2006-01-02T15:04:05.000Z"

	exAt, _ := time.Parse(layout, expireAt)

	mockUrlService := new(MockUrlService)
	mockUrlService.On("Create", url, exAt).Return("", errors.New("insert fail"))

	rr := httptest.NewRecorder()

	router := gin.Default()
	urlController := &UrlController{
		UrlService: mockUrlService,
	}

	router.POST(api, urlController.CreateUrl)

	reqBody, _ := json.Marshal(gin.H{
		"url":      url,
		"expireAt": expireAt,
	})

	request, err := http.NewRequest(http.MethodPost, "/api/v1/urls", bytes.NewBuffer(reqBody))
	assert.NoError(t, err)

	router.ServeHTTP(rr, request)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	mockUrlService.AssertExpectations(t)
}

func Test_Create(t *testing.T) {
	mockUrlRepo := new(repo.MockUrlRepo)
	urlId := utils.RandStringRunes(6)

	url := &model.Urls{
		UrlId:    urlId,
		ExpireAt: time.Now(),
		Url:      "https:\\www.google.com",
	}

	mockUrlRepo.On("Create", url.UrlId, url.Url, url.ExpireAt).Return(nil)
	mockUrlRepo.Create(url.UrlId, url.Url, url.ExpireAt)
	mockUrlRepo.AssertExpectations(t)

}
