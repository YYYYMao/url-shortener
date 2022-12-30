package urls

import (
	"net/http"
	"os"
	"time"
	repo "urlshortener/repositories"
	"urlshortener/utils"
	"urlshortener/utils/resHandler"

	"github.com/gin-gonic/gin"
)

type UrlParam struct {
	Url      string    `json:"url" binding:"required"`
	ExpireAt time.Time `json:"expireAt" binding:"required"`
}

type UrlResponse struct {
	UrlId    string `json:"id"`
	ShortUrl string `json:"shortUrl"`
}

// @Summary create short url
// @Schemes
// @Description create short url
// @Tags urls
// @Accept json
// @Produce json
// @param data body UrlParam true "url and expire time 2022-12-30T15:03:43.4Z"
// @Success 200 {object} UrlResponse
// @Success 400 {object} resHandler.ErrResponse
// @Success 500 {object} resHandler.ErrResponse
// @Router /api/v1/urls [post]
func CreateUrl(c *gin.Context) {
	var param UrlParam
	if err := c.BindJSON(&param); err != nil {
		resHandler.SendResponse(c, http.StatusBadRequest, err, nil)
		return
	}

	urlId := utils.RandStringRunes(6)

	if err := repo.CreateUrl(urlId, param.Url, param.ExpireAt); err != nil {
		resHandler.SendResponse(c, http.StatusInternalServerError, err, nil)
		return
	}
	result := UrlResponse{
		UrlId:    urlId,
		ShortUrl: os.Getenv("DOMAIN") + urlId,
	}
	resHandler.SendResponse(c, http.StatusOK, nil, result)
}
