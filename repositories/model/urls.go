package model

import (
	"time"

	"gorm.io/gorm"
)

type Urls struct {
	gorm.Model
	UrlId     string    `gorm:"primary_key;size:10;not null;unique;index" json:"url_id"`
	Url       string    `gorm:"not null;" json:"url"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	ExpireAt  time.Time `gorm:"TIMESTAMP" json:"expire_at"`
}
