package urls

import (
	"github.com/gin-gonic/gin"
)

func RouteUrls(r *gin.Engine) {
	posts := r.Group("/api/v1/")
	{
		posts.POST("/urls", CreateUrl)
	}
}
