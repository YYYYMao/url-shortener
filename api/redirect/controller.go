package redirect

import (
	"net/http"
	"urlshortener/utils/resHandler"

	"github.com/gin-gonic/gin"
)

type RedirectController struct {
	RedirectService RedirectService
}

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
func (u *RedirectController) Redirect(c *gin.Context) {
	urlId := c.Param("url_id")
	if url, err := u.RedirectService.GetUrl(c.Request.Context(), urlId); err == nil {
		c.Redirect(http.StatusFound, url)
		return
	} else {
		resHandler.SendResponse(c, http.StatusNotFound, err, nil)
		return
	}
}
