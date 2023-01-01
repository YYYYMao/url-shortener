package redirect

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"
	repo "urlshortener/repositories"
	"urlshortener/utils"
	"urlshortener/utils/redis"
)

var expirePeriod = 7 * 24 * time.Hour

type cacheKey struct {
	Exp int64
	Url string
}

type redirectService struct {
	UrlRepo repo.UrlRepository
	Cache   redis.Repository
}

type RedirectService interface {
	GetUrl(c context.Context, urlId string) (string, error)
}

func NewRedirectService(urlRepo repo.UrlRepository, cacheRepo redis.Repository) RedirectService {
	return &redirectService{
		UrlRepo: urlRepo,
		Cache:   cacheRepo,
	}
}

func (s *redirectService) GetUrl(ctx context.Context, urlId string) (string, error) {
	now := time.Now().Unix()

	if !utils.VerifyUrlId(urlId) {
		return "", errors.New("not found")
	}

	if result, err := s.Cache.Get(ctx, urlId); err == nil {
		var data cacheKey
		if unmarshalErr := json.Unmarshal([]byte(result), &data); unmarshalErr == nil {
			v := data.Exp
			if v > now {
				return data.Url, nil
			} else {
				return "", errors.New("url expired")
			}
		}
	}

	if result, err := s.UrlRepo.SelectByUrlId(urlId); err == nil {
		if result.ExpireAt.Unix() > now {
			data := cacheKey{
				Exp: result.ExpireAt.Unix(),
				Url: result.Url,
			}

			if value, err := json.Marshal(data); err == nil {
				if insertErr := s.Cache.Set(ctx, result.UrlId, value, expirePeriod); insertErr != nil {
					fmt.Println("redis insert err: ", insertErr)
				}
			} else {
				fmt.Println("redis Marshal err: ", data, err)
			}

			return result.Url, nil
		}
	}

	return "", errors.New("not found")
}
