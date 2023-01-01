package redirect

import (
	"github.com/gin-gonic/gin"
)

func RouteUrls(r *gin.Engine, redirectController *RedirectController) {
	url := r.Group("")
	{
		url.GET("/:url_id", redirectController.Redirect)
	}
}
