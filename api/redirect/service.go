package redirect

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
	repo "urlshortener/repositories"
	"urlshortener/utils"
	"urlshortener/utils/redis"
	"urlshortener/utils/resHandler"

	"github.com/gin-gonic/gin"
)

var expirePeriod = 7 * 24 * time.Hour

// @Summary redirect url
// @Schemes
// @Description redirect url
// @Tags urls
// @Produce json
// @param url_id path string true "url_id"
// @Success 302
// @Success 404 {object} resHandler.ErrResponse
// @Success 500 {object} resHandler.ErrResponse
// @Router /{url_id} [get]
func RedirectHandler(c *gin.Context) {
	urlId := c.Param("url_id")
	now := time.Now().Unix()

	if !utils.VerifyUrlId(urlId) {
		resHandler.SendResponse(c, http.StatusNotFound, errors.New("not found"), nil)
		return
	}
	redisRepo := redis.NewRedisRepository(redis.Client)
	if result, err := redisRepo.Get(c.Request.Context(), urlId); err == nil {
		data := make(map[string]interface{})
		if unmarshalErr := json.Unmarshal([]byte(result), &data); unmarshalErr == nil {
			if v, ok := data["exp"].(int64); ok {
				if v > now {
					c.Redirect(http.StatusFound, data["url"].(string))
					return
				} else {
					resHandler.SendResponse(c, http.StatusNotFound, errors.New("url expired"), nil)
					return
				}
			}
		}
	}

	if result, err := repo.SelectUrl(urlId); err == nil {

		if result.ExpireAt.Unix() > now {
			data := make(map[string]interface{})
			data["exp"] = result.ExpireAt.Unix()
			data["url"] = result.Url

			if value, err := json.Marshal(data); err == nil {
				if insertErr := redisRepo.Set(c, result.UrlId, value, expirePeriod); insertErr != nil {
					fmt.Println("redis insert err: ", insertErr)
				}
			} else {
				fmt.Println("redis Marshal err: ", data, err)
			}
			c.Redirect(http.StatusFound, result.Url)
			return
		}
	}

	resHandler.SendResponse(c, http.StatusNotFound, errors.New("not found"), nil)

}
