package urlsRepo

import (
	"errors"
	"time"
	"urlshortener/repositories/model"

	"gorm.io/gorm"
)

var fields = []string{"url_id", "url", "expire_at"}

type urlRepository struct {
	db *gorm.DB
}

func NewUrlRepository(db *gorm.DB) UrlRepository {
	return &urlRepository{
		db: db,
	}
}

type UrlRepository interface {
	SelectByUrlId(urlId string) (*model.Urls, error)
	Create(urlId string, url string, expireAt time.Time) error
	IsUrlIdExist(urlId string) bool
}

func (r *urlRepository) SelectByUrlId(urlId string) (*model.Urls, error) {
	url := &model.Urls{}
	if err := r.db.Select(fields).Where("url_id=?", urlId).First(&url); err.Error != nil {
		return nil, err.Error
	} else {
		return url, nil
	}
}

func (r *urlRepository) Create(urlId string, url string, expireAt time.Time) error {
	if !r.IsUrlIdExist(urlId) {
		return errors.New("url exists")
	}
	newUrl := model.Urls{
		UrlId:    urlId,
		Url:      url,
		ExpireAt: expireAt,
	}
	insertErr := r.db.Model(&model.Urls{}).Create(&newUrl).Error
	return insertErr
}

func (r *urlRepository) IsUrlIdExist(urlId string) bool {
	var url model.Urls

	if err := r.db.Select(fields).Where("url_id=?", urlId).First(&url); err.RowsAffected == 0 {
		return true
	}
	return false
}
