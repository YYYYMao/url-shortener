package urls

import (
	"errors"
	"time"
	repo "urlshortener/repositories"
	"urlshortener/utils"
)

type urlService struct {
	UrlRepo repo.UrlRepository
}

type UrlService interface {
	Create(url string, expireAt time.Time) (string, error)
}

func NewUrlService(urlRepo repo.UrlRepository) UrlService {
	return &urlService{
		UrlRepo: urlRepo,
	}
}

func (s *urlService) Create(url string, expireAt time.Time) (string, error) {
	urlId := utils.RandStringRunes(6)
	if err := s.UrlRepo.Create(urlId, url, expireAt); err != nil {
		return "", errors.New("db create url fail " + err.Error())
	}
	return urlId, nil
}
