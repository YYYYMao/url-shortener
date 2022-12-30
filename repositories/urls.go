package repo

import (
	"fmt"
	"time"
	"urlshortener/repositories/model"
	"urlshortener/utils/pg"
)

var fields = []string{"url_id", "url", "expire_at"}

func SelectUrl(urlId string) (*model.Urls, error) {
	url := &model.Urls{}
	if err := pg.Db.Select(fields).Where("url_id=?", urlId).First(&url); err.Error != nil {
		return nil, err.Error
	} else {
		return url, nil
	}
}

func CreateUrl(urlId string, url string, expireAt time.Time) error {
	if !CheckUrl(urlId) {
		return fmt.Errorf("url exists.")
	}
	newUrl := model.Urls{
		UrlId:    urlId,
		Url:      url,
		ExpireAt: expireAt,
	}
	insertErr := pg.Db.Model(&model.Urls{}).Create(&newUrl).Error
	return insertErr
}

func CheckUrl(urlId string) bool {
	var url model.Urls

	if err := pg.Db.Select(fields).Where("url_id=?", urlId).First(&url); err.RowsAffected == 0 {
		return true
	}
	return false
}
