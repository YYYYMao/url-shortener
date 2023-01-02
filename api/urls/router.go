package urls

import (
	"github.com/gin-gonic/gin"
)

func RouteUrls(r *gin.Engine, urlController *UrlController) {

	posts := r.Group("/api/v1/")
	{
		posts.POST("/urls", urlController.CreateUrl)
	}
}
