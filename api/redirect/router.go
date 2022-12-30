package redirect

import (
	"github.com/gin-gonic/gin"
)

func RouteUrls(r *gin.Engine) {
	url := r.Group("")
	{
		url.GET("/:url_id", RedirectHandler)
	}
}
